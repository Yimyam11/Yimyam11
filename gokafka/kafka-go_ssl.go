package kafkaconsume

import (

	//"geofencing/config"
	//"geofencing/logic/process"

	// "geofencing/logic/process"
	//"geofencing/models"

	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Consume func
func Consume() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "10.95.148.163:9092,10.95.148.164:9092,10.95.148.165:9092",
		"group.id":          "logs",
		"auto.offset.reset": "smallest",
		//"sasl.kerberos.principal": config.Config.KafkaConsumePrincipal,
		//"sasl.kerberos.keytab":    config.Config.KafkaConsumeKeytab,
		//"security.protocol":       "sasl_plaintext",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"dmc-ocs"}, nil)
	// c.SubscribeTopics([]string{"mldd-scab", "^aRegex.*[Tt]opic"}, nil)

	// start := time.Now()
	// i := 0

	///// yim ///
	retrieved := c.Poll(0)
	switch message := retrieved.(type) {
	case *kafka.Message:
		//fmt.Printf(string(message.value()))
		fmt.Printf("%% data : %v\n", message)
	case kafka.Error:
		fmt.Printf("%% Error: %v\n", message)
	default:
		fmt.Printf("Ignored %v\n", message)
	}

	// for {
	// 	msg, err := c.ReadMessage(-1)
	// 	if err == nil {
	// 		logInfo := log.JSON{
	// 			"type":    "receiveMessage",
	// 			"rawData": string(msg.Value),
	// 		}
	// 		log.Debugj(logInfo)

	// 		input := &models.Input{}
	// 		if err := json.Unmarshal(msg.Value, input); err != nil {
	// 			log.Error("invalid msg : " + string(msg.Value))
	// 		}

	// 		process.Do(input)
	// 	} else {
	// 		log.Error("Consumer error : " + err.Error())
	// 	}
	// 	// i++

	// 	// if i == 10000 {
	// 	//  fmt.Println(time.Since(start))
	// 	//  //  os.Exit(0)
	// 	// }
	// }
}
