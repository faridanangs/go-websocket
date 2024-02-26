package ws

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type RoomHandlerIPLM struct {
	RoomService *RoomServiceIPLM
	hub         *Hub
}

func NewRoomHandler(roomService *RoomServiceIPLM, hub *Hub) *RoomHandlerIPLM {
	return &RoomHandlerIPLM{
		RoomService: roomService,
		hub:         hub,
	}
}

func (h *RoomHandlerIPLM) Create(ctx *fiber.Ctx) error {
	reqUser := CreateRoom{}
	if err := ctx.BodyParser(&reqUser); err != nil {
		return WebResponseJosn(ctx, 404, "Bad Request", err.Error())
	}
	resp := h.RoomService.Create(context.Background(), ctx, reqUser)

	h.hub.Rooms[resp.ID] = &RoomHub{
		ID:        resp.ID,
		Name:      resp.Name,
		User:      make(map[string]*WsUser),
		CreatedAt: resp.CreatedAt,
	}

	return WebResponseJosn(ctx, 200, "Ok", resp)
}

func (h *RoomHandlerIPLM) GetByID(ctx *fiber.Ctx) error {
	roomID := ctx.Params("roomId")

	resp := h.RoomService.GetByID(context.Background(), ctx, roomID)
	return WebResponseJosn(ctx, 200, "Ok", resp)
}

func (h *RoomHandlerIPLM) GetAll(ctx *fiber.Ctx) error {
	resp := h.RoomService.GetAll(context.Background(), ctx)
	return WebResponseJosn(ctx, 200, "Ok", resp)
}

func (h *RoomHandlerIPLM) JoinRoom(conn *websocket.Conn, ctx *fiber.Ctx) {
	roomID := conn.Params("roomId")
	userID := conn.Query("userId")
	userName := conn.Query("username")

	resp := h.RoomService.GetByID(context.Background(), ctx, roomID)
	if resp.ID != roomID {
		WebResponseJosn(ctx, 404, "Room not found", nil)
		return
	}
	fmt.Print(h.hub.Rooms[roomID])
	if _, ok := h.hub.Rooms[roomID]; !ok {
		h.hub.Rooms[roomID] = &RoomHub{
			ID:        resp.ID,
			Name:      resp.Name,
			User:      make(map[string]*WsUser),
			CreatedAt: resp.CreatedAt,
		}
	}

	wsusr := &WsUser{
		ID:        userID,
		Name:      userName,
		Message:   make(chan *Message, 10),
		CreatedAt: time.Now().Local(),
		RoomID:    roomID,
		Conn:      conn,
	}

	msg := &Message{
		ID:        userID,
		Username:  userName,
		Content:   userName + " barusaja memasuki ruang obrolan",
		RoomID:    roomID,
		CreatedAt: time.Now().Local(),
	}

	h.hub.Register <- wsusr
	h.hub.Broadcast <- msg

	go wsusr.writeMessage(h.hub)
	wsusr.readMessage(h.hub)

}
