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

	pb "github.com/MariumAZ/Zenly-tests/tree/main/Go-store/pkg/gopher"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)
const (
    port = ":9000"
)


// server is used to implement gopher.GopherServer.

type GopherServer struct {
    //CreateQuery(ctx context.Context, req *pb.GopherRequest)(*pb.GopherReply, error)
	//GetQuery(ctx context.Context, req *pb.GopherRequest)(*pb.GopherReply, error)
	//UpdateQuery(ctx context.Context, req *pb.GopherRequest)(*pb.GopherReply, error)
}

type Gopher struct {
    URL string `json: "url"`
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen %v", err)
		}
		// Create new gRPC server with (blank) options
		grpcServer := grpc.NewServer()
		// Create BlogService type 
		srv := &GopherServer{}
		pb.RegisterGopherServer(grpcServer, srv)
		log.Printf("GRPC server listening on %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
            log.Fatalf("failed to serve: %v", err)
        }
	},
}


// implements set query 
func (s *GopherServer) CreateQuery (ctx context.Context, req *pb.GopherRequest)(*pb.GopherReply, error){  
	res := &pb.GopherReply{}
	if req == nil {
        fmt.Println("request must not be nil")
        return res, xerrors.Errorf("request must not be nil")
    }

    if req.Name == "" {
        fmt.Println("name must not be empty in the request")
        return res, xerrors.Errorf("name must not be empty in the request")
    }
} 


//implements get query 
func (s *GopherServer) GetQuery (ctx context.Context, req *pb.GopherRequest)(*pb.GopherReply, error){
		
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
