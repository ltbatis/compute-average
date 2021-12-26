package main

import (
	"fmt"
	"log"

	"github.com/ltbatista/compute-average/average/averagepb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, i'm the avg client...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := averagepb.NewAverageServiceClient(cc)

	doClientStreaming(c)
}

func doClientStreaming(c averagepb.AverageServiceClient) {
	fmt.Println("Sarting to do a Client Streaming RPC...")
	//TODO: implementar client streaming requests
}
