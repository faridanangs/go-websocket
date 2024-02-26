package ws

import (
	"context"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func (usr *WsUser) writeMessage(hub *Hub) {
	defer func() {
		usr.Conn.Close()
	}()

	for {
		msg, ok := <-usr.Message
		if !ok {
			return
		}
		hub.RoomServiceIPLM.CreateMessage(context.Background(), &fiber.Ctx{}, msg)
		err := usr.Conn.WriteJSON(msg)
		if err != nil {
			log.Printf("Error writing JSON message to connection: %v", err)
			return
		}

	}

}
func (usr *WsUser) readMessage(h *Hub) {
	defer func() {
		h.UnRegister <- usr
		usr.Conn.Close()
	}()

	for {
		_, m, err := usr.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseGoingAway) {
				log.Printf("error reading message: %v", err)
			}
			break
		}

		msg := &Message{
			ID:        usr.ID,
			Username:  usr.Name,
			Content:   string(m),
			RoomID:    usr.RoomID,
			CreatedAt: usr.CreatedAt,
		}
		h.Broadcast <- msg
	}

}
