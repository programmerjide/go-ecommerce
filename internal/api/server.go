package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/programmerolajide/go-ecommerce/config"
	"github.com/programmerolajide/go-ecommerce/internal/api/rest"
	"github.com/programmerolajide/go-ecommerce/internal/api/rest/handlers"
	"github.com/programmerolajide/go-ecommerce/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	// connect the ORM
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection failed with error: %v\n", err)
	}

	// run migration
	db.AutoMigrate(&domain.User{})

	restHandler := &rest.RestHandler{
		App: app,
		DB:  db,
	}

	setupRoutes(restHandler)

	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
}
