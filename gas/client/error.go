package client

import "fmt"

// CodeSuccess 是请求正确响应时给出的 code 值
const CodeSuccess = 0

// APIError 是接口本身报错时给出的 error 类型
type APIError struct {
	Code    int
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// Ensure 检查 code & message 并在响应不正确时构造错误实例
func Ensure(code int, message string) error {
	if code != CodeSuccess {
		return &APIError{code, message}
	}
	return nil
}
