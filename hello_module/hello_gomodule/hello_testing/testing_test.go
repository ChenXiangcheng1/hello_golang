package xxx_test

import (
	"fmt"
	"testing"
)

/*
	doc: https://devdocs.io/go/testing/index
	go help test
	go help testflag
*/

/*
	go help testfunc
	func TestXxx(t *testing.T) { ... }  // 单元测试
	func BenchmarkXxx(b *testing.B) { ... }  // 基准测试
	func FuzzXxx(f *testing.F) { ... }  // 模糊测试
	func ExampleType_Method() { ... }  // 示例函数
*/

func TestXxx(t *testing.T) {
	// fmt.Printf("%T %v\n", t, t) // 对于测试，仅当测试失败或设置了-test.v标志时，才会打印文本
	// t.Logf("%T %v \n", t, t)
	got := custom_abs(-1)
	if got != 1 {
		// 测试函数不应该有返回，使用t.Errorf()代替
		t.Errorf("Abs(-1) = %d; want 1", got) // 断言函数：Error、fatal、Fail
		t.Fatalf("over %v\n", got)            // t.Logf() + t.failNow()
	}
}

func BenchmarkXxx(b *testing.B) {
	// fmt.Printf("%T %v\n", b, b)
	// b.Logf("%T %v \n", b, b)
	b.Logf("%T %T %v\n", b, b.N, b.N)
	for b.Loop() { // 等价于 for i := 0; i < b.N; i++ {}
		custom_abs(-1)
	}
}

func ExamplePrintln() {
	fmt.Println("The output of\nthis example.")
	// Unordered output: 4
	// 2
	// 1
	// Output: The output of
	// this example.
}

// 命令行：
// 本地目录模式(当前目录执行)
// go test  // go test时读取os.Stdin会返回EOF

// 包列表模式(会缓存 -count=1禁用缓存)
// go test -timeout 30s github.com/chenxiangcheng1/hello-golang/hello_module/hello_gomodule/hello_testing
// ok      github.com/chenxiangcheng1/hello-golang/hello_module/hello_gomodule/hello_testing       0.001s

// -run指定测试函数
// go test -run ^TestXxx$ github.com/chenxiangcheng1/hello-golang/hello_module/hello_gomodule/hello_testing
// ok      github.com/chenxiangcheng1/hello-golang/hello_module/hello_gomodule/hello_testing       0.001s

// -v 打印完整的测试输出(测试失败也会)
// go test -v github.com/chenxiangcheng1/hello-golang/hello_module/hello_gomodule/hello_testing
// === RUN   TestXxx
// --- PASS: TestXxx (0.00s)
// PASS
// ok      github.com/chenxiangcheng1/hello-golang/hello_module/hello_gomodule/hello_testing       0.001s

// -bench 基准测试
// go test -bench=^BenchmarkXxx$ github.com/chenxiangcheng1/hello-golang/hello_module/hello_gomodule/hello_testing
// goos: linux
// goarch: amd64
// pkg: github.com/chenxiangcheng1/hello-golang/hello_module/hello_gomodule/hello_testing
// cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
// BenchmarkXxx-12         1000000000               0.2583 ns/op
// --- BENCH: BenchmarkXxx-12
//     hello_testing__test.go:19: *testing.B int 1
//     hello_testing__test.go:19: *testing.B int 100
//     hello_testing__test.go:19: *testing.B int 10000
//     hello_testing__test.go:19: *testing.B int 1000000
//     hello_testing__test.go:19: *testing.B int 100000000
//     hello_testing__test.go:19: *testing.B int 1000000000
// PASS
// ok      github.com/chenxiangcheng1/hello-golang/hello_module/hello_gomodule/hello_testing       0.288s

// -benchmem
// go test -benchmem -bench=^BenchmarkXxx$ github.com/chenxiangcheng1/hello-golang/hello_module/hello_gomodule/hello_testing
// goos: linux
// goarch: amd64
// pkg: github.com/chenxiangcheng1/hello-golang/hello_module/hello_gomodule/hello_testing
// cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
// BenchmarkXxx-12         1000000000               0.2611 ns/op          0 B/op          0 allocs/op
// 基准测试名称				循环次数			每次操作的平均时间		每次操作的内存分配		每次操作的内存分配次数
// --- BENCH: BenchmarkXxx-12
//     hello_testing__test.go:19: *testing.B int 1
//     hello_testing__test.go:19: *testing.B int 100
//     hello_testing__test.go:19: *testing.B int 10000
//     hello_testing__test.go:19: *testing.B int 1000000
//     hello_testing__test.go:19: *testing.B int 100000000
//     hello_testing__test.go:19: *testing.B int 1000000000
// PASS
// ok      github.com/chenxiangcheng1/hello-golang/hello_module/hello_gomodule/hello_testing       0.292s
