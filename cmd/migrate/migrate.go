package main

import (
	"context"
	"fmt"
	pb "github.com/dalbar/gedis/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddr := "gedis:5431"
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Unable to connect to db instance", err.Error())
		return
	}
	defer conn.Close()
	client := pb.NewGedisClient(conn)

	testContext := context.Background()
	client.PutPair(testContext, &pb.Pair{Key: "animal", Value: "dog"})
	client.PutPair(testContext, &pb.Pair{Key: "fruit", Value: "orange"})
	client.PutPair(testContext, &pb.Pair{Key: "foo", Value: "bar"})
	client.PutPair(testContext, &pb.Pair{Key: "furniture", Value: "table"})

	if err != nil {
		fmt.Println("Unable to write value from db", err.Error())
		return
	}

	dump, _ := client.Dump(testContext, &pb.Empty{})
	for i, p := range dump.Pairs {
		fmt.Println(fmt.Sprintf("%d: (%s, %s)", i+1, p.Key, p.Value))
	}
}
