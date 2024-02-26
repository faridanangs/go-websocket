package ws

import (
	"context"
	"database/sql"
	"time"
	"ws_postgre/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RoomServiceIPLM struct {
	DB                 *sql.DB
	RoomRepositoryIPLM *RoomRepositoryIPLM
}

func NewRoomService(db *sql.DB, roomRepository *RoomRepositoryIPLM) *RoomServiceIPLM {
	return &RoomServiceIPLM{
		DB:                 db,
		RoomRepositoryIPLM: roomRepository,
	}
}

func (service *RoomServiceIPLM) Create(ctx context.Context, fctx *fiber.Ctx, req CreateRoom) Room {
	tx, err := service.DB.Begin()
	helper.HelperError(err, "Error creating room service database transaction")
	defer helper.DBTransaction(tx)

	user := Room{
		ID:        uuid.NewString(),
		Name:      req.Name,
		UserID:    req.UserID,
		CreatedAt: time.Now(),
	}

	resp := service.RoomRepositoryIPLM.Save(ctx, tx, user)
	return RoomResponseWeb(resp)

}
func (service *RoomServiceIPLM) GetByID(ctx context.Context, fctx *fiber.Ctx, roomID string) Room {
	tx, err := service.DB.Begin()
	helper.HelperError(err, "Error GetbyID room service database transaction")
	defer helper.DBTransaction(tx)

	resp, err := service.RoomRepositoryIPLM.GetByID(ctx, tx, roomID)
	if err != nil {
		fctx.Status(404).JSON(fiber.Map{"err": err.Error(), "code": 404, "status": fiber.StatusNotFound})
	}

	return RoomResponseWeb(resp)
}
func (service *RoomServiceIPLM) GetAll(ctx context.Context, fctx *fiber.Ctx) []Room {
	tx, err := service.DB.Begin()
	helper.HelperError(err, "Error getall room service database transaction")
	defer helper.DBTransaction(tx)

	respponses := service.RoomRepositoryIPLM.GetAll(ctx, tx)

	return RoomResponsesWeb(respponses)

}

func (service *RoomServiceIPLM) CreateMessage(ctx context.Context, fctx *fiber.Ctx, req *Message) {
	tx, err := service.DB.Begin()
	helper.HelperError(err, "Error creating message service database transaction")
	defer helper.DBTransaction(tx)

	message := Message{
		Username:  req.Username,
		Content:   req.Content,
		ID:        req.ID,
		RoomID:    req.RoomID,
		CreatedAt: time.Now(),
	}
	service.RoomRepositoryIPLM.SaveMessage(ctx, tx, message)

}
