package main

import (
	"log"
	"ww-api-gateway/api"
)

func main() {
	dbConnectionString := "mongodb://root:password123@localhost:27017"
	port := "8080"
	apiConfig, err := api.NewConfig(dbConnectionString, port)
	if err != nil {
		panic(err)
	}
	log.Fatalln(api.Start(apiConfig))
}
