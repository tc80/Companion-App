package main

import (
	"context"
	"fmt"
	"os"
	"time"

	vh "gRPC_experiment/videohandler"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:21000"
	defaultName = "world"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure()) // look at transport security later
	if err != nil {
		fmt.Printf("Error connecting: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	c := vh.NewVideoHandlerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() // figure out exactly the purpose of cancel, context.background

	r, err := c.GetVideos(ctx, &vh.VideoRequest{
		Type:  vh.VideoType_ENCODED,
		Query: "test",
	})

	if err != nil {
		fmt.Printf("Server error: %v\n", err)
		os.Exit(1)
	}

	sr := r.Response
	quv := r.Queued
	egv := r.Encoding
	edv := r.Encoded

	if r == nil {
		fmt.Println("nil")
	} else if r != nil {
		fmt.Println("not nil")
	}
	fmt.Println("r is: ", r)
	fmt.Println("response is nil: ", sr == nil)
	fmt.Println("queued is nil: ", quv == nil)
	if quv != nil {
		for i, q := range quv {
			fmt.Printf("index: %v, value is nil?: %v\n", i, q == nil)
			fmt.Println("priority: ", q.Priority)
			fmt.Println("video is nil: ", q.Video == nil)
		}
	}

	// seems like server cannot have nil values when passing to client, but default pointer is nil
	os.Exit(5)

	// NO NEED FOR ERRORS NOW THAT WE RETURN IT FROM SERVER ANYWAYS

	_ = egv
	_ = edv

	fmt.Println(r.Response)

	// if sr.Status == nil {
	// 	fmt.Println("NIL")
	// }
	// os.Exit(2)

	if sr == nil {
		fmt.Println("nil")
	}
	if sr.Status == vh.Status_FAILURE {
		fmt.Printf("Failure: %v\n", sr.Error)
		os.Exit(1)
	}
	fmt.Println("Success, received")
	fmt.Printf("STATUS: %v\n", sr.Status)

	for _, q := range quv {
		fmt.Println("video: ", q) // try w/o directive
	}
}
