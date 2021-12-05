package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {

	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码: %d 已经存在,请更换", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d,错误信息:%s", e.Code(), e.Msg())
}

// 返回错误Code
func (e *Error) Code() int {
	return e.code
}

// 返回错误内容
func (e *Error) Msg() string {
	return e.msg
}

// 返回错误内容
func (e *Error) Details() []string {
	return e.details
}

// 包括描述信息
func (e *Error) WithDetails(details ...string) *Error {

	newError := *e
	newError.details = []string{}
	newError.details = append(newError.details, details...)
	return &newError
}

// 返回错误的状态码
func (e *Error) StatusCode() int {

	switch e.Code() {
	case Success.code:
		return http.StatusOK
	case ServerError.code:
		return http.StatusInternalServerError
	case NotFound.code:
		return http.StatusNotFound
	case InvalidParams.code:
		return http.StatusBadRequest
	case UnauthorizedAuthNoExist.code:
		fallthrough
	case UnauthorizedTokenGenerate.code:
		fallthrough
	case UnauthorizedTokenError.code:
		fallthrough
	case UnauthorizedTokenTimeOut.code:
		return http.StatusUnauthorized
	case ToolManyRequests.code:
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
