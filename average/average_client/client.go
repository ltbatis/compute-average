package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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
	requests := make([]*averagepb.AverageRequest, len(os.Args)-1)
	for i := 1; i < len(os.Args); i++ {
		numero, _ := strconv.Atoi(os.Args[i])
		requests[i-1] = &averagepb.AverageRequest{
			Average: &averagepb.Average{
				Number: int32(numero),
			},
		}
	}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Average: %v", err)
	}

	// we interate over our slice and send each message individually
	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from Averag: %v", err)
	}
	fmt.Printf("Average response: %v\n", res)
}
