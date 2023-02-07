package broker

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

var writer *kafka.Writer

var Connections = connectionVariables{}

type connectionVariables struct {
	NotificationWriter *kafka.Writer

	AlertWriter *kafka.Writer
}

// setting up broker for the system
func SetupBroker() {
	notificationBroker := os.Getenv("BROKER_NOTIFICATIONS")
	// alertBroker := os.Getenv("BROKER_ALERT")
	brokers := os.Getenv("BROKER_HOST")

	brokerUrls := strings.Split(brokers, ",")

	clientId := os.Getenv("BROKER_CLIENT_ID")

	fmt.Println("Setting up connection to the broker")

	ConfigureNotificationBroker(brokerUrls, clientId, notificationBroker)

	// ConfigureAlertBroker(brokerUrls, clientId, alertBroker)

}

// configuring notifications broker topic
func ConfigureNotificationBroker(kafkaBrokerUrls []string, clientId string, topic string) (*kafka.Writer, error) {

	fmt.Println(kafkaBrokerUrls, clientId, topic)
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientId,
	}

	config := kafka.WriterConfig{
		Brokers:          kafkaBrokerUrls,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	Connections.NotificationWriter = kafka.NewWriter(config)

	return Connections.NotificationWriter, nil
}

// // configuring alert broker topic
// func ConfigureAlertBroker(kafkaBrokerUrls []string, clientId string, topic string) {
// 	dialer := &kafka.Dialer{
// 		Timeout:  10 * time.Second,
// 		ClientID: clientId,
// 	}

// 	config := kafka.WriterConfig{
// 		Brokers:          kafkaBrokerUrls,
// 		Topic:            topic,
// 		Balancer:         &kafka.LeastBytes{},
// 		Dialer:           dialer,
// 		WriteTimeout:     10 * time.Second,
// 		ReadTimeout:      10 * time.Second,
// 		CompressionCodec: snappy.NewCompressionCodec(),
// 	}
// 	Connections.AlertWriter = kafka.NewWriter(config)
// }
