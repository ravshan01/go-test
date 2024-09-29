package core

type Response struct {
	Data  interface{}   `json:"data"`
	Error ResponseError `json:"error,omitempty"`
}

type ResponseError struct {
	Message string `json:"message"`
}

type BadRequestResponse struct {
	Response
}
