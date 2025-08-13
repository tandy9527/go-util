package resp

import "github.com/tandy9527/go-util/errno"

// Response 通用返回结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// 成功返回
func Success(data interface{}) Response {
	return Response{
		Code: errno.Success.Code,
		Msg:  errno.Success.Msg,
		Data: data,
	}
}

// 错误返回
func Error(e errno.Errno, data interface{}) Response {
	return Response{
		Code: e.Code,
		Msg:  e.Msg,
		Data: data,
	}
}
