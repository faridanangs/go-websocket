package ws

import (
	"time"
	internaluser "ws_postgre/internals/internal_user"

	"github.com/gofiber/contrib/websocket"
)

type Room struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	UserID    string `json:"user_id"`
	User      internaluser.UserResponse
	CreatedAt time.Time `json:"created_at"`
}

type CreateRoom struct {
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

type Message struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	RoomID    string    `json:"room_id"`
	CreatedAt time.Time `json:"created_at"`
}

type MessageResp struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	RoomID    string    `json:"room_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type WsUser struct {
	ID        string
	Name      string
	Message   chan *Message
	CreatedAt time.Time
	RoomID    string
	Conn      *websocket.Conn
}
