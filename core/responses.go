package core

type Response[T any] struct {
	Data  *T             `json:"data"`
	Error *ResponseError `json:"error,omitempty"`
}

type ResponseError struct {
	Message string `json:"message"`
}

func NewBadRequestResponse(message string) Response[interface{}] {
	return Response[interface{}]{
		Error: &ResponseError{Message: message},
	}
}
