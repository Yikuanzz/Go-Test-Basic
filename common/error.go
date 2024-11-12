package common

import "net/http"

// errors 重写错误，实现 Error 方法，表示带有 HTTP 状态码的错误。
type errors int

var (
	// ErrNotFound 在资源未找到时返回。
	ErrNotFound errors = http.StatusNotFound
	// ErrHasExists 在资源已存在时返回。
	ErrHasExists errors = http.StatusConflict
	// ErrInternal 在内部服务器错误时返回。
	ErrInternal errors = http.StatusInternalServerError
	// ErrBadRequest 在请求错误时返回。
	ErrorBadRequest errors = http.StatusBadRequest
)

// Error 返回与HTTP状态码关联的错误消息。
func (e errors) Error() string {
	return http.StatusText(int(e))
}

func (e errors) StatusCode() int {
	return int(e)
}

func StatusCode(err error) int {
	return err.(errors).StatusCode()
}
