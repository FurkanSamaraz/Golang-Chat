package controllers

import (
	"fmt"

	api_model "github.com/FurkanSamaraz/Golang-Chat/internal/pkg/model"

	api_structure "github.com/FurkanSamaraz/Golang-Chat/internal/pkg/structures"

	"github.com/gofiber/websocket/v2"
)

var clients = make(map[string]*api_structure.Client)
var rooms = make(map[string]*api_structure.Room)

// WsHandler handles WebSocket connections and message sending
// @Summary       WebSocket Handler
// @Description   Handles WebSocket connections and message sending
// @Tags          WebSocket
// @Router        /ws [get]
func WsHandler(c *websocket.Conn) {
	user := c.Locals("user").(*api_structure.Claims)

	// Kullanıcı listesine kullanıcı ekle
	client := &api_structure.Client{Conn: c, Username: user.Name}
	clients[user.Name] = client
	fmt.Println("clients", len(clients), clients, c.RemoteAddr())

	// Gelen yeni mesajlar için süresiz dinle
	// Tanımladığımız WebSocket noktamız üzerinden
	for {
		var msg api_structure.Chat
		err := client.Conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("error reading json", err)
			break
		}

		fmt.Println("received message", msg)

		// Mesajın bir hedef odası olup olmadığını kontrol edin
		if msg.Target != nil {
			// Odayı al veya oluştur
			roomID := msg.Target.ID
			room, exists := rooms[roomID]
			if !exists {
				room = &api_structure.Room{
					ID:      roomID,
					Name:    msg.Target.Name,
					Clients: make(map[string]*api_structure.Client),
				}
				rooms[roomID] = room
			}

			// Odaya kullanıcıyı ekle
			room.Clients[client.Username] = client

			// Odadaki kullanıcılara mesajı yayınlayın
			for _, c := range room.Clients {
				if c.Username != client.Username {
					err = c.Conn.WriteJSON(msg)
					if err != nil {
						fmt.Println("error writing json", err)
						break
					}

					// Gönderici ve alıcıyı birbirinin kişi listelerine ekleyin
					api_model.AddToContactList(c.Username, client.Username)
					api_model.AddToContactList(client.Username, c.Username)
				}
			}

			// Mesajı Redis'e kaydet
			api_model.SaveChatHistory(msg)

		} else {
			// Tüm bağlı istemcilere mesaj yayınlayın (Broadcast)
			for _, c := range clients {
				if c.Username != client.Username {
					err = c.Conn.WriteJSON(msg)
					if err != nil {
						fmt.Println("error writing json", err)
						break
					}

					// Gönderici ve alıcıyı birbirinin kişi listelerine ekleyin
					api_model.AddToContactList(c.Username, client.Username)
					api_model.AddToContactList(client.Username, c.Username)
				}
			}

			// Mesajı Redis'e kaydet
			api_model.SaveChatHistory(msg)
		}
	}

	fmt.Println("exiting", c.RemoteAddr().String())
	delete(clients, user.Name)
}
