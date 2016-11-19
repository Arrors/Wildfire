package servers

import (
	"golang.org/x/net/context"

	pb "common/protos"
)

func (s *server) SayHallo(ctx context.Context, in *pb.HalloRequest) (*pb.HalloReply, error) {
	return &pb.HalloReply{Message: "Hello " + in.Name}, nil
}

func init() {
	pb.RegisterGreeterServer(s, &server{})
}
