package main

import (
	"context"
	"flag"
	"fmt"
	db "github.com/dalbar/gedis/pkg/db"
	nodeInterface "github.com/dalbar/gedis/pkg/node"

	pb "github.com/dalbar/gedis/pkg/grpc"
	"google.golang.org/grpc"
)

type DatabaseNode struct {
	ctx                 context.Context
	database            *db.Database
	id                  string
	instanceKeyTable    map[string]string
	instanceClientTable map[string]pb.GedisClient
	connections         []*grpc.ClientConn
	pb.UnimplementedGedisServer
}

// TODO I don't like global variables
var (
	address        *string
	id             *string
	totalInstances *int64
	initialSize    *int
	port           *int
	isReplicaMode  *bool

	leader *string
)

func init() {
	// TODO proper input validation -> let's assume happy path for now
	id = flag.String("id", "0", "unique identifier")

	totalInstances = flag.Int64("n", 1, "total instances including leader")

	initialSize = flag.Int("allocateN", 0, "pre-allocate the number of keys to store")

	port = flag.Int("port", 8000, "port to serve db interface")

	isReplicaMode = flag.Bool("replicaMode", false, "allows full replicas among all nodes")

	leader = flag.String("leader", "0", "should be a valid address to connect to (including port)")

	address = flag.String("listen", "", "alternative address for other instances to connect to")
}

func main() {
	// TODO no sane usage of flags -> we can crash
	flag.Parse()

	node := nodeInterface.New(
		fmt.Sprintf("%s:%d", *id, *port),
		*initialSize,
		*isReplicaMode,
	)

	node.Run(*port, *leader, *totalInstances, *address)
}
