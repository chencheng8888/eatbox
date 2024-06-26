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
		panic(fmt.Sprintf("错误码%d已存在，请更换一个", code))

	}
	codes[code] = msg
	return &Error{
		code: code,
		msg:  msg,
	}
}
func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d,错误信息: %s", e.Code(), e.Msg())
}
func (e *Error) Code() int {
	return e.code
}
func (e *Error) Msg() string {
	return e.msg
}
func (e *Error) Details() []string {
	return e.details
}
func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}
func (e *Error) WithDetails(details ...string) *Error {
	e.details = []string{}
	for _, d := range details {
		e.details = append(e.details, d)
	}
	return e
}
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case NotFound.Code():
		return http.StatusNoContent
	case ToJSONError.Code():
		return http.StatusInternalServerError
	case ErrRedisSet.Code():
		return http.StatusInternalServerError
	case ErrRedisDel.Code():
		return http.StatusInternalServerError
	case JSONUnmarshalError.Code():
		return http.StatusInternalServerError
	case MySQLErr.Code():
		return http.StatusInternalServerError
	case WXAPIError.Code():
		return http.StatusInternalServerError
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case PageInvalid.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedAuthIsEmpty.Code():
		fallthrough
	case Unauthorized.Code():
		fallthrough
	case UnauthorizedTokenExpired.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError

}
