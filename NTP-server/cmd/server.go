/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/MariumAZ/Zenly-tests/tree/main/NTP-server/time"
	ntp "github.com/beevik/ntp"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Fetch NTP and gives timestamp",
	Long: `----------------------------------------------------.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
	}
	// instantiate a new server
		s := grpc.NewServer()
	// register the new server
		pb.RegisterTimeServer(s, &Server{})
	//fmt.Println(ntp.Time("time.apple.com"))
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
	}
	},
}


const (
	port = ":50051"
)

// define the server struct that will implement the server
type Server struct {
	pb.UnimplementedTimeServer
}

func (s *Server) GetTime(ctx context.Context, req *pb.TimeRequest) (*pb.TimeReply, error){
	// connect to a remote ntp server 
	ntpserver := req.Ntpserver 
	if  ntpserver == "" {
		fmt.Println("ntpserver name must not be empty in the request")		
	}
	log.Printf("Received: %v", ntpserver)
	time, error := ntp.Time(ntpserver)
	if error != nil {
		log.Fatalf("Failed %v : ", error)
	
	}
	timestamp := time.Unix()
    return &pb.TimeReply{Timestamp : timestamp}, nil
	}
func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
