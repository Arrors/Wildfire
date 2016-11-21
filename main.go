package main

import (
	"log"
	"net"

	"Wildfire/register"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	register.Regist(lis)
}
