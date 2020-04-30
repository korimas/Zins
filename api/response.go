package api

import "github.com/zpdev/zins/common/errutils"

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
