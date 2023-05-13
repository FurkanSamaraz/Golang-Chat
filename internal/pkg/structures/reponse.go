package structures

// Chat Mesajı
type StatusMessage struct {
	Message string `json:"message"`
}

func (p *StatusMessage) TableName() string {
	return "calendar.users"
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}
