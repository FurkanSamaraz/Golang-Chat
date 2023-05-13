package structures

type Chat struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
	Target  *Room  `json:"target"`
}
