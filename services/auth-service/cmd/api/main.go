package main

import (
	"log"

	"donation-platform/auth-service/internal/config"
	"donation-platform/auth-service/internal/database"
	"donation-platform/auth-service/internal/router"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := router.Setup(db, cfg)

	port := cfg.AppPort
	if port == "" {
		port = "8081"
	}

	log.Printf("Auth Service running on :%s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
