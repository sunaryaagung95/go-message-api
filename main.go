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
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Can't load env:%s", err)
	}
	fmt.Println("Ger env values")

	server.ConnectDB(
		os.Getenv("DB_HOST"),
		"5432",
		"postgres",
		"messenger",
		os.Getenv("DB_PASSWORD"),
	)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on localhost:%s", port)
	server.Serve(fmt.Sprintf(":%s", port))
}
