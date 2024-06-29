package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/programmerolajide/go-ecommerce/config"
	"github.com/programmerolajide/go-ecommerce/internal/dto"
	"net/http"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Get("/health", HealthCheck)

	app.Listen(config.ServerPort)
}

func HealthCheck(ctx *fiber.Ctx) error {
	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: config.SUCCESS.Description,
			Data:    nil, // You can set this to any data you want to include
		},
	}
	return ctx.Status(http.StatusOK).JSON(response)
}
