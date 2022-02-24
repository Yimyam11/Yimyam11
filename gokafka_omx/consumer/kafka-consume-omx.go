package kafkaconsume

import (
	"fmt"

	"gokafka_omx/consumer/config/config.go"

	//"geofencing/logic/process"
	//	"geofencing/models"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/gommon/log"
)

// Consume func
func Consume() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":       config.Config.KafkaConsumeHost,
		"group.id":                config.Config.KafkaConsumeGroup,
		"auto.offset.reset":       config.Config.KafkaConsumeOffsetReset,
		"sasl.kerberos.principal": config.Config.KafkaConsumePrincipal,
		"sasl.kerberos.keytab":    config.Config.KafkaConsumeKeytab,
		"security.protocol":       "sasl_plaintext",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{config.Config.KafkaConsumeTopic}, nil)
	// c.SubscribeTopics([]string{"mldd-scab", "^aRegex.*[Tt]opic"}, nil)

	// start := time.Now()
	// i := 0
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			logInfo := log.JSON{
				"type":    "receiveMessage",
				"rawData": string(msg.Value),
			}
			log.Debugj(logInfo)

			//input := &models.Input{}
			//if err := json.Unmarshal(msg.Value, input); err != nil {
			log.Error("invalid msg : " + string(msg.Value))
			fmt.Println(string(msg.Value))
			//}

			//process.Do(input)
		} else {
			log.Error("Consumer error : " + err.Error())
		}
		// i++

		// if i == 10000 {
		// 	fmt.Println(time.Since(start))
		// 	// 	os.Exit(0)
		// }
	}
}
