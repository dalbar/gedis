package node

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	db "github.com/dalbar/gedis/pkg/db"
	pb "github.com/dalbar/gedis/pkg/grpc"
)

type DatabaseNode struct {
	ctx context.Context

	database       *db.Database
	id             string
	initialSize    int
	isReplicaMode  bool
	isLeader       bool
	totalInstances int64

	instances           sync.Map
	connectedInstances  int64
	instanceKeyTable    map[string]string
	instanceClientTable map[string]pb.GedisClient
	connections         []*grpc.ClientConn

	pb.UnimplementedGedisServer
}

func New(id string, initialSize int, isReplicaMode bool) *DatabaseNode {
	node := DatabaseNode{
		ctx:                context.Background(),
		id:                 id,
		isReplicaMode:      isReplicaMode,
		database:           db.New(initialSize),
		connectedInstances: 1,

		instanceKeyTable: map[string]string{},
	}

	return &node
}

func (node *DatabaseNode) Run(port int, leader string, totalInstances int64, address string) {
	log.Printf("received leader %q, my node id %q", leader, node.id)
	node.totalInstances = totalInstances

	listener, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: port})
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGedisServer(s, node)
	reflection.Register(s)

	defer func() {
		for _, connection := range node.connections {
			connection.Close()
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs

		node.Cleanup()

		os.Exit(0)
	}()

	if node.id != leader {
		cc, _ := grpc.Dial(leader, grpc.WithTransportCredentials(insecure.NewCredentials()))
		client := pb.NewGedisClient(cc)
		instanceId := node.id

		if len(address) > 0 {
			instanceId = address
		}

		for {
			// TODO timeout channel to prevent endless waiting
			response, err := client.Heartbeat(node.ctx, &pb.InstanceKeyBroadCastReply{Instance: instanceId})
			if err != nil {
				log.Printf("Leader is not ready yet: %q", err.Error())
				time.Sleep(time.Second * 3)
				continue
			}

			node.InitiateClientTable(response.Instances)
			break
		}
	} else {
		node.isLeader = true
	}

	log.Printf("instance with id %q running", node.id)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (node *DatabaseNode) Cleanup() {
	for _, connection := range node.connections {
		connection.Close()
	}

	// TODO kill all connections when leader dies
}

func (node *DatabaseNode) GetValue(ctx context.Context, req *pb.KeyRequest) (*pb.Pair, error) {
	log.Printf("%q received a request to get value of key %q", node.id, req.Key)

	instance, ok := node.instanceKeyTable[req.Key]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Key %q not found in database", req.Key)
	}

	if instance != node.id {
		pair, err := node.instanceClientTable[instance].GetValue(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error(), req.Key)
		}

		return pair, nil
	}

	value, _ := node.database.Read(req.Key)
	return &pb.Pair{Key: req.Key, Value: value}, nil
}

func (node *DatabaseNode) PutPair(ctx context.Context, pair *pb.Pair) (*pb.Empty, error) {
	log.Printf("%q received a request to save value for pair (%s, %s)", node.id, pair.Key, pair.Value)

	node.database.Write(pair.Key, pair.Value)
	node.instanceKeyTable[pair.Key] = node.id

	log.Printf("sending broadcast for key %q", pair.Key)

	var wg sync.WaitGroup

	for _, client := range node.instanceClientTable {
		wg.Add(1)
		go func(targetClient pb.GedisClient) {
			node.BroadcastTillAck(ctx, targetClient, pair.Key, pb.InstanceKeyBroadcast_WRITE)
			wg.Done()
		}(client)
	}

	log.Printf("Waiting for broadcast to finish for key %q", pair.Key)

	wg.Wait()

	return &pb.Empty{}, nil
}

// TODO fix messy api
func (node *DatabaseNode) BroadcastTillAck(ctx context.Context, client pb.GedisClient, key string, operation pb.InstanceKeyBroadcast_Operation) {
	for {
		ack, err := client.Broadcast(ctx, &pb.InstanceKeyBroadcast{Instance: node.id, Key: key, Operation: operation})

		if err != nil {
			log.Println("Unable to broadcast write. Retrying!")
			time.Sleep(time.Second * 3)
			continue
		}

		log.Printf("successfully informed %q about %q", ack.Instance, key)
		break
	}
}

