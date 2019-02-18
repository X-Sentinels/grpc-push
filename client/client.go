//
// client.go
//
// gRPC Push Notification Client
//

package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	pb "github.com/X-Sentinels/grpc-push/protos"

	"google.golang.org/grpc"
)

// register ...
func register(client pb.PushNotifClient, name string) error {
	log.Printf("Calling Register RPC")

	stream, err := client.Register(context.Background(), &pb.RegistrationRequest{ClientName: name})
	if err != nil {
		log.Fatalf("Register failed %v", err)
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to recv response")
		}
		log.Print(resp, err)
	}
	return nil
}

// main ...
func main() {
	addr := flag.String("a", "127.0.0.1:50051", "grpc server address")
	flag.Parse()
	// init important structures
	rand.Seed(time.Now().UTC().UnixNano())
	name := fmt.Sprintf("%d", rand.Intn(50))

	// Setup a connection with the server
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := pb.NewPushNotifClient(conn)

	go register(client, name)
	select {}
}
