package helper

import "github.com/programmerolajide/go-ecommerce/internal/dto"

import (
	"github.com/gofiber/fiber/v2"
)

func RespondWithError(ctx *fiber.Ctx, status int, code string, message string) error {
	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  code,
			Message: message,
			Data:    nil,
		},
	}
	return ctx.Status(status).JSON(response)
}
