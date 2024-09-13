package global

import (
	"database/sql"
	"github.com/segmentio/kafka-go"
)

var (
	Mysql         *sql.DB
	MyKafkaWriter *kafka.Writer
	MyKafkaReader *kafka.Reader
)
