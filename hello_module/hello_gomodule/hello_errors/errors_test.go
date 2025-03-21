package hello_errors_test

import (
	"errors"
	"fmt"
	"testing"
)

/*
golang error和exception区别:
Errors as value, golang error被看作一等公民(值)
Monadic 结构化 层层unwarp  好处: 无隐藏控制流
完全控制error错误链(没强制用户处理)
	强制用户错误检查if-else 好处: 没有uncaught exception
	没强制用户处理 缺点: 不如union type严格约束, 多返回值可能存在歧义, 可能上游没处理error
缺少一个语法糖单行错误处理  // golang混淆了BUG(编程错误)和异常(业务逻辑的一部分)，在可能BUG处也要求设置检查点

Exceptions
特殊隐藏控制流 需要进行stack trace堆栈跟踪 当需要时pop比return性能差
不需要次次错误检查 含栈信息
没强制用户错误检查(懒惰的编程习惯 一直冒泡到主程序) 没强制用户处理

主要区别在于：值 控制流
*/

/*
errors历史:
包装: fmt.Errorf(fmtStr, ...any) -> *errors.errorString
解包: strings.Container(err.Error()) 判断具体错误类型
判断具体错误类型: err.(*CustomError)使用.类型断言

go1.13 对处理error变为处理error链表
结构化包装(嵌套 本质是链表): fmt.Errorf("%w", err) -> *fmt.wrapError
解包: errors.Unwrap(err) -> err
判断具体错误类型: errors.Is(err, target) -> bool、errors.As(err, &target) -> bool

go1.20
结构化包装(组合): errors.Join(err, err) -> *errors.joinError
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
//	type joinError struct {
//		errs []error
//	}
//
// fmt  // 包装/解包
// fmt.Errorf(fmtStr, ...any) -> *errors.errorString  // 和 errors.New()一样
// fmt.Errorf(fmtStr(%w), ...any) -> *fmt.wrapError
func TestMain(t *testing.T) {
	err1 := errors.New("failed_to_aaa")       // 本质是调用errorString
	fmt.Printf("%T %q\n", err1, err1)         // *errors.errorString
	err2 := fmt.Errorf("%v", "failed_to_bbb") // 本质是调用wrapError
	fmt.Printf("%T %q\n", err2, err2)         // *errors.errorString

	err3 := fmt.Errorf("%w: %v", err1, err2) // *fmt.wrapError  // %w  // 错误链
	fmt.Printf("%T %q\n", err3, err3)
	ret := errors.Unwrap(err3)
	fmt.Printf("%T %q\n", ret, ret)
	err4 := errors.Join(err1, err2) // *errors.joinError "err1\nerr2"  // \n分割  // 比fmt.Errorf("%w\n%w", err1, err2)性能提升75%
	fmt.Printf("%T %q\n", err4, err4)
	ret = errors.Unwrap(err4)
	fmt.Printf("%T %q\n", ret, ret)

	// errors.Is(err, target) -> bool
	// errors.As(err, &target) -> bool
	fmt.Printf("%v\n", errors.Is(err3, err3)) // true
	fmt.Printf("%v\n", errors.Is(err3, err1)) // true
	fmt.Printf("%v\n", errors.Is(err1, err3)) // false

	em := &CustomError{}
	if errors.As(err3, &em) { // 二级指针(pointer to type)  // 提取target
		fmt.Printf("%T %v\n", em, em)
	} else {
		fmt.Printf("???\n")
	}
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
		return
	} // if err直接return，避免if-else嵌套
	fmt.Printf("%T %v\n", ret, ret)
}

func must(v interface{}, err error) interface{} {
	if err != nil {
		panic(fmt.Sprintf("unexpected error: %v", err))
	}
	return v
}

func dosomething() (int, error) {
	return 0, nil
	// return -1, errors.New("failed to dosomething")
}

func TestBug(t *testing.T) {
	ret := must(dosomething()) // 无需if-else
	fmt.Printf("%T %v\n", ret, ret)
}
