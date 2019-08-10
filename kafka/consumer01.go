package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

func main() {
	consumer, e := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if e != nil {
		panic(e)
	}
	defer func() {
		if err:=consumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	partitionConsumer, e := consumer.ConsumePartition("test", 0, 0)
	if e != nil {
		panic(e)
	}

	defer func() {
		if err:=partitionConsumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for ; ; {
		message := <- partitionConsumer.Messages()
		fmt.Println("message =", string(message.Value))
	}


}
