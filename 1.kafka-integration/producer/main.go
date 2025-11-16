package main

import (
    "github.com/IBM/sarama"
    "log"
)

func main() {
    config := sarama.NewConfig()
    config.Producer.Return.Successes = true
    
    producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
    if err != nil {
        log.Fatal(err)
    }
    defer producer.Close()
    
    msg := &sarama.ProducerMessage{
        Topic: "test-topic",
        Value: sarama.StringEncoder("Hello Kafka"),
    }
    
    partition, offset, err := producer.SendMessage(msg)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Sent to partition %d, offset %d", partition, offset)
}