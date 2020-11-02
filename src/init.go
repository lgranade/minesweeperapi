// +build !unit
// Only runs as daemon, skipped at unit tests

package main

import (
	"log"
	"os"
	"strconv"
)

func init() {
	var err error

	params.DBHost = os.Getenv("DB_HOST")
	params.DBPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Can't load config, err: ", err)
	}
	params.DBUser = os.Getenv("DB_USER")
	params.DBPassword = os.Getenv("DB_PASSWORD")
	params.DBName = os.Getenv("DB_NAME")
	params.DBSSLMode = os.Getenv("DB_SSLMODE")
}
