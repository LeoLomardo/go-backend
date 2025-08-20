package user

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"` // O - ta ai pra  q senha nao seja enviada em respostas JSON
}
