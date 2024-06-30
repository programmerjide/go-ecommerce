package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/programmerolajide/go-ecommerce/internal/helper"
	"gorm.io/gorm"
)

type RestHandler struct {
	App  *fiber.App
	DB   *gorm.DB
	Auth helper.Auth
}
