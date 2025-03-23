package main

import (
	"context"
	shipping "go-grpc-micro"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:8085", opts...)
	if err != nil {
		log.Fatalf("Failed to connect shipping service. Err: %v", err)
	}

	defer conn.Close()

	shippingClient := shipping.NewShippingServiceClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	log.Println("Creating shipping...")
	_, errCreate := shippingClient.Create(ctx, &shipping.CreateShippingRequest{UserId: 23})
	if errCreate != nil {
		log.Printf("Failed to create shipping. Err: %v", errCreate)
	} else {
		log.Println("Shipping is created successfully.")
	}
}
