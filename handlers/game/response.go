package game

import "github.com/zuno90/go-ws/utils"

type Response struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message,omitempty" default:""`
	Data    any                    `json:"data,omitempty"`
	Error   []*utils.ErrorResponse `json:"errors,omitempty"`
}

func Resp(s bool, m string, d any, e []*utils.ErrorResponse) *Response {
	return &Response{
		Success: s,
		Message: m,
		Data:    d,
		Error:   e,
	}
}
