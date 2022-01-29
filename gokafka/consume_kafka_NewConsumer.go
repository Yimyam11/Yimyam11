package main

import (
        "github.com/confluentinc/confluent-kafka-go/kafka"
        "github.com/labstack/gommon/log"
)

func main(){
        Consume()
}




// Consume func
func Consume() {
        c, err := kafka.NewConsumer(&kafka.ConfigMap{
            "bootstrap.servers": "localhost:9092",
            "group.id":          "logs",
            "auto.offset.reset" : "earliest",
//            "debug":"all",
        })

        if err != nil {
                panic(err)
        }

        c.SubscribeTopics([]string{"myTopic1"}, nil)
        for {
                msg, err := c.ReadMessage(-1)
                if err == nil {
                        log.Info("++++++++ msg : " + string(msg.Value))
                } else {
                        log.Error("Consumer error : " + err.Error())
                }
        }
}
 