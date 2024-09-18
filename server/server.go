package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"mit-api/internal/cache"
	"mit-api/internal/database"
	routers "mit-api/routes"
)

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	database.New()
	cache.Connect()

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      routers.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
