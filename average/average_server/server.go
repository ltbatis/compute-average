package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/ltbatista/compute-average/average/averagepb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Average(stream averagepb.AverageService_AverageServer) error {
	fmt.Println("Average function was invoked with a streaming request...")
	dividendo := 0
	divisor := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			result := fmt.Sprintf("The average is: %f", float64(dividendo)/float64(divisor))
			return stream.SendAndClose(&averagepb.AverageResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		dividendo += int(req.GetAverage().GetNumber())
		divisor++
	}
}

func main() {
	fmt.Println("Hello, I'm the avg server!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	averagepb.RegisterAverageServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
