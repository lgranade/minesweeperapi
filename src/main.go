package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Params represents all env variables recevied at runtime
type Params struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

var params Params

func main() {
	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, syscall.SIGTERM)
	signal.Notify(stopChannel, syscall.SIGINT)

	log.Println("Starting api server ...")

	err := start()
	if err != nil {
		log.Fatal("Error starting api, err: ", err)
	}
	log.Println("Started")

	<-stopChannel
	log.Println("Received TERMSIGNAL")

	stop()

	log.Println("Shutted down gracefully")
}
