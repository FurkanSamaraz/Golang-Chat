package services

import "main/internal/pkg/structures"

// Kullanıcıyı Odaya Ekle
func AddClientToRoom(room *structures.Room, username string, client *structures.Client) {
	room.AddClientToRoom(username, client)
}

// Kullanıcıyı Odadan Çıkar
func RemoveClientFromRoom(room *structures.Room, username string) {
	room.RemoveClientFromRoom(username)
}
