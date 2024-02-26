package ws

import (
	"database/sql"
	"time"
)

type RoomHub struct {
	ID        string
	Name      string
	User      map[string]*WsUser
	CreatedAt time.Time
}

type Hub struct {
	Rooms           map[string]*RoomHub
	Register        chan *WsUser
	UnRegister      chan *WsUser
	Broadcast       chan *Message
	RoomServiceIPLM *RoomServiceIPLM
}

func NewHub(db *sql.DB, roomRepo *RoomRepositoryIPLM) *Hub {
	return &Hub{
		Rooms:           make(map[string]*RoomHub),
		Register:        make(chan *WsUser),
		UnRegister:      make(chan *WsUser),
		Broadcast:       make(chan *Message, 5),
		RoomServiceIPLM: NewRoomService(db, roomRepo),
	}
}

func (h *Hub) Run() {

	for {
		select {
		case usr := <-h.Register:
			if _, ok := h.Rooms[usr.RoomID]; ok {
				if _, ok := h.Rooms[usr.RoomID].User[usr.ID]; !ok {
					h.Rooms[usr.RoomID].User[usr.ID] = usr
				}
			}
		case usr := <-h.UnRegister:
			if _, ok := h.Rooms[usr.RoomID]; ok {
				if _, ok := h.Rooms[usr.RoomID].User[usr.ID]; ok {
					if len(h.Rooms[usr.RoomID].User) != 0 {
						h.Broadcast <- &Message{
							ID:        usr.ID,
							Username:  usr.Name,
							Content:   usr.Name + " barus saja meninggalkan ruang obrolan",
							RoomID:    usr.RoomID,
							CreatedAt: time.Now(),
						}
					}
					delete(h.Rooms[usr.RoomID].User, usr.ID)
					close(usr.Message)
				}
			}
		case msg := <-h.Broadcast:
			if _, ok := h.Rooms[msg.RoomID]; ok {
				for _, r := range h.Rooms[msg.RoomID].User {
					r.Message <- msg
				}
			}
		}
	}
}
