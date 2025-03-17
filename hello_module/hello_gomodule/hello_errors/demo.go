package hello_errors_test

import "fmt"

// 自定义error(实现builtin/error接口)
//
//	type error interface {
//		Error() string
//	}
type CustomError struct {
	Code int64
	Msg  string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}

func Foo() (string, error) {
	return "", &CustomError{Code: 100, Msg: "failed to xxx"}
}
