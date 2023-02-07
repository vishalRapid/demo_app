package broker

import (
	"time"

	"github.com/segmentio/kafka-go"
)

// push new message to new topic
func PushData(connection *kafka.Conn, message string) {
	connection.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err := connection.WriteMessages(
		kafka.Message{Value: []byte(message)},
	)

	if err != nil {
		panic(err)
	}
}
