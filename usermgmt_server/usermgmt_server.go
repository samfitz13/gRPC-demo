package main

import (
	pb "example.com/go-usermgmt-grpc/usermgmt"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	pb.UnimplementedUserManagementServer
}
