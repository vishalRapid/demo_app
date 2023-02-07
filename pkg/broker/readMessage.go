package broker

import (
	"fmt"
	"time"
)

// function to read message from broker
func ReadMessage() {
	Connections.NotificationConn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := Connections.NotificationConn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}
}
