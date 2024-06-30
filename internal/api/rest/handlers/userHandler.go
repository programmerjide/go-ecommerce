package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/programmerolajide/go-ecommerce/config"
	"github.com/programmerolajide/go-ecommerce/internal/api/rest"
	"github.com/programmerolajide/go-ecommerce/internal/dto"
	"github.com/programmerolajide/go-ecommerce/internal/helper"
	"github.com/programmerolajide/go-ecommerce/internal/repository"
	"github.com/programmerolajide/go-ecommerce/internal/service"
	"net/http"
)

type UserHandler struct {
	// svc UserService
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {

	app := rh.App

	// create an instance of user service & inject to handler
	svc := service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
	}

	handler := UserHandler{
		svc: svc,
	}

	// Public endpoints
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// Private endpoints
	app.Get("/verify", handler.GetVerificationCode)
	app.Post("/verify", handler.Verify)
	app.Post("/profile", handler.CreateProfile)
	app.Get("/profile/:id", handler.GetProfile)

	app.Post("/cart", handler.AddToCart)
	app.Get("/cart", handler.GetCart)
	app.Post("/order", handler.CreateOrder)
	app.Post("/order/:id", handler.GetOrder)
	app.Get("/orders", handler.GetOrders)

	app.Post("/order/:id", handler.BecomeSeller)

}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignupRequestDTO{}
	if err := ctx.BodyParser(&user); err != nil {
		return helper.RespondWithError(ctx, http.StatusBadRequest, config.INVALID_PAYLOAD.Code, config.INVALID_PAYLOAD.Description)
	}

	token, err := h.svc.Signup(user)
	if err != nil {
		return helper.RespondWithError(ctx, http.StatusBadRequest, config.FAILED.Code, "Error occurred while processing the request")
	}

	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: token,
			Data:    nil,
		},
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {

	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: config.SUCCESS.Description,
			Data:    nil,
		},
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {

	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: config.SUCCESS.Description,
			Data:    nil,
		},
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {

	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: config.SUCCESS.Description,
			Data:    nil,
		},
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {

	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: config.SUCCESS.Description,
			Data:    nil,
		},
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {

	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: config.SUCCESS.Description,
			Data:    nil,
		},
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {

	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: config.SUCCESS.Description,
			Data:    nil,
		},
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {

	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: config.SUCCESS.Description,
			Data:    nil,
		},
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) CreateOrder(ctx *fiber.Ctx) error {

	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: config.SUCCESS.Description,
			Data:    nil,
		},
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {

	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: config.SUCCESS.Description,
			Data:    nil,
		},
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {

	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: config.SUCCESS.Description,
			Data:    nil,
		},
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {

	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: config.SUCCESS.Description,
			Data:    nil,
		},
	}
	return ctx.Status(http.StatusOK).JSON(response)
}
