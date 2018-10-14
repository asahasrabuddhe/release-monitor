package main

import (
	"context"
	"github.com/asahasrabuddhe/rmpb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := rmpb.NewReleaseMonitorServiceClient(conn)

	go func() {
		client.Start(context.Background(), &rmpb.StartRequest{Interval: 3})
	}()
	time.Sleep(6 * time.Second)
	client.Stop(context.Background(), &rmpb.StopRequest{Restart: false})
}
