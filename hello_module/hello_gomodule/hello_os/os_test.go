package hello_os_test

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
	"testing"
)

/*
os标准库本质就是对POSIX的封装
*/

func TestEnv(t *testing.T) {
	envStrArr := os.Environ()
	fmt.Printf("%T %v\n\n", envStrArr, envStrArr)

	envMap := make(map[string]string)
	for _, envStr := range envStrArr {
		parts := strings.SplitN(envStr, "=", 2)
		envMap[parts[0]] = parts[1]
	}
	for k := range envMap {
		fmt.Printf("%v\n", k)
	}
	fmt.Println()

	env, ok := os.LookupEnv("http_proxy") // http_proxy GOPATH
	if ok {
		fmt.Printf("%T %v\n", env, env)
	} else {
		fmt.Printf("the env variable is not exist\n")
	}
}

func TestAdd(t *testing.T) {
	err := os.MkdirAll("./testdata/dir", 0o666) // 创建目录
	if err != nil {
		t.Fatalf("failed to os.MkdirAll(): %v", err)
	}
	file, err := os.Create("./testdata/dir/data.txt") // 创建文件 // 本质OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	if err != nil {
		t.Errorf("failed to os.Create(): %v", err)
	}
	fmt.Printf("%T %v\n", file, file)
}

func TestDelete(t *testing.T) {
	err := os.RemoveAll("./testdata")
	if err != nil {
		t.Errorf("failed to os.RemoveAll(): %v", err)
	}
}

// filepath.Join()
// filepath.Dir()
// filepath.Base()
func TestSelect(t *testing.T) {
	err := os.Chdir("./testdata") // 更改current working directory
	if err != nil {
		t.Fatalf("failed to os.Chdir(): %v", err)
	}

	dir, err := os.Getwd() // absolute path name of current working directory
	if err != nil {
		t.Errorf("failed to os.Getwd(): %v", err)
	}
	fmt.Printf("%T %v\n", dir, dir)

	fileStat, err := os.Stat("./testdata")
	if err != nil {
		t.Errorf("failed to os.Stat(): %v", err)
		fmt.Printf("%v\n", errors.Is(err, fs.ErrExist))
	}
	fmt.Printf("%T %v\n", fileStat, fileStat)
	fmt.Printf("%v %v\n", fileStat.Name(), fileStat.Size())
}

func TestOpenFile(t *testing.T) {
	// /dev/tty
	// flag: O_RDONLY O_WRONLY O_RDWR | O_CREATE O_APPEND(尾) O_TRUNC(头)
	// perm: 八进制, 文件已存在时忽略perm
	file, err := os.OpenFile("/dev/tty", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o666)
	if err != nil {
		t.Errorf("failed to os.OpenFile(): %v", err)
	}
	fmt.Printf("%T %v %v\n", file, file, file.Name())

	n, err := file.WriteString("tty: Hello world!")
	if err != nil {
		t.Errorf("failed to file.WriteString(): %v", err)
	}
	fmt.Printf("%v\n", n)
	err = file.Close()
	if err != nil {
		t.Errorf("failed to file.Close(): %v", err)
	}

	// testdata
	file, err = os.OpenFile("./testdata/data1.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		t.Errorf("failed to os.OpenFile(): %v", err)
	}
	content := make([]byte, 1024)
	n, err = file.Read(content)
	if err != nil {
		if errors.Is(err, io.EOF) {
			fmt.Printf("%v %T %v\n", n, err, err)
		} else {
			t.Errorf("failed to file.Read(): %v %v", n, err)
		}
	}
	str := string(content[:n])
	fmt.Printf("%T %v %v %v\n", str, str, n, len(content))
}
