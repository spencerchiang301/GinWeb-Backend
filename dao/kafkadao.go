package dao

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"web/global"
)

func SendMessage() error {
	writer := global.MyKafkaWriter

	// Send 100 messages to Kafka
	for i := 1; i <= 20000; i++ {
		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(fmt.Sprintf("Key-%d", i)),
				Value: []byte(fmt.Sprintf("Hello Kafka Golang user! Message number: %d", i)),
			},
		)

		log.Printf("Hello Kafka Golang user! Message number: %d", i)

		if err != nil {
			log.Printf("Error writing message to Kafka: %v", err)
			return err
		}
	}

	log.Printf("Successfully sent 100 messages to Kafka")
	return nil
}

func ReceiveTopic(consumerID string) error {
	reader := global.MyKafkaReader

	// Ensure the reader is properly closed when the function exits
	defer func(reader *kafka.Reader) {
		err := reader.Close()
		if err != nil {
			log.Printf("Error closing Kafka reader: %v", err)
		}
	}(reader)

	// Goroutine to consume messages from Kafka
	go func() {
		for {
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Printf("could not read message: %v", err)
				return
			}
			log.Printf("%s received message: %s\n", consumerID, string(msg.Value))
		}
	}()

	// Block until a signal is received (e.g., Ctrl+C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	fmt.Println("Shutting down consumer...")

	return nil
}
