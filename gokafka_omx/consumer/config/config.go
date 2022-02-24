package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

//Envconfig reconfig from env
type Envconfig struct {
	//MongoHost               string `default:"localhost:27017"`
	//MongoUser               string `default:""`
	//MongoPass               string `default:""`
	KafkaConsumeHost        string `default:"localhost"`
	KafkaConsumeTopic       string `default:"PROD-json-MB-Postpaid-Aftersales-Online"`
	KafkaConsumeGroup       string `default:"OMX_DMC"`
	KafkaConsumeKeytab      string `default:"keytab/U-SVC-DMC.keytab"`
	KafkaConsumePrincipal   string `default:"gof_uat@TYB.RFT"`
	KafkaConsumeOffsetReset string `default:"earliest"`
	//KafkaProduceHost        string `default:"node1"`
	//KafkaProduceTopic       string `default:"whatsup-sms-geo-topic"`
	IntervalLoadConfig string `default:"1m"`
	LogLevel           string `default:"debug"`
}

var (
	// Config for this Application
	Config Envconfig
)

func init() {
	err := envconfig.Process("", &Config)
	if err != nil {
		log.Fatal(err.Error())
	}
}
