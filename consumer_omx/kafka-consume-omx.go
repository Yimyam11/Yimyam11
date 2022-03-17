package main

import (
	"consumer/config"
	"fmt"
	"strings"
 	"consumer/process"
	"consumer/formatomx"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/gommon/log"
	"encoding/json"
)

func main() {

	fmt.Println("Helllo")
	Consume()
}
//var  KafkaFormat formatomx.Omxkafka
	// Config for this Application
 
// Consume func
func Consume() {
 
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
	 
	 c.SubscribeTopics(strings.Split(config.Config.KafkaConsumeTopic,","),nil)
 
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			logInfo := log.JSON{
				"type":    "receiveMessage",
				"rawData": string(msg.Value),
			}
			log.Debugj(logInfo)
		 	//input := &KafkaFormat.kafka_message_input{}
			 input :=  &formatomx.Omxkafka{}
		 	if  err := json.Unmarshal(msg.Value,input); err != nil {
				log.Error("invalid msg : " + string(msg.Value))
			}else {
			 	 
				  if input.Order.OrderType == 3 || input.Order.OrderType == 4  {
				 	res := process.Do(input) 
					fmt.Println(res)
				  }
			}		  
		} else {
			log.Error("Consumer error : " + err.Error())
		}
	}	
}		 
 