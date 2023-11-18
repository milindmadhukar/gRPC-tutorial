package gapi

import "github.com/milindmadhukar/gRPC-tutorial/pb"

type Server struct {
  pb.UnimplementedAPIServer
  // NOTE: You can add other fields
}

func NewServer() (*Server, error) {
  return &Server{}, nil
}
