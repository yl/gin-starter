package responses

import (
	"go-trading/models"
	"net/http"
)

func UserResponse(user *models.User) ItemResponse {
	data := make(Item)
	data["mobile"] = user.Mobile

	return ItemResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}
}
