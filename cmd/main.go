package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mikejeuga/currency_converter/config"
	"github.com/mikejeuga/currency_converter/src/clients/rateapi"
	"github.com/mikejeuga/currency_converter/src/domain/currency_conversion"
	"github.com/mikejeuga/currency_converter/src/web"
	"log"
)

func main() {
	var rateApiConfig rateapi.Config
	err := envconfig.Process("", &rateApiConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	var c config.Config
	err = envconfig.Process("kfkf", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	client := rateapi.NewClient(rateApiConfig)
	gateway := currency_conversion.NewService(client)
	server := web.NewServer(c, gateway)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
