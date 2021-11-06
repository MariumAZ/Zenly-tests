package main

import (
    "fmt"
    ntp "github.com/beevik/ntp"
	_"context"
	pb "github.com/MariumAZ/Zenly-tests/tree/main/NTP-server/time"

)

// define the server struct that will implement the server 
type Server struct {
	pb.UnimplementedTimeServer
}
/*
func (s *Server) GetTime(ctx context.Context, req *pb.TimeRequest) (*pb.TimeReply){
	// connect to a remote ntp server 
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	
	// check if ntp server is missing
	if req.Ntpserver == "" {
		fmt.Println("ntpserver name must not be empty in the request")		
	}
	

	}

*/
func main() {
	//time, _:= ntp.Time("0.beevik-ntp.pool.ntp.org")
	time, _ := ntp.Time("time.microsoft.com")
	fmt.Println(time)
}
