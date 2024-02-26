package internaluser

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type UserHandlerIPLM struct {
	UserService *UserServiceIPLM
}

func NewUserHandler(userService *UserServiceIPLM) *UserHandlerIPLM {
	return &UserHandlerIPLM{
		UserService: userService,
	}
}

func (h *UserHandlerIPLM) Create(ctx *fiber.Ctx) error {
	reqUser := CreateUser{}
	if err := ctx.BodyParser(&reqUser); err != nil {
		return WebResponseJosn(ctx, 404, "Bad Request", err.Error())
	}

	resp := h.UserService.Create(context.Background(), ctx, reqUser)
	return WebResponseJosn(ctx, 200, "Ok", resp)
}
func (h *UserHandlerIPLM) GetByEmail(ctx *fiber.Ctx) error {
	userEmail := ctx.Params("email")

	resp := h.UserService.GetByEmail(context.Background(), ctx, userEmail)
	return WebResponseJosn(ctx, 200, "Ok", resp)
}
func (h *UserHandlerIPLM) GetByID(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")

	resp := h.UserService.GetByEmail(context.Background(), ctx, userID)
	return WebResponseJosn(ctx, 200, "Ok", resp)
}
func (h *UserHandlerIPLM) GetAll(ctx *fiber.Ctx) error {
	resp := h.UserService.GetAll(context.Background(), ctx)
	return WebResponseJosn(ctx, 200, "Ok", resp)
}
