package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	fmt.Println("Connected to NATS server")

	// Subject to subscribe to
	subject := "demo.messages"

	// Subscribe to messages
	sub, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Printf("Received message: %s\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	fmt.Printf("Subscribed to subject: %s\n", subject)
	fmt.Println("Waiting for messages... (Press Ctrl+C to exit)")

	// Wait for interrupt signal to gracefully shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\nShutting down consumer...")
}
