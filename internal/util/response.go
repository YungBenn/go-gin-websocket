package util

type dataResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type messageResponse struct {
	Message string `json:"message"`
}

func DataResponse(message string, data interface{}) dataResponse {
	response := dataResponse{
		Message: message,
		Data: data,
	}

	return response
}

func MessageResponse(message string) messageResponse {
	response := messageResponse{
		Message: message,
	}

	return response
}