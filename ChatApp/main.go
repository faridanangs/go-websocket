package main

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Body       string
	Connection *websocket.Conn
}

var channels = make(map[string]*[]*websocket.Conn)
var register = make(chan *websocket.Conn)
var textpipe = make(chan *Message)
var unregister = make(chan *websocket.Conn)

func socketListen() {
	for {
		select {
		case connection := <-register:
			(*channels[connection.Params("id")]) = append((*channels[connection.Params("id")]), connection)
		case message := <-textpipe:
			for _, connection := range *channels[message.Connection.Params("id")] {
				connection.WriteMessage(websocket.TextMessage, []byte(message.Body))
			}
		case connection := <-unregister:
			for i := len(*channels[connection.Params("id")]) - 1; i >= 0; i-- {
				if (*channels[connection.Params("id")])[i] == connection {
					(*channels[connection.Params("id")]) = append(
						(*channels[connection.Params("id")]),
						(*channels[connection.Params("id")])[i+1:]...)
				}
			}
			// jiks tidak ada in chatroom, delete chatroom section
			if len(*channels[connection.Params("id")]) == 0 {
				delete(channels, connection.Params("id"))
			}
		}
	}
}
func main() {
	app := fiber.New()
	go socketListen()

	app.Get("/ws/chat/:id/ws", websocket.New(func(c *websocket.Conn) {
		defer func() {
			// ketika membaca pesan gagal, server akn mencoba untuk menolak user dan  chanel
			unregister <- c
		}()

		// mencoba untuk berkoleksi dengan user koneksi
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				return
			}

			// setelah membaca pesan, ciba unutk mengirim pesan ke orang lain di chatroom
			textpipe <- &Message{
				Body:       string(message),
				Connection: c,
			}
		}

	}))

	app.Listen(":3000")
}
