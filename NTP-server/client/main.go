package main

import (
	"context"
	"fmt"
	"net"
	"log"
	"google.golang.org/grpc"
	pb "github.com/MariumAZ/Zenly-tests/tree/main/NTP-server/time"
	ntp "github.com/beevik/ntp"
)

const (
	port = ":50051"
)

// define the server struct that will implement the server
type Server struct {
	pb.UnimplementedTimeServer
}

func (s *Server) GetTime(ctx context.Context, req *pb.TimeRequest) (*pb.TimeReply){
	// connect to a remote ntp server 
	ntpserver := req.Ntpserver 
	if  ntpserver == "" {
		fmt.Println("ntpserver name must not be empty in the request")		
	}
	log.Printf("Received: %v", ntpserver)
	time, _ := ntp.Time(ntpserver)
	timestamp := time.Unix()
    return &pb.TimeReply{Timestamp : timestamp}
	}


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// instantiate a new server
	s := grpc.NewServer()
	// register the new server
	pb.RegisterTimeServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
