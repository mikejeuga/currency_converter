package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/mikejeuga/currency_converter/src/gateway/rateapi"
	"log"
)

func main() {
	var rateApiConfig rateapi.Config

	err := envconfig.Process("", &rateApiConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	client := rateapi.NewClient(rateApiConfig)

	rate, err := client.GetFXRate("GBP", "USD")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(rate)
}
