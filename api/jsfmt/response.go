package jsfmt

import (
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/model"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Error   interface{} `json:"error"`
	Result  interface{} `json:"result"`
	Success bool        `json:"success"`
}

func NormalResponse(result interface{}) *Response {
	return &Response{
		Error:   nil,
		Success: true,
		Result:  result,
	}
}

func ErrorResponse(zinError *errutils.ZinError) *Response {
	return &Response{
		Error: Error{
			Code:    zinError.Errno,
			Message: zinError.Msg,
		},
		Success: false,
		Result:  nil,
	}
}

type LoginResponse struct {
	User  *model.User  `json:"user"`
	Token *model.Token `json:"token"`
}
