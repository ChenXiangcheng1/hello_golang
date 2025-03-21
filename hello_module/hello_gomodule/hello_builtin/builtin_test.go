package hello_builtin_test

import (
	"fmt"
	"testing"
)

/*
G()->F()->panic()
F_goroutine终止并执行F_defer, G_goroutine终止并执行G_defer,
直到所有执行goroutine的函数终止, 程序以非零exit码终止
该goroutine称为panic goroutine 该终止序列称为panicking sequence, 可以通过recover()进行控制
func panic(v any)
*/

/*
在F.defer(err_val)->recover() (在延迟函数中 而不是在延迟函数中再调用的函数中 执行调用recover)  // 一定是在F.defer()中调用
停止panicking sequence(恢复正常执行), 并返回error value
recover()-> err_val | nil(F.defer之外、当前goroutine无panic)
func recover() any
*/

func TestPanic(t *testing.T) {
	r1 := recover()
	fmt.Printf("%T %v\n", r1, r1) // nil

	defer func() {
		hello_recover()
		if r := recover(); r != nil {
			fmt.Printf("%T %v\n", r, r) // string
		}
	}()

	panic(struct {
		Code int
		Msg  string
	}{9001, "failed to xxx"})
}

func hello_recover() {
	fmt.Println("hello_recover")
	// if r := recover(); r != nil {  // unreachable
	// 	fmt.Println("hello_recover: return err_val")
	// } else {
	// 	fmt.Println("hello_recover: return nil")
	// }
}

/*
嵌套panic
panic: failed to yyy
panic: failed to xxx [recovered]
panic: failed to xxx
*/
func TestNestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%T %v\n", r, r) // xxx
		}
		if r := recover(); r != nil {
			fmt.Printf("%T %v\n", r, r)
		} else {
			fmt.Println("nil")
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%T %v\n", r, r) // yyy
		} else {
			fmt.Println("nil")
		}
		panic("failed to xxx") // NestPanic
	}()

	panic("failed to yyy")
}

func TestNestfunctionWithPanic(t *testing.T) {
	defer func() {
		fmt.Println("TestNestPanic")
		if r := recover(); r != nil {
			fmt.Printf("%T %v\n", r, r)
		} else {
			fmt.Println("nil")
		}
	}()

	hello_panic()
}

func hello_panic() {
	defer func() {
		fmt.Println("hello_panic")
		if r := recover(); r != nil {
			fmt.Printf("%T %v\n", r, r)
		}
	}()

	panic("test")
}
