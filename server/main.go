package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("unable to listen on port 8080: %v", err)
	}

	srv := grpc.NewServer()
	//
	srv.Serve(listener)
}
