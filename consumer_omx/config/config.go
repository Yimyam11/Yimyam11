package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

//KafkaClient {
///	com.sun.security.auth.module.Krb5LoginModule required
//	useKeyTab=true
//	keyTab="/keytab/U-SVC-DMC.keytab"
/////	principal="U-SVC-DMC@TRUE.TH";
//};

//Envconfig reconfig from env
type Envconfig struct {
	//MongoHost               string `default:"localhost:27017"`
	//MongoUser               string `default:""`
	//MongoPass               string `default:""`
	KafkaConsumeHost        string `default:"tykbpr01:9092,tykbpr02:9092,tykbpr03:9092"`
	KafkaConsumeTopic       string `default:"PROD-json-MB-Postpaid-Aftersales-Online,PROD-json-MB-Postpaid-Aftersales-Bulk"` 
	KafkaConsumeGroup       string `default:"OMX_DMC"`
	KafkaConsumeKeytab      string `default:"config/keytab/U-SVC-DMC.keytab"`
	KafkaConsumePrincipal   string `default:"U-SVC-DMC@TRUE.TH"`
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
