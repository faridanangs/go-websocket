package ws

import (
	"time"
	internaluser "ws_postgre/internals/internal_user"

	"github.com/gofiber/fiber/v2"
)

func RoomResponseWeb(req Room) Room {
	return Room{
		ID:        req.ID,
		Name:      req.Name,
		UserID:    req.UserID,
		CreatedAt: req.CreatedAt,
		// User:      RoomUserResponse(req.User),
		User: internaluser.UserResponse{
			ID:        req.UserID,
			Name:      req.User.Name,
			Email:     req.User.Email,
			CreatedAt: req.CreatedAt,
		},
	}
}

// func RoomUserResponse(req []internaluser.UserResponse) []internaluser.UserResponse {
// 	var users []internaluser.UserResponse
// 	for _, user := range req {
// 		users = append(users, internaluser.UserResponse{
// 			ID:        user.ID,
// 			Name:      user.Name,
// 			Email:     user.Email,
// 			CreatedAt: user.CreatedAt,
// 		})
// 	}

// 	return users
// }

func RoomResponsesWeb(req []Room) []Room {
	var responses []Room
	for _, u := range req {
		responses = append(responses, RoomResponseWeb(u))
	}

	return responses
}

func WebResponseJosn(ctx *fiber.Ctx, code int, sttus string, data any) error {
	return ctx.Status(code).JSON(internaluser.WebResponse{
		Code:   code,
		Status: sttus,
		Data:   data,
	})
}

func MessageResponseWeb(req Message) Message {
	return Message{
		ID:        req.ID,
		Username:  req.Username,
		Content:   req.Content,
		RoomID:    req.RoomID,
		CreatedAt: time.Now(),
	}
}
