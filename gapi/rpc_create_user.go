package gapi

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/milindmadhukar/gRPC-tutorial/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
  fmt.Println("CreateUser called")

  fmt.Println("Username:", req.GetUsername())

  if regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(req.GetMail()) {
    fmt.Println("Mail:", req.GetMail())
  } else {
    return nil, status.Errorf(codes.InvalidArgument, "Invalid mail address")
  }

  resp := &pb.CreateUserResponse{
  	User: &pb.User{
  		Username:  req.GetUsername(),
  		Name:      req.GetName(),
  		Mail:      req.GetMail(),
  		CreatedAt: timestamppb.New(time.Now()),
  	},
  }


  return resp, nil


}
