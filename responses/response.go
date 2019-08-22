package responses

type Item map[string]interface{}

type ItemResponse struct {
	Code    int `json:"code"`
	Message string `json:"message"`
	Data    Item `json:"data"`
}
