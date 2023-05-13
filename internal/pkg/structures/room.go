package structures

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"-"`
}

// Yeni Oda Oluştur
func (r *Room) AddClientToRoom(username string, client *Client) {
	r.Clients[username] = client
}

// Odayı Sil
func (r *Room) RemoveClientFromRoom(username string) {
	delete(r.Clients, username)
}
