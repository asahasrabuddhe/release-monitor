package server

import (
	"context"
	"fmt"
	"github.com/asahasrabuddhe/rmpb"
	"log"
	"time"
)

type ReleaseMonitorServer struct {
	quit   chan struct{}
	ticker *time.Ticker
}

func (rms *ReleaseMonitorServer) Start(ctx context.Context, req *rmpb.StartRequest) (*rmpb.StartResponse, error) {
	rms.quit = make(chan struct{})
	rms.ticker = time.NewTicker(time.Duration(req.GetInterval()) * time.Second)
	var counter int
	log.Println("Starting.")
	go func() {
		for {
			select {
			case <-rms.ticker.C:
				fmt.Println("Work done ", counter)
				counter++
			case <-rms.quit:
				log.Println("Received Stop Signal. Stopping.")
				rms.ticker.Stop()
				return
			}
		}
	}()
	return &rmpb.StartResponse{}, nil
}

func (rms *ReleaseMonitorServer) Stop(context.Context, *rmpb.StopRequest) (*rmpb.StopResponse, error) {
	log.Println("Stop called")
	close(rms.quit)
	return &rmpb.StopResponse{}, nil
}
