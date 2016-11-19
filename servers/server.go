package servers

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

const (
	port = ":50051"
)

var s = grpc.NewServer()

func init() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
