package main

import (
	// "log"

	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/config"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/database"
)

func main() {
	// Load environment variables
	cfg := config.Load()

	// database connect
	database.ConnectMongo(cfg.MongoUri, cfg.DbName)

	// Initialize new app
	app := SetupApp()

	// Listen at a given port
	// log.Println(cfg.Port)
	app.Listen("0.0.0.0:" + cfg.Port)
}