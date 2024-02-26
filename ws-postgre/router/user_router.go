package router

import (
	"database/sql"
	"time"
	internaluser "ws_postgre/internals/internal_user"
	"ws_postgre/internals/ws"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func InitialRouter(db *sql.DB, app *fiber.App, hub *ws.Hub) {
	InitialRouterUser(db, app)
	InitialRouterRoom(db, app, hub)

}
func InitialRouterRoom(db *sql.DB, app *fiber.App, hub *ws.Hub) {
	roomRepo := ws.NewRoomRepositoryIPLM()
	roomService := ws.NewRoomService(db, roomRepo)
	roomHandler := ws.NewRoomHandler(roomService, hub)

	usr := app.Group("/api/room")
	usr.Post("/create", roomHandler.Create)
	usr.Get("/:roomId", roomHandler.GetByID)
	usr.Get("/", roomHandler.GetAll)
	usr.Get("/join-room/:roomId", websocket.New(func(c *websocket.Conn) {
		roomHandler.JoinRoom(c, &fiber.Ctx{})
	}, websocket.Config{
		Filter: func(c *fiber.Ctx) bool {
			return true
		},
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: time.Hour * 24,
	}))
}
func InitialRouterUser(db *sql.DB, app *fiber.App) {
	userRepo := internaluser.NewUserRepositoryIPLM()
	userService := internaluser.NewUserService(db, userRepo)
	userHandler := internaluser.NewUserHandler(userService)

	usr := app.Group("/api/user")
	usr.Post("/create", userHandler.Create)
	usr.Get("/:email", userHandler.GetByEmail)
	usr.Get("/:id", userHandler.GetByID)
	usr.Get("/", userHandler.GetAll)
}
