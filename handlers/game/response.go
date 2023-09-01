package game

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type Response struct {
	Success bool             `json:"success"`
	Message string           `json:"message,omitempty" default:""`
	Data    any              `json:"data,omitempty"`
	Error   []*ErrorResponse `json:"errors,omitempty" default:""`
}

func Resp(s bool, m string, d any, e []*ErrorResponse) *Response {
	return &Response{
		Success: s,
		Message: m,
		Data:    d,
		Error:   e,
	}
}
