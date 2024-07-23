package main

import (
	"database/Assignment/http-client/internal/handlers"
	"database/Assignment/http-client/internal/repository"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database connection
	dbURL := os.Getenv("DATABASE_URL")
	db, err := repository.NewDatabase(dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repository
	userRepo := repository.NewUserRepository(db)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(*userRepo)

	// Router setup
	router := mux.NewRouter()
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	router.HandleFunc("/search", userHandler.SearchUsersByAddress).Methods("GET")

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
