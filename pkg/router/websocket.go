package router

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func WebSocketRoutes(app *fiber.App) {
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		defer c.Close()
		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				break
			}
			// Echo the received message back to the client
			if err = c.WriteMessage(mt, msg); err != nil {
				break
			}
		}
	}))
}
