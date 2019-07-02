package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"math/rand"
	"testing"
	"time"
)

func TestProducer(t *testing.T) {
	fmt.Printf("producer_test\n")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("producer_test create producer error :%s\n", err.Error())
		return
	}

	defer producer.AsyncClose()

	// send message
	msg := &sarama.ProducerMessage{
		Topic: "yikeapp",
		Key:   sarama.StringEncoder("go_test"),
	}

	value := fmt.Sprintf("this is message. time=%v", time.Now())
	for {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Scanln(&value)
		msg.Value = sarama.ByteEncoder(value)
		fmt.Printf("input [%s]\n", value)

		// send to chain
		producer.Input() <- msg

		select {
		case suc := <-producer.Successes():
			fmt.Printf("offset: %d,  timestamp: %s", suc.Offset, suc.Timestamp.String())
		case fail := <-producer.Errors():
			fmt.Printf("err: %s\n", fail.Err.Error())
		}
	}
}

func TestConsumer(t *testing.T) {
	fmt.Printf("consumer_test")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	// consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()

	partition_consumer, err := consumer.ConsumePartition("yikeapp", 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partition_consumer.Close()

	for {
		select {
		case msg := <-partition_consumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		case err := <-partition_consumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}
}
