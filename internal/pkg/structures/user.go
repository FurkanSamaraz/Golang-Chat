package structures

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (p *User) TableName() string {
	return "calendar.users"
}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Jwt     string      `json:"jwt"`
	Data    interface{} `json:"data"`
}
