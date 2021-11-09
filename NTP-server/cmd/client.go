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
	"fmt"
    "log"
	"context"
	"os"
	"time"
	"google.golang.org/grpc"
	pb "github.com/MariumAZ/Zenly-tests/tree/main/NTP-server/time"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Query NTP server",
	Long: ` -------------------------------------`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("client called")
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect : %v", err)
	    }
		defer conn.Close()
		c := pb.NewTimeClient(conn)
		// contact the server and print its response
		name := defaultName
		if len(os.Args) > 2 {
			name = os.Args[2]
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		r, err := c.GetTime(ctx, &pb.TimeRequest{Ntpserver: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("The Timestamp for %s is :  %d", name, r.GetTimestamp())
	},
}


const (
	address     = "localhost:50051"
	defaultName = "time.apple.com"
)
func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
