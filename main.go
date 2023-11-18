package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/milindmadhukar/gRPC-tutorial/gapi"
	"github.com/milindmadhukar/gRPC-tutorial/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	server, err := gapi.NewServer()
	if err != nil {
		panic(err)
	}

	go runGRPCServer(server)
	runGRPCtoRestGatewayServer(server)
}

func runGRPCServer(server *gapi.Server) {
	grpcServer := grpc.NewServer()

	pb.RegisterAPIServer(grpcServer, server)
	reflection.Register(grpcServer) // NOTE: Allows the client to discover the server and what RPCs are available and how to call them, kind of like self-documentations -TechSchool

	grpcListener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		panic(err)
	}

	log.Printf("gRPC to HTTP REST Gateway running on port %v\n", grpcListener.Addr().String())
	if err := grpcServer.Serve(grpcListener); err != nil {
		panic(err)
	}
}

func runGRPCtoRestGatewayServer(server *gapi.Server) {
  jsonOpts := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
      DiscardUnknown: true,
    },
	})

	grpcMux := runtime.NewServeMux(jsonOpts)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := pb.RegisterAPIHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	httpListener, err := net.Listen("tcp", "localhost:9091")
	if err != nil {
		panic(err)
	}

	log.Printf("gRPC server running on port %v\n", httpListener.Addr().String())
	if err := http.Serve(httpListener, mux); err != nil {
		panic(err)
	}

}
