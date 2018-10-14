package main

import (
	"fmt"
	"github.com/asahasrabuddhe/release-monitor/server"
	"github.com/asahasrabuddhe/rmpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	rmpb.RegisterReleaseMonitorServiceServer(grpcServer, &server.ReleaseMonitorServer{})

	grpcServer.Serve(lis)
}
