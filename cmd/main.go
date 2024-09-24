package main

import (
	"fmt"
	"log"
	"mit-api/internal/cache"
	"mit-api/internal/database"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
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

	cfg := config{
		address: os.Getenv("ADDRESS"),
		db: dbConfig{
			database: os.Getenv("DB_DATABASE"),
			password: os.Getenv("DB_PASSWORD"),
			username: os.Getenv("DB_USERNAME"),
			port:     os.Getenv("DB_PORT"),
			host:     os.Getenv("DB_HOST"),
		},
		cache: cacheConfig{
			redisAddr: os.Getenv("REDIS_ADDR"),
			password:  os.Getenv("REDIS_PASSWORD"),
			redisDB:   os.Getenv("REDIS_DB"),
		},
	}

	// Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	// Connect to cache
	dbNumber, err := strconv.Atoi(cfg.cache.redisDB)
	if err != nil {
		log.Fatalf("Failed to convert redisDB to int: %v", err)
	}
	cache.Connect(cfg.cache.redisAddr, cfg.cache.password, dbNumber)

	// Connect to database
	db, err := database.Connect(cfg.db.host, cfg.db.username, cfg.db.password, cfg.db.database, cfg.db.port)
	if err != nil {
		logger.Fatal(err.Error())
	}

	app := &application{
		config: cfg,
		logger: sugar,
		db:     db,
	}

	server := app.newServer()

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("Cannot start server: %s", err))
	}
}
