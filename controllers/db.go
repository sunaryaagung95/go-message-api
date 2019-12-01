package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	"github.com/sunaryaagung95/go-message-api/models"
)

//Server model
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

//ConnectDB func
func (server *Server) ConnectDB(dbHost, dbPort, dbUser, dbName, dbPassword string) {
	var err error

	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		dbHost,
		dbPort,
		dbUser,
		dbName,
		dbPassword,
	)
	server.DB, err = gorm.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("%s", dbName)
		log.Fatal("Error:", err)
	}
	fmt.Println("Connected to db:", dbName)
	server.DB.Debug().AutoMigrate(models.User{})
	server.Router = mux.NewRouter()
	server.RunRouters()
}

//Serve func
func (server *Server) Serve(addr string) {
	headers := handlers.AllowedHeaders([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedMethods([]string{"*"})
	fmt.Println("Listening to Port: 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(server.Router)))
}
