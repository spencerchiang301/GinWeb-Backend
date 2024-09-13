package messaing

import (
	"github.com/segmentio/kafka-go"
	"gopkg.in/ini.v1"
	"log"
	"time"
)

type KafkaConfig struct {
	Brokers   []string
	Partition int
	MinBytes  int
	MaxBytes  int
}

// Load Kafka configuration from the config file
func loadKafkaConfig() (*KafkaConfig, error) {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		return nil, err
	}

	kafkaConfig := &KafkaConfig{
		Brokers: []string{
			cfg.Section("Kafka").Key("brokers_1").String(),
			cfg.Section("Kafka").Key("brokers_2").String(),
			cfg.Section("Kafka").Key("brokers_3").String(),
		},
		Partition: cfg.Section("Kafka").Key("Partition").MustInt(),
		MinBytes:  cfg.Section("Kafka").Key("MinBytes").MustInt(10000),
		MaxBytes:  cfg.Section("Kafka").Key("MaxBytes").MustInt(10000000),
	}

	return kafkaConfig, nil
}

func KafkaWriter() *kafka.Writer {
	config, err := loadKafkaConfig()
	if err != nil {
		log.Fatalf("Failed to load Kafka config: %v", err)
	}

	writer := kafka.Writer{
		Addr:         kafka.TCP(config.Brokers...),
		Topic:        "test-topic",
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    10,
		BatchTimeout: 500 * time.Millisecond,
		Async:        true,
	}
	return &writer
}

func KafkaReader() *kafka.Reader {
	config, err := loadKafkaConfig()
	if err != nil {
		log.Fatalf("Failed to load Kafka config: %v", err)
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     config.Brokers, // Kafka broker addresses
		Topic:       "test-topic",
		GroupID:     "test-group",      // Use GroupID for consumer group (remove Partition if using this)
		StartOffset: kafka.FirstOffset, // Start from the beginning of the topic
		//MinBytes:  config.MinBytes,
		//MaxBytes:  config.MaxBytes,
	})

	return reader
}
