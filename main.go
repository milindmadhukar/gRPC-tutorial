package main

import (
	"log"
	"net"

	"github.com/milindmadhukar/gRPC-tutorial/gapi"
	"github.com/milindmadhukar/gRPC-tutorial/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
  server, err := gapi.NewServer()
  if err != nil {
    panic(err)
  }

  grpcServer := grpc.NewServer()

  pb.RegisterAPIServer(grpcServer, server)
  reflection.Register(grpcServer) // NOTE: Allows the client to discover the server and what RPCs are available and how to call them, kind of like self-documentations -TechSchool

  listener, err := net.Listen("tcp", "localhost:9090")
  if err != nil {
    panic(err)
  }

  log.Printf("gRPC server running on port %v\n", listener.Addr().String())

  if err := grpcServer.Serve(listener); err != nil {
    panic(err)
  }

}
