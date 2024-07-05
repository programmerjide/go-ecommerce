package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/programmerolajide/go-ecommerce/config"
	"github.com/programmerolajide/go-ecommerce/internal/api/rest"
	"github.com/programmerolajide/go-ecommerce/internal/dto"
	"github.com/programmerolajide/go-ecommerce/internal/helper"
	"github.com/programmerolajide/go-ecommerce/internal/repository"
	"github.com/programmerolajide/go-ecommerce/internal/service"
	"log"
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
		Auth: rh.Auth,
	}

	handler := UserHandler{
		svc: svc,
	}

	publicRoutes := app.Group("api/v1/users")

	// Public endpoints
	publicRoutes.Post("/register", handler.Register)
	publicRoutes.Post("/login", handler.Login)

	privateRoutes := publicRoutes.Group("/", rh.Auth.Authorize)

	// Private endpoints
	privateRoutes.Get("/verify", handler.GetVerificationCode)
	privateRoutes.Post("/verify", handler.Verify)
	privateRoutes.Post("/profile", handler.CreateProfile)
	privateRoutes.Get("/profile", handler.GetProfile)

	privateRoutes.Post("/cart", handler.AddToCart)
	privateRoutes.Get("/cart", handler.GetCart)
	privateRoutes.Post("/order", handler.CreateOrder)
	privateRoutes.Post("/order/:id", handler.GetOrder)
	privateRoutes.Get("/orders", handler.GetOrders)

	privateRoutes.Post("/order/:id", handler.BecomeSeller)

}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignupRequestDto{}
	if err := ctx.BodyParser(&user); err != nil {
		return helper.RespondWithError(ctx, http.StatusBadRequest, config.INVALID_PAYLOAD.Code, config.INVALID_PAYLOAD.Description)
	}

	signupResponseData, err := h.svc.Signup(user)
	if err != nil {
		return helper.RespondWithError(ctx, http.StatusBadRequest, config.FAILED.Code, "Error occurred while processing the request")
	}

	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: config.SUCCESS.Description,
			Data:    signupResponseData,
		},
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {

	loginDto := dto.UserLoginDto{}
	if err := ctx.BodyParser(&loginDto); err != nil {
		return helper.RespondWithError(ctx, http.StatusBadRequest, config.INVALID_PAYLOAD.Code, config.INVALID_PAYLOAD.Description)
	}

	token, err := h.svc.Login(loginDto.Email, loginDto.Password)
	if err != nil {
		return helper.RespondWithError(ctx, http.StatusBadRequest, config.FAILED.Code, "Invalid credentials")
	}

	loginResponseData := dto.LoginResponseData{
		AccessToken:  token,
		RefreshToken: "",
	}

	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: config.SUCCESS.Description,
			Data:    loginResponseData,
		},
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	// create verification code and update to user profile in DB
	code, err := h.svc.GetVerificationCode(user)
	if err != nil {
		log.Printf("Error generating verification code: %v", err)
		return helper.RespondWithError(ctx, http.StatusBadRequest, config.FAILED.Code, "Error generating verification code")
	}

	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: "Verification code generated successfully",
			Data: dto.VerificationCodeData{
				VerificationCode: code,
			},
		},
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)

	var req dto.VerificationCodeInput

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid verification code",
		})
	}

	err := h.svc.VerifyCode(user.ID, req.Code)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: "User verified successfully",
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

	user := h.svc.Auth.GetCurrentUser(ctx)

	response := dto.DefaultApiResponse{
		BaseResponse: dto.BaseResponse[any]{
			Status:  config.SUCCESS.Code,
			Message: "user profile retrieved successfully",
			Data:    user,
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
