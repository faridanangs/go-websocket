package internaluser

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"ws_postgre/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserServiceIPLM struct {
	DB                 *sql.DB
	UserRepositoryIPLM *UserRepositoryIPLM
}

func NewUserService(db *sql.DB, userRepository *UserRepositoryIPLM) *UserServiceIPLM {
	return &UserServiceIPLM{
		DB:                 db,
		UserRepositoryIPLM: userRepository,
	}
}

func (service *UserServiceIPLM) Create(ctx context.Context, fctx *fiber.Ctx, req CreateUser) UserResponse {
	tx, err := service.DB.Begin()
	helper.HelperError(err, "Error creating user service database transaction")
	defer helper.DBTransaction(tx)

	fmt.Print("req", req)

	user := User{
		ID:        uuid.NewString(),
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
		Password:  req.Password,
	}

	resp := service.UserRepositoryIPLM.Save(ctx, tx, user)
	fmt.Print("resp", resp)
	return UserResponseWeb(resp)

}
func (service *UserServiceIPLM) GetByEmail(ctx context.Context, fctx *fiber.Ctx, userEmail string) UserResponse {
	tx, err := service.DB.Begin()
	helper.HelperError(err, "Error Getbyemail user service database transaction")
	defer helper.DBTransaction(tx)

	resp, err := service.UserRepositoryIPLM.GetByEmail(ctx, tx, userEmail)
	if err != nil {
		fctx.Status(404).JSON(fiber.Map{"err": err.Error(), "code": 404, "status": fiber.StatusNotFound})
	}

	return UserResponseWeb(resp)
}
func (service *UserServiceIPLM) GetByID(ctx context.Context, fctx *fiber.Ctx, userID string) UserResponse {
	tx, err := service.DB.Begin()
	helper.HelperError(err, "Error GetbyID user service database transaction")
	defer helper.DBTransaction(tx)

	resp, err := service.UserRepositoryIPLM.GetByID(ctx, tx, userID)
	if err != nil {
		fctx.Status(404).JSON(fiber.Map{"err": err.Error(), "code": 404, "status": fiber.StatusNotFound})
	}

	return UserResponseWeb(resp)
}
func (service *UserServiceIPLM) GetAll(ctx context.Context, fctx *fiber.Ctx) []UserResponse {
	tx, err := service.DB.Begin()
	helper.HelperError(err, "Error getall user service database transaction")
	defer helper.DBTransaction(tx)

	respponses := service.UserRepositoryIPLM.GetAll(ctx, tx)

	return UserResponsesWeb(respponses)

}