func (node *DatabaseNode) DeleteValue(ctx context.Context, req *pb.KeyRequest) (*pb.Empty, error) {
	log.Printf("%q received a request to delete value of key %q", node.id, req.Key)

	instance, ok := node.instanceKeyTable[req.Key]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Key %q not found in database", req.Key)
	}

	if instance != node.id {
		resp, err := node.instanceClientTable[instance].DeleteValue(ctx, req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	err := node.database.Delete(req.Key)
	if err != nil {
		return nil, err
	}

	delete(node.instanceKeyTable, req.Key)

	// TODO prefer channel
	var wg sync.WaitGroup

	for _, client := range node.instanceClientTable {
		wg.Add(1)
		go func(targetClient pb.GedisClient) {
			node.BroadcastTillAck(ctx, targetClient, req.Key, pb.InstanceKeyBroadcast_DELETE)
			wg.Done()
		}(client)
	}

	log.Printf("Waiting for broadcast to finish for key %q", req.Key)

	wg.Wait()

	return &pb.Empty{}, nil
}

func (node *DatabaseNode) Broadcast(ctx context.Context, req *pb.InstanceKeyBroadcast) (*pb.InstanceKeyBroadCastReply, error) {
	log.Printf("received a broadcast request from instance %q for key %q", req.Instance, req.Key)

	switch req.Operation {
	case pb.InstanceKeyBroadcast_WRITE:
		if node.isReplicaMode {
			node.instanceKeyTable[req.Key] = node.id
			node.database.Write(req.Key, req.Value)
		} else {
			node.instanceKeyTable[req.Key] = req.Instance
		}
	case pb.InstanceKeyBroadcast_DELETE:
		delete(node.instanceKeyTable, req.Key)
	}

	return &pb.InstanceKeyBroadCastReply{Instance: node.id}, nil
}

func (node *DatabaseNode) Dump(ctx context.Context, _ *pb.Empty) (*pb.DatabaseDump, error) {
	log.Println("Received a request to dump all data")

	dump := pb.DatabaseDump{}
	dumpChan := make(chan *pb.Pair, len(node.instanceKeyTable))

	var wg sync.WaitGroup

	for key, instance := range node.instanceKeyTable {
		wg.Add(1)
		go func(targetInstance string, targetKey string) {
			log.Printf("Dumping value for key %s", targetKey)

			// TODO maybe a more elegant solution
			if targetInstance == node.id {
				value, _ := node.database.Read(targetKey)
				dumpChan <- &pb.Pair{Key: targetKey, Value: value}
			} else {
				resp, err := node.instanceClientTable[targetInstance].GetValue(ctx, &pb.KeyRequest{Key: targetKey})
				if err != nil {
					dumpChan <- &pb.Pair{Key: targetKey, Value: "__Not-Found"}
				} else {
					dumpChan <- resp
				}
			}
			wg.Done()
		}(instance, key)
	}

	wg.Wait()
	close(dumpChan)

	for pair := range dumpChan {
		dump.Pairs = append(dump.Pairs, pair)
	}

	return &dump, nil
}

func (node *DatabaseNode) Heartbeat(ctx context.Context, instance *pb.InstanceKeyBroadCastReply) (*pb.HeartbeatResponse, error) {
	// leader gets called by other instances
	if !node.isLeader {
		return nil, nil
	}

	if node.connectedInstances == node.totalInstances {
		var instances []*pb.InstanceKeyBroadCastReply
		node.instances.Range(func(instance, _ any) bool {
			if instance != nil {
				instances = append(instances, &pb.InstanceKeyBroadCastReply{Instance: instance.(string)})
			}
			return true
		})

		instances = append(instances, &pb.InstanceKeyBroadCastReply{Instance: node.id})

		// TODO concurrency bug
		node.InitiateClientTable(instances)

		return &pb.HeartbeatResponse{Instances: instances}, nil
	}

	if _, ok := node.instances.Load(instance.Instance); !ok {
		node.instances.Store(instance.Instance, struct{}{})
		atomic.AddInt64(&node.connectedInstances, 1)
	}

	return nil, errors.New("waiting for other instances to connect")
}

func (node *DatabaseNode) InitiateClientTable(instances []*pb.InstanceKeyBroadCastReply) {
	node.instanceClientTable = make(map[string]pb.GedisClient, len(instances))

	for _, instance := range instances {
		if instance.Instance == node.id {
			continue
		}

		log.Printf("creating connection to %q", instance.Instance)

		conn, err := grpc.Dial(instance.Instance, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Unable to connect to db instance %q", instances)
		}

		node.connections = append(node.connections, conn)
		node.instanceClientTable[instance.Instance] = pb.NewGedisClient(conn)
	}
}
