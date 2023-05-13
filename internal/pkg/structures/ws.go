package structures

import "github.com/gofiber/websocket/v2"

type Client struct {
	Conn     *websocket.Conn
	Username string `json:"username"`
	RoomID   string `json:"-"`
}
