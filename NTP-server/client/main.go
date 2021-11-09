package main

import (
	"context"
	"log"
	"os"
	"time"
	"google.golang.org/grpc"
	pb "github.com/MariumAZ/Zenly-tests/tree/main/NTP-server/time"
	//ntp "github.com/beevik/ntp"
)

const (
	address     = "localhost:50051"
	defaultName = "time.apple.com"
)

func main() {
    // set up a connection to the server 
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}
	defer conn.Close()
	c := pb.NewTimeClient(conn)
	// contact the server and print its response
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c.GetTime(ctx, &pb.TimeRequest{Ntpserver: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("The Timestamp for %s is :  %d", name, r.GetTimestamp())
}
