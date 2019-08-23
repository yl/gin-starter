package requests

type LoginRequest struct {
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}
