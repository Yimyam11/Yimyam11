package main

import (
	"consumer/config"
	"consumer/formatomx"
	"consumer/logger"
	"consumer/process"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/gommon/log"
)

func main() {

	fmt.Println("Helllo")
	Consume()
}

// Config for this Application

// Consume func
func Consume() {
	filename := "ResultNotificationOMX_Kafka.txt"
	var OutputOmx []byte
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":          config.Config.KafkaConsumeHost,
		"group.id":                   config.Config.KafkaConsumeGroup,
		"auto.offset.reset":          config.Config.KafkaConsumeOffsetReset,
		"sasl.kerberos.principal":    config.Config.KafkaConsumePrincipal,
		"sasl.kerberos.keytab":       config.Config.KafkaConsumeKeytab,
		"security.protocol":          "sasl_ssl",
		"sasl.mechanisms":            "GSSAPI",
		"sasl.kerberos.service.name": "kafka",
		"ssl.ca.location":            "config/cert/ca.pem",
		"ssl.cipher.suites":          "DHE-DSS-AES256-GCM-SHA384",
		//"debug" : "broker",
		//"ssl.truststore.location":    "truststore/tykbpr.client.truststore.jks",
		//"ssl.truststore.password":    "maserati10m",
		//	"sasl.jaas.config":           "com.sun.security.auth.module.Krb5LoginModule required useTicketCache=true;",
	})

	if err != nil {
		panic(err)

	}

	c.SubscribeTopics(strings.Split(config.Config.KafkaConsumeTopic, ","), nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			logInfo := log.JSON{
				"type":    "receiveMessage",
				"rawData": string(msg.Value),
			}
			log.Debugj(logInfo)
			input := &formatomx.Omxkafka{}

			if err := json.Unmarshal(msg.Value, input); err != nil {
				log.Error("invalid msg : " + string(msg.Value))
			} else {
				//	fmt.Println(string(msg.Value))
				// if ((input.Order.Channel == "DMC" || input.Order.Channel == "DMC1" || input.Order.Channel == "DMC2"  || input.Order.Channel == "DMC3" ) && (input.Order.OrderType == 3 || input.Order.OrderType == 4  || input.Order.OrderType == 41))  {
				if input.Order.Channel == "DMC" || input.Order.Channel == "DMC1" || input.Order.Channel == "DMC2" || input.Order.Channel == "DMC3" {
					//	fmt.Println(string(msg.Value))
					if input.Order.OrderType == 3 || input.Order.OrderType == 4 {
						OutputOmx, err = json.Marshal(process.Do(input))
						if err != nil {
							fmt.Println(err)
						}
						logger.LogFile(OutputOmx, filename)
					}
					//logger.LogFile(*(*[]byte)(unsafe.Pointer(input)), filename)
					logger.LogFile([]byte(string(msg.Value)), filename)
					//	// print out //
					// 	fmt.Println(string(OutputOmx))
					//	// write to file ////	fmt.Println(input)

				}
			}
		} else {
			log.Error("Consumer error : " + err.Error())
		}
	}
}
