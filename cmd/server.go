package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"

	route "mit-api/api/routes"
)

type application struct {
	config config
	logger *zap.SugaredLogger
	db     *gorm.DB
	cache  *redis.Client
}

type config struct {
	address string
	db      dbConfig
	cache   cacheConfig
}

type dbConfig struct {
	database string
	host     string
	password string
	port     string
	username string
}

type cacheConfig struct {
	redisAddr string
	password  string
	redisDB   string
}

func (app *application) newServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

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
