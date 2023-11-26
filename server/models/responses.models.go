package models

type SuccessResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

type SuccessResponseWithData struct {
	Success bool `json:"success"`
	Data interface{} `json:"data"`
}

type ErrResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	ErrorDetails interface{} `json:"error_details"`
}

type LoginResponse struct {
	Success bool `json:"success"`
	Token string `json:"token"`
	Data interface{} `json:"data"`
}