package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sunaryaagung95/go-message-api/controllers"
)

var server = controllers.Server{}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		loadEnv()
		port = "8080"
	}

	server.ConnectDB(
		os.Getenv("DB_HOST"),
		"5432",
		"postgres",
		"messenger",
		os.Getenv("DB_PASSWORD"),
	)
	log.Printf("Listening on localhost:%s", port)
	server.Serve(fmt.Sprintf(":%s", port))
}

func loadEnv() {
	if os.Getenv("DB_HOST") == "" {
		var err error
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Can't load env:%s", err)
		}
		fmt.Println("Ger env values")
	}
}
