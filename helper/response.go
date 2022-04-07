package helper

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct {
}

func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{}
}
