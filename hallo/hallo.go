package hallo

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"servers/hallo/apidoc"
)

type server struct{}

var s = grpc.NewServer()

func (s *server) SayHallo(ctx context.Context, in *helloworld.HalloRequest) (*helloworld.HalloReply, error) {
	return &helloworld.HalloReply{Message: "Hello " + in.Name}, nil
}

func registAllServers() {
	helloworld.RegisterGreeterServer(s, &server{})
}

// RegistServer Regist service for listener
func RegistServer(lis net.Listener) {

	registAllServers()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
