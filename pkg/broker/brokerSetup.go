package broker

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

var Connections = connectionVariables{}

type connectionVariables struct {
	NotificationConn *kafka.Conn
}

// setting up broker for the system
func SetupBroker() {
	notificationBroker := os.Getenv("BROKER_NOTIFICATIONS")
	brokers := os.Getenv("BROKER_HOST")

	clientId := os.Getenv("BROKER_CLIENT_ID")

	ConfigureNotificationBroker(brokers, clientId, notificationBroker)

}

// configuring notifications broker  topic connection
func ConfigureNotificationBroker(kafkaBrokerUrls string, clientId string, topic string) {

	fmt.Println(kafkaBrokerUrls, clientId, topic)
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaBrokerUrls, topic, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	Connections.NotificationConn = conn
}
