package structs

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func Resp(s bool, m string, d any) *Response {
	return &Response{
		Success: s,
		Message: m,
		Data:    d,
	}
}
