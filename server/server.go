package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	vh "gRPC_experiment/videohandler"

	"google.golang.org/grpc"
)

const (
	port = ":21000"
)

type server struct{}

func (s *server) GetVideos(ctx context.Context, in *vh.VideoRequest) (*vh.VideoResponse, error) {
	test := []*vh.QueuedVideo{&vh.QueuedVideo{
		Video: &vh.Video{
			Title: "test",
			Time:  123,
			Size:  234,
		},
		Priority: 333,
	}, &vh.QueuedVideo{
		Video: &vh.Video{
			Title: "test",
			Time:  123,
			Size:  234,
		},
		Priority: 333,
	}, &vh.QueuedVideo{
		Video: &vh.Video{
			Title: "test",
			Time:  123,
			Size:  234,
		},
		Priority: 333,
	}}
	//test := []*vh.QueuedVideo{&vh.QueuedVideo{}}
	v := &vh.VideoResponse{Encoded: nil, Response: nil, Queued: test, Encoding: nil}
	_ = v
	v2 := &vh.VideoResponse{}
	_ = v2
	return v, nil
}

func (s *server) CancelEncoding(ctx context.Context, in *vh.CancelEncodingRequest) (*vh.StatusResponse, error) {
	return &vh.StatusResponse{}, nil
}

func (s *server) UpdatePriority(ctx context.Context, in *vh.UpdatePriorityRequest) (*vh.StatusResponse, error) {
	return &vh.StatusResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
		os.Exit(1)
	}
	s := grpc.NewServer()
	vh.RegisterVideoHandlerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
