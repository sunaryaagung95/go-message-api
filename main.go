package main

import "fmt"

import "github.com/sunaryaagung/hello-world/controllers"

import "github.com/joho/godotenv"

import "log"

import "os"

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
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	server.Serve(":8080")
}
