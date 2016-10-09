package main

import (
	"github.com/influxdata/surgemq/service"
	"log"
)

func main() {
	log.Print("starting...\n")

	// Create a new server
	svr := &service.Server{
		KeepAlive:        300,           // seconds
		ConnectTimeout:   2,             // seconds
		SessionsProvider: "mem",         // keeps sessions in memory
		Authenticator:    "mockSuccess", // always succeed
		TopicsProvider:   "mem",         // keeps topic subscriptions in memory
	}

	// Listen and serve connections at localhost:1883
	svr.ListenAndServe("tcp://:1883")

	log.Print("exit...\n")
}
