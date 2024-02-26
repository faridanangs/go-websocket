package main

import (
	"ws_postgre/db"
	"ws_postgre/internals/ws"
	"ws_postgre/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := db.ConnectionToDatabase()
	roomRepo := ws.NewRoomRepositoryIPLM()

	hub := ws.NewHub(db, roomRepo)
	router.InitialRouter(db, app, hub)
	go hub.Run()

	app.Listen(":8000")
}
