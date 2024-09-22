package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	route "mit-api/api/routes"
	"mit-api/internal/cache"
	"mit-api/internal/database"
)

type application struct {
	config config
}

type config struct {
	address string
}

func (app *application) newServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	database.New()
	cache.Connect()

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      route.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
