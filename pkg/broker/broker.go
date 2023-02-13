package broker

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

var Connections = connectionVariables{}

type connectionVariables struct {
	NotificationConn *kafka.Conn
}

// setting up broker for the system
func SetupBroker() {

	brokerTopic := os.Getenv("BROKER_TOPIC")
	brokers := os.Getenv("BROKER_HOST")

	groupId := os.Getenv("BROKER_GROUP_ID")

	ConfigureNotificationBroker(brokers, groupId, brokerTopic)

}

// configuring notifications broker  topic connection
func ConfigureNotificationBroker(kafkaBrokerUrls string, groupId string, topic string) {

	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaBrokerUrls, topic, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	Connections.NotificationConn = conn
}

func PushMessage(message string) {
	Connections.NotificationConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err := Connections.NotificationConn.WriteMessages(
		kafka.Message{Value: []byte(message)},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

}

func ReadMessages(ctx context.Context) {
	brokerTopic := os.Getenv("BROKER_TOPIC")
	brokers := os.Getenv("BROKER_HOST")

	groupId := os.Getenv("BROKER_GROUP_ID")

	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{brokers},
		Topic:       brokerTopic,
		Partition:   0,
		GroupID:     groupId,
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
		StartOffset: kafka.LastOffset,
	})

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("Received message %s\n", string(m.Value))
		r.CommitMessages(ctx, m)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
