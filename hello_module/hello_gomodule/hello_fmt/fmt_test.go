package hello_fmt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

/*
General:
%v 打印value的默认格式化
%+v 添加字段名称
%#v 浮点无穷, NaNs被打印为±Inf, NaN
%T 打印value的类型
%% 转义打印%
......
*/
func TestFormat(t *testing.T) {
	fmt.Printf("%q\n", "1\t2\n") // 用于查看原始字符串，显示字面量而不是真换行
}

// 填充
func TestPad(t *testing.T) {
	f := 3.1415
	fmt.Printf("|%5.2f|\n", f)                // 左填充
	fmt.Printf("|%*.*f|\n", 5, 2, f)          // 左填充+*参数
	fmt.Printf("|%[2]*.[1]*[3]f|\n", 2, 5, f) // 左填充+*参数+索引
	fmt.Printf("|%-5.2f|\n", f)               // 右填充
	fmt.Printf("|%05.2f|\n", f)               // 0填充
}

func TestFprintf(t *testing.T) {
	f := 3.14

	ret, err := fmt.Fprintf(os.Stdout, "%T %v\n", f, f)
	if err != nil {
		t.Errorf("Fprintf error: %v\n", err)
	} else {
		fmt.Printf("%T %v\n", ret, ret) // 字节数
	}

	fileObj, err := os.OpenFile("./testdata/data.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		t.Errorf("OpenFile error: %v\n", err)
	}
	fmt.Printf("%T %v\n", fileObj, fileObj) // *os.File
	ret2, err := fmt.Fprintln(fileObj, "test")
	if err != nil {
		t.Errorf("Fprintln error: %v\n", err)
	} else {
		fmt.Printf("%T %v\n", ret2, ret2)
	}
}

func TestSprintf(t *testing.T) {
	f := 3.14
	str := fmt.Sprintf("%v\n", f)
	fmt.Printf("%T %vover\n", str, str)
}

// fmt.Errorf(str) -> *errors.errorString
func TestErrorf(t *testing.T) {
	err := fmt.Errorf("%v\n", "error")
	fmt.Printf("%T %v\n", err, err) // *errors.errorString
}

// fmt.Printf() 输出到os.Stdout标准输出
// fmt.Fprintf() 输出到io.Writer
// fmt.Sprintf() 输出到string
// fmt.Errorf() 返回*errors.errorString

// 运行时从os.Stdin标准输入读取变量值
// fmt.Scan() 以空格作为分隔符
// fmt.Scanln() 以空格作为分隔符, 以换行符作为结束符
// fmt.Scanf(fmtStr, ...any) 以空格作为分隔符，按格式化字符串读取输入
// fmt.Fscanf(io.Reader, string, ...any) 从io.Reader读取变量值
// fmt.Sscanf(string, ...any) 从string读取变量值
func TestScanf(t *testing.T) {
	var f float64
	var str string
	input := strings.NewReader("13.14q2hello world")
	// go test时读取os.Stdin会返回EOF
	// ret, err := fmt.Scanf("1%f2%s", &f, &str)
	ret, err := fmt.Fscanf(input, "1%fq2%s", &f, &str)
	if err != nil {
		t.Errorf("Scanf error: %v\n", err)
	}
	fmt.Printf("%T %v\n", ret, ret) // 返回成功读取的参数个数
	fmt.Printf("%T %v\n", f, f)
	fmt.Printf("%T %v\n", str, str)
}

// bufio.NewReader(io.Reader) -> *bufio.Reader 能读取空格
func TestBufio(t *testing.T) {
	input := strings.NewReader("13.14q2hello world\n") // *strings.Reader
	// reader := bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(input) // *bufio.Reader
	fmt.Printf("Please input a float and a string: \n")
	ret, err := reader.ReadString('\n')
	if err != nil {
		t.Errorf("Scanf error: %v\n", err)
	}
	fmt.Printf("%T %v\n", ret, ret)
}
