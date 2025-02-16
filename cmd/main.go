package main

import (
	"context"
	"log"
	"net"
	"noteapp/internal/handlers"
	"noteapp/kafka"
	"noteapp/pkg/db"
	"noteapp/proto/generated"
	"noteapp/redis"

	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db.InitDB()

	go func() {
		err := redis.InitRedis("redis", "6379")
		if err != nil {
			log.Fatalf("❌ Failed to initialize Redis: %v", err)
		}
		log.Println("✅ Redis initialized successfully")
	}()

	// Initialize Kafka producer
	kafkaProducer := kafka.NewProducer()
	defer kafkaProducer.Close() // Fix: Now `Close()` exists

	// Start Kafka consumer in a separate goroutine
	go func() {
		log.Println("🚀 Starting Kafka Consumer...")
		kafka.StartConsumer(ctx) // Fix: Now accepts `context.Context`
	}()

	// Start Redis-to-Kafka processing in a goroutine
	go func() {
		log.Println("🔄 Pushing Redis updates to Kafka...")
		kafkaProducer.PushUpdatesToKafka(ctx) // Fix: Now accepts `context.Context`
	}()

	// Start gRPC server
	server := grpc.NewServer()
	generated.RegisterUserServiceServer(server, &handlers.UserHandler{})
	generated.RegisterNoteServiceServer(server, &handlers.NoteHandler{})

	listener, err := net.Listen("tcp", ":7000")
	if err != nil {
		log.Fatalf("❌ Failed to listen on port 7000: %v", err)
	}

	log.Println("✅ gRPC server is running on port 7000 🚀")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("❌ Failed to serve gRPC: %v", err)
	}
}
