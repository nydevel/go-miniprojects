package main

import (
	"fmt"
	"log"
	"time"

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

	// Subject to publish messages to
	subject := "demo.messages"

	// Publish messages in a loop
	counter := 0
	for {
		counter++
		message := fmt.Sprintf("Message #%d at %s", counter, time.Now().Format(time.RFC3339))

		err := nc.Publish(subject, []byte(message))
		if err != nil {
			log.Printf("Error publishing message: %v", err)
			continue
		}

		fmt.Printf("Published: %s\n", message)

		// Wait before sending next message
		time.Sleep(2 * time.Second)
	}
}
