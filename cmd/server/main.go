package main

import (
	"context"
	"em/internal/config"
	"em/internal/db"
	"em/internal/service"
	"em/internal/web"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// loads values from .env
func init() {
	if err := godotenv.Load("config/config.env"); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime) // launching the logger

	// Init, Connect and Migrate database
	database := db.NewRepository(&config.PSQLConnection{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DB"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Username: os.Getenv("POSTGRES_USER"),
	}, logger).Connect()

	psql, _ := database.DB()
	db.ApplyMigrations(database)
	log.Print("Database connected")

	// Init WebController, Service, Routes
	wApp := fiber.New()
	srv := service.CreateNewService(logger, database)
	ctrl := web.CreateNewWebController(wApp, logger, srv)
	ctrl.RegisterRouters()

	// start service and graceful shutdown
	go func() {
		if err := wApp.Listen(":1200"); err != nil {
			logger.Fatal("Error while listening")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	psql.Close()
	defer cancel()

	if err := wApp.Shutdown(); err != nil { // try to stop server
		logger.Print("Failed to stop server")

		return
	}

	logger.Print("Server stopped")
}
