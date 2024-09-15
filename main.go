package main

import (
	"fmt"
	"log"
	"mit-api/server"

	"github.com/joho/godotenv"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @host localhost:8080
// @BasePath /v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Provide your Bearer token in the format: Bearer <token>
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	server := server.NewServer()

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
