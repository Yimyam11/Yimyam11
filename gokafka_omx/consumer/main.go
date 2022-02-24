package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	servers := []string{"10.95.148.163:9092"}
	consumer, err := sarama.NewConsumer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition("dmc-ocs", 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)

	}
	defer partitionConsumer.Close()
	for {
		select {
		case err := <-partitionConsumer.Errors():
			fmt.Println(err)
		case msg := <-partitionConsumer.Messages():
			fmt.Println(string(msg.Value))
		}
	}

}
