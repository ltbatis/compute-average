package main

import (
	"context"
	"fmt"
	"log"
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
	//TODO: implementar leitura de argumentos do terminal e percorrer por eles
	requests := []*averagepb.AverageRequest{
		&averagepb.AverageRequest{
			Average: &averagepb.Average{
				Number: 1,
			},
		},
		&averagepb.AverageRequest{
			Average: &averagepb.Average{
				Number: 2,
			},
		},
		&averagepb.AverageRequest{
			Average: &averagepb.Average{
				Number: 3,
			},
		},
		&averagepb.AverageRequest{
			Average: &averagepb.Average{
				Number: 4,
			},
		},
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
