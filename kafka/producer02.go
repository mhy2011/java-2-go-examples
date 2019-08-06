package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func main() {
	config := sarama.NewConfig()
	producer, e := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if e != nil {
		panic(e)
	}

	for i := 0; i < 10; i++ {
		content := fmt.Sprintf("Hello_%d", i)
		message := &sarama.ProducerMessage{Topic: "test01", Value: sarama.StringEncoder(content)}
		producer.Input() <- message

	}

	time.Sleep(1e9)
}
