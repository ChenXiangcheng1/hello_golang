package hello_errors_test

import (
	"errors"
	"fmt"
	"testing"
)

/*
golang error和exception区别:
Errors as value
Monadic 结构化 层层unwarp (无隐藏控制流)
强制用户错误检查if-else 使用错误链
缺少一个语法糖单行错误处理

Exceptions
特殊隐藏控制流 栈展开走特殊控制流比return性能差
不需要次次错误检查 含栈信息
*/

/*
errors历史:
包装: fmt.Errorf()
解包: strings.Container(err.Error()) 判断具体错误类型
判断具体错误类型: err.(*CustomError)使用.类型断言

go1.13 对错误的处理变为对错误链的处理
结构化包装: fmt.Errorf("%w", err) -> *fmt.wrapError
解包: errors.Unwrap(err) -> err
判断具体错误类型: errors.Is(err, target) -> bool、errors.As(err, &target) -> bool

go1.20
结构化包装: errors.Join(err, err) -> *errors.joinError
*/

// errors
//
//	type errorString struct {
//		s string
//	}
//
//	type wrapError struct {
//		msg string
//		err error
//	}
//
// fmt  // 包装/解包
// fmt.Errorf(fmtStr, ...any) -> *errors.errorString  // 和 errors.New()一样
// fmt.Errorf(fmtStr(%w), ...any) -> *fmt.wrapError
func TestMain(t *testing.T) {
	err1 := errors.New("failed_to_aaa")
	fmt.Printf("%T %q\n", err1, err1) // *errors.errorString
	err2 := fmt.Errorf("%v", "failed_to_bbb")
	fmt.Printf("%T %q\n", err2, err2) // *errors.errorString

	err3 := fmt.Errorf("%w: %v", err1, err2) // *fmt.wrapError  // %w  // 错误链
	fmt.Printf("%T %q\n", err3, err3)
	err4 := errors.Join(err1, err2) // *errors.joinError "err1\nerr2"  // \n分割  // 比fmt.Errorf("%w\n%w", err1, err2)快75%
	fmt.Printf("%T %q\n", err4, err4)

	// errors.Is(err, target) -> bool
	// errors.As(err, &target) -> bool
	fmt.Printf("%v\n", errors.Is(err3, err3)) // true
	fmt.Printf("%v\n", errors.Is(err3, err1)) // true
	fmt.Printf("%v\n", errors.Is(err1, err3)) // false

	ret := errors.Unwrap(err3)
	fmt.Printf("%T %q\n", ret, ret)
	ret = errors.Unwrap(err4)
	fmt.Printf("%T %q\n", ret, ret)
}

// 自定义error(实现builtin/error接口)
//
//	type error interface {
//		Error() string
//	}
func TestCustomError(t *testing.T) {
	ret, err := Foo()
	if v, ok := err.(*CustomError); ok { // 类型断言
		fmt.Printf("%T %v %d %s\n", v, v, v.Code, v.Msg)
	} else if err != nil {
		t.Errorf("%T %v", err, err)
		// return nil, fmt.Errorf("xxx error: %v", err)
	} else {
		fmt.Printf("%T %v\n", ret, ret)
	}
}
