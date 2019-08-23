package requests

type RegisterRequest struct {
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewRegisterRequest() *RegisterRequest {
	return &RegisterRequest{}
}
