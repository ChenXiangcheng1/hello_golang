package main

/*
import "fmt"
import "time"
*/
import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"time"

	mylib1 "github.com/chenxiangcheng1/hello-golang/base/lib1" // 别名
	. "github.com/chenxiangcheng1/hello-golang/base/lib2"      // 直接导入当前包
	_ "github.com/chenxiangcheng1/hello-golang/base/lib3"      // 匿名，只执行init函数. _用于避免autofix删除该行
)

// =============================================================================
// 4main
// %v %T
// %p: 指针类型变量
// %t: bool
func hello_print() {
	var a int = 100
	fmt.Printf("%d %T\n", a, a)
	str := "gogogo"
	fmt.Printf("%s %T\n", str, str)
}

// =============================================================================
// 5声明变量
var gA int = 100 // 全局变量(函数体外)必须使用var声明, 不能使用:=

// 声明多个变量
var gB, gC int = 100, 200

var gD, gE = 100, "hello"

var (
	gF, gG int  = 100, 200
	gH, gI      = 100, "hello"
	gJ     bool = true
)

// 茴香豆的四种吃法
func hello_var() {
	var a int
	fmt.Println("a=", a)

	var b int = 100
	fmt.Println("b=", b)

	// var c = 100
	c := 100
	fmt.Println("c=", c)

	d := 100 // 推荐
	fmt.Println("d=", d)
}

// =============================================================================
// 6iota
// iota常量计数器
const (
	SHANGHAI = iota      // 0
	WENZHOU              // 1  // 基础iota
	BEIJING  = iota      // 2
	NINGBO   = 1         // 1
	SHAOXIN              // 1  // 基础常量
	HENAN    = iota      // 5
	AA       = iota      // 6
	BB       = iota * 10 // 70
)

func hello_const() {
	const a int = 100
	fmt.Println("a=", a)
	fmt.Println("SHANGHAI=", SHANGHAI)
	fmt.Println("WENZHOU=", WENZHOU)
	fmt.Println("BEIJING=", BEIJING)
	fmt.Println("NINGBO=", NINGBO)
	fmt.Println("SHAOXIN=", SHAOXIN)
	fmt.Println("HENAN=", HENAN)
	fmt.Println("AA=", AA)
	fmt.Println("BB=", BB)
}

// =============================================================================
// 7函数
// 大写表示对外开放
func Foo(a, b int) (int, int) {
	return a + b, a - b
}

// 按位置返回
func foo1(a int, b int) (int, int) {
	return a + b, a - b
}

// 按变量名返回
func foo2(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

func foo3(a int, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func hello_func() {
	sum, sub := foo1(100, 200)
	fmt.Println("sum=", sum, "sub=", sub)

	sum, sub = foo2(100, 200)
	fmt.Println("sum=", sum, "sub=", sub)

	sum, sub = foo3(100, 200)
	fmt.Println("sum=", sum, "sub=", sub)
}

// =============================================================================
// 8模块
func hello_lib() {
	mylib1.Lib1Test()
	Lib2Test()
}

// =============================================================================
// 10指针
func swap1(a int, b int) {
	var temp int = a
	a = b
	b = temp
}

func swap(pa *int, pb *int) { // 指针类型变量*int
	var temp int = *pa // 解引用*pa
	*pa = *pb
	*pb = temp
}

func hello_point() {
	var a int = 100
	var b int = 200
	swap(&a, &b)
	// swap1(a, b)
	fmt.Println("a =", a, "b =", b)
}

// =============================================================================
// 11defer
// defer压栈，外层函数返回后，defer语句弹栈执行
// 意图：资源清理(替代finally、析构函数)
// 但可能被滥用
func deferFunc() {
	fmt.Println("deferFunc")
}

func returnFunc() int {
	fmt.Println("returnFunc")
	return 0
}

func hello_defer() int {
	defer fmt.Println("111")
	defer fmt.Println("222")
	defer deferFunc()
	return returnFunc()
}

func hello_defer_closure() {
	x := 10
	y := 1
	defer func(y int) {
		fmt.Println(y) // 参数预计算，该参数在defer被声明时立即计算
		fmt.Println(x) // 20  // 闭包引用的变量, 用于延迟计算
	}(y)
	y = 2
	x = 20
}

// =============================================================================
// 12数组和函数形参传递方式(都是值拷贝)
// 区别：
// 其实本质是：slice是含底层数组指针的结构体 拷贝后指向相同的底层数组；而数组是值类型 拷贝是进行完整的拷贝。
func changeArray(myArray [5]int) { // 值拷贝
	myArray[0] = 999 // 修改副本
}

func changeSlice(mySlice []int) { // 值拷贝含指针
	mySlice[0] = 999               // 修改共享底层数组
	mySlice = append(mySlice, 200) // 修改副本
}

func traverseArray(myArray [5]int) {
	for index, value := range myArray { // range必须与for关键字结合使用
		fmt.Println(index, value)
	}
}

func traverseSlice(mySlice []int) {
	// 遍历
	// for i := 0; i < len(mySlice); i++ {
	// 	fmt.Println(mySlice[i])
	// }
	for index, value := range mySlice {
		fmt.Println(index, value)
	}
	// for _, value := range mySlice {
	// 	fmt.Println(value)
	// }
}

func hello_array() {
	var myArray1 [5]int
	myArray2 := [5]int{1, 2, 3}
	mySlice := []int{1, 2, 3} // slice切片
	fmt.Printf("%T %T %T %T\n", myArray1, myArray2, mySlice, myArray2[:])
	changeArray(myArray2)
	traverseArray(myArray2)

	changeSlice(myArray2[:])
	changeSlice(mySlice)
	traverseSlice(myArray2[:])
	traverseSlice(mySlice)
}

// =============================================================================
// 13slice切片
func printSlice(mySlice []int) {
	fmt.Printf("%v %T %d %d\n", mySlice, mySlice, len(mySlice), cap(mySlice))
}

func hello_slice() {
	// 声明与定义(分配空间)
	var mySlice []int              // 未分配空间 nil
	mySlice2 := []int{1, 2, 3}     // 分配空间
	mySlice3 := make([]int, 5, 10) // 分配空间

	printSlice(mySlice)
	printSlice(mySlice2)
	printSlice(mySlice3)
	mySlice = make([]int, 5, 10)
	printSlice(mySlice)

	// 判断
	if mySlice == nil { // nil判断是否初始化 len判断是否为空
		fmt.Println("mySlice is nil")
	} else {
		fmt.Println("mySlice is not nil")
	}

	// 追加，容量不足会扩容为两倍
	printSlice(mySlice2)
	// n := 0  // 先扩容再append
	// for _, err := range iter {
	// 	if err != nil {
	// 		n++
	// 	}
	// }
	// mySlice2 = make([]int, 0, n)
	mySlice2 = append(mySlice2, 4)

	printSlice(mySlice2)

	// 截取(子列表), golang切片共享底层数组, python(list)切片会拷贝, python(numpy)切片共享内存
	// 其实就是有性能需求的就通过共享内存实现
	subSlice1 := mySlice2[0:2]
	subSlice2 := mySlice2[:2]
	printSlice(subSlice1)
	printSlice(subSlice2)
	subSlice1[0] = 999
	mySlice2[1] = 888
	printSlice(subSlice1)
	printSlice(mySlice2)

	// 深拷贝
	copySlice := make([]int, 5, 10)
	copy(copySlice, mySlice2)
	printSlice(copySlice)
}

// =============================================================================
// 15map
func printMap(myMap map[string]string) {
	fmt.Printf("%v %T %p %d\n", myMap, myMap, myMap, len(myMap))
}

func traverseMap(myMap map[string]string) {
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}

func changeMap(myMap map[string]string, key string, value string) {
	myMap[key] = value
}

func hello_map() {
	// 声明与定义
	var myMap1 map[string]string // 未分配空间 nil
	myMap2 := map[string]string{ // 分配空间
		"a": "aaa",
		"b": "bbb",
		"c": "ccc",
	}
	myMap3 := make(map[string]string, 2) // 分配空间

	printMap(myMap1)
	printMap(myMap2)
	printMap(myMap3)
	myMap1 = make(map[string]string, 2)
	printMap(myMap1)

	// 判断
	if myMap3 == nil {
		fmt.Println("myMap3 is nil")
	} else {
		fmt.Println("myMap3 is not nil")
	}

	// 遍历
	traverseMap(myMap2)

	// 追加
	myMap2["d"] = "ddd"
	myMap2["e"] = "eee"
	myMap2["d"] = "999" // 修改
	printMap(myMap2)

	// 删除
	delete(myMap2, "a")
	printMap(myMap2)

	// 函数形参传递方式
	changeMap(myMap2, "b", "888") // 底层包含指针的结构体值传递
	printMap(myMap2)
}

// =============================================================================
// 17struct结构体
type myint int // C++ typedef

func hello_type() {
	var a myint = 1
	fmt.Println(a)
}

type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	book.title = "b"
	book.auth = "bbb"
}

func hello_struct() {
	// 值类型和引用类型除了在形参传递方式以外还有一个区别：
	// 值类型(固定大小)(基本类型、数组、struct(类))声明时未分配空间
	// 引用类型(动态大小)(Slice, Map, Chan, Interface)声明时立即分配空间

	// 声明与定义
	var book1 Book // 分配空间
	book1.title = "a"
	book1.auth = "aaa"
	// book2 := Book{title: "b", auth: "bbb"} // 分配空间
	// book3 := Book{"c", "ccc"}

	fmt.Printf("%v %T %p\n", book1, book1, &book1)
	changeBook(book1) // 值拷贝
	fmt.Printf("%v %T %p\n", book1, book1, &book1)

	// reflectArg(struct {  // 匿名结构体
	// 	name string
	// 	age  int
	// }{"Nemesis", 25})
}

// =============================================================================
// 18OOP#Class = struct + 方法接收器(method receiver)
// 接收器：在函数名之前指定接收器变量名和类型，通过接收器struct实例调用
// 类名大写表示public
type Hero struct {
	Name  string // 属性名大写表示public
	Ad    int
	Level int
}

func (this *Hero) Show() {
	fmt.Println("Name=", this.Name)
	fmt.Println("Ad=", this.Ad)
	fmt.Println("Level=", this.Level)
}

// 指针接收器
// 指针接收器遇到指针直接传递指针hero_point.GetName()，而值接收器遇到指针会自动解引用(*hero_pont).Setname()
func (this *Hero) GetName() string {
	return this.Name
}

// 值接收器
func (this Hero) SetName(name string) {
	this.Name = name // 和Java、Python不同：golang中类(struct)是值传递
}

func (this *Hero) SetAd(ad int) {
	// 和C++指针不同：C++ (*pointer).attr, Goalng pointer.attr
	// this.Ad = ad  // Golang语法糖
	(*this).Ad = ad
}

func hello_class() {
	hero := Hero{Name: "Nemesis", Ad: 100, Level: 1}
	hero.Show()
	hero.SetName("Zeus")
	hero.SetAd(200)
	hero.Show()
}

// =============================================================================
// 19OOP#inherit
type Human struct {
	name   string
	gender string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human      // 匿名嵌入是继承  // 组合优于继承
	StructType struct {
		str string
	}
	// 属性名不能为关键字(struct)、父类名称(Huamn)
	human Human // 命名字段是组合
	Level int
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Show() {
	fmt.Printf("%v\n", this)
	fmt.Println(this.Human)
	fmt.Println(this.name)
	fmt.Println(this.gender)
	fmt.Println(this.StructType)
	fmt.Println(this.human)
	fmt.Println(this.Level)
}

func hello_inherit() {
	human := Human{"Nemesis", "female"}
	human.Eat()
	human.Walk()

	// var superMan SuperMan
	// superMan.name = "Nemesis"
	// superMan.gender = "man"
	// superMan.StructType.str = "struct"
	// superMan.human = Human{name: "human", gender: "man"}
	// superMan.Level = 1
	superMan := SuperMan{Human: Human{name: "Nemesis", gender: "man"}, StructType: struct{ str string }{str: "struct"}, human: Human{name: "human", gender: "man"}, Level: 1}
	// superMan := SuperMan{Human{"Zues", "man"}, struct{ str string }{"str"}, Human{"human", "man"}, 1}
	// superMan := SuperMan{}
	superMan.Eat()
	superMan.Walk()
	superMan.Fly()
	superMan.Show()
}

// =============================================================================
// 20OOP#多态(抽象接口有多种形式 父类指向子类) interface
// Goalng interface本质是：一个包含数据指针和类型信息的struct，数据指针指向接口具体实现，类型信息含方法表
// Golang接口不能包含field字段、不支持默认实现
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

type Dog struct {
	color string
}

// golang实现接口，不需要显式声明
// 使用值接收器，Cat、*Cat类型都实现了该接口
func (this Cat) Sleep() {
	fmt.Println("Cat.Sleep()...")
}

func (this Cat) GetColor() string {
	return this.color
}

func (this Cat) GetType() string {
	return "Cat"
}

// 使用指针接收器，*Dog类型实现了该接口，Dog没有
func (this *Dog) Sleep() {
	fmt.Println("Dog.Sleep()...")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

// func changeAnimal(animal AnimalIF) {
// 	animal.
// }

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Printf("%v %T %p\n", animal, animal, animal)
}

func hello_interface() {
	cat := Cat{"white"}
	dog := Dog{"yellow"}
	// 值接收器
	// showAnimal(cat)
	// showAnimal(&cat)
	// 指针接收器
	// showAnimal(&dog)  // 推荐
	// showAnimal(dog)  // 不行

	// 值接收器
	var animal1 AnimalIF = cat  // 编译器可以调用Cat的值接收器
	var animal2 AnimalIF = &cat // 当需要调用值接收器时，编译器会生成指针接收器包装值接收器
	// 指针接收器
	var animal3 AnimalIF = &dog // 指针接收器遇到指针直接传递指针
	// var animal4 AnimalIF = dog  // 不行，因为dog是值类型，值类型可能是临时值无法获得地址，编译器无法调用指针接收器
	showAnimal(animal1)
	showAnimal(animal2)
	showAnimal(animal3)
}

// =============================================================================
type PetIF interface {
	AnimalIF
	GetAge() int
}

type CommonImpl struct{}

func (this CommonImpl) Run() {
	fmt.Println("CommonImpl.Run()...")
}

type PetPig struct {
	CommonImpl
}

func (this PetPig) Sleep() {
	fmt.Println("PetPig.Sleep()...")
}

func (this PetPig) GetColor() string {
	return "pink"
}

func (this PetPig) GetType() string {
	return "PetPig"
}

func (this PetPig) GetAge() int {
	return 1
}

// =============================================================================

// =============================================================================
// 21interface{}空接口 和类型断言机制(运行时类型检查 与泛型编译时类型检查不同)(类似Java Object、C++ std::any)
func processAnything(arg interface{}) {
	fmt.Printf("%v %T\n", arg, arg)

	// 类型断言
	// 类型断言的本质是：获取arg的pair type，检查是否实现了interface指定的方法
	var v string = arg.(string) // 普通赋值，类型断言失败会panic()
	v, ok := arg.(string)       // 多值返回，类型断言失败不会panic()
	// panic运行时错误, 后续不执行，直接打印 panic: interface conversion: interface {} is int, not string
	fmt.Printf("%v %T\n", v, v)
	fmt.Printf("%v %T\n", ok, ok) // bool

	if ok {
		fmt.Println("is string", v)
	} else {
		fmt.Println("not string")
	}

	// Golang if语法
	// if <condition> {}
	// if <initialization_statement>; <condition> {}

	if v, ok := arg.(int); ok {
		fmt.Println("if: is int", v)
	} else {
		fmt.Println("if: not int", v)
	}

	// Golang switch类型选择(类型断言的语法糖)
	switch v := arg.(type) {
	default:
		fmt.Printf("%v %T\n", v, v)
	}
	switch v := arg.(type) {
	case int:
		fmt.Println("switch: is int", v)
	case string:
		fmt.Println("switch: is string", v)
	case interface{ fooRetStr() fmt.Stringer }:
		fmt.Println("switch: is interface{ fooRetStr() fmt.Stringer }", v)
	default:
		fmt.Println("switch: unknown", v)
	}
	fmt.Println()
}

func hello_typeassert() {
	processAnything(42)      // int
	processAnything("hello") // string
	processAnything(struct{}{})
}

// =============================================================================
// 22pair 变量结构pair<type=static|concrete, value>
func hello_pair() {
	// tty: pair<type:*os.File文件描述符, value:"/dev/tty">
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0) // /dev/tty是当前终端设备文件，驱动会读取该文件的内容打印到当前终端
	fmt.Printf("%v %T\n", tty, tty)
	fmt.Printf("%v %T\n", err, err)
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}

	var r io.Reader = tty // io.Reader是一个接口含Read()方法 // 通过`go doc os.File`发现实现io.Reader.Read()的接收器类型为*os.File，所以要传递*os.File给io.Reader
	fmt.Printf("%v %T\n", r, r)
	var w io.Writer = r.(io.Writer) // 类型断言
	fmt.Printf("%v %T\n", w, w)
	w.Write([]byte("hello\n"))
}

// =============================================================================
// 23reflect反射机制
func reflectArg(arg interface{}) {
	fmt.Printf("%v %T\n", arg, arg)
	inputType := reflect.TypeOf(arg)   // *reflect.rtype
	inputValue := reflect.ValueOf(arg) // reflect.Value
	fmt.Printf("%v %T\n", inputType, inputType)
	fmt.Printf("%v %T\n", inputValue, inputValue)

	kind := inputType.Kind()
	fmt.Printf("%v %T\n", kind, kind) // reflect.Kind
	if inputType.Kind() == reflect.Struct {
		// 字段
		numField := inputType.NumField()
		fmt.Printf("%v %T\n", numField, numField)
		for i := 0; i < inputType.NumField(); i++ {
			field := inputType.Field(i)
			field_name := field.Name
			field_type := field.Type
			fmt.Printf("%v %T\n", field, field) // reflect.StructField{Name; PkgPath; Type; Tag; Offset; Index例如[0]; Anonymous是否是匿名字段}
			fmt.Printf("%v %T\n", field_name, field_name)
			fmt.Printf("%v %T\n", field_type, field_type)
			value := inputValue.Field(i)
			canInterface := value.CanInterface()
			fmt.Printf("%v %T\n", canInterface, canInterface)
			if canInterface {
				valueIF := inputValue.Field(i).Interface()
				fmt.Printf("%v %T\n", value, value)
				fmt.Printf("%v %T\n", valueIF, valueIF)
			}
		}
		// 方法
		numMethod := inputType.NumMethod()
		fmt.Printf("%v %T\n", numMethod, numMethod)
		for i := 0; i < inputType.NumMethod(); i++ {
			method := inputType.Method(i)
			method_name := method.Name
			method_type := method.Type
			fmt.Printf("%v %T\n", method, method)
			fmt.Printf("%v %T\n", method_name, method_name)
			fmt.Printf("%v %T\n", method_type, method_type)
		}
	}
}

func hello_reflect() {
	var num float64 = 3.14
	reflectArg(num)
	reflectArg(struct {
		name string
		age  int
	}{"Nemesis", 25})
}

// =============================================================================
// 24结构体标签with reflect
// 结构体标签常用于：json编解码、ORM映射
type resume struct { // 摘要
	Name   string `info:"name" doc:"测试"`
	Gender string `info:"gender"`
}

func findTag(arg interface{}) {
	inputType := reflect.TypeOf(arg)
	fmt.Printf("%v %T\n", inputType, inputType) // *reflect.rtype
	t := reflect.TypeOf(arg).Elem()             // Elem()获取指针指向的类型，即从指针类型获取值类型. 具体实现类型还是*reflect.rtype
	fmt.Printf("%v %T\n", t, t)                 // *reflect.rtype
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag
		fmt.Printf("%v %T\n", tag, tag) // reflect.StructTag
		taginfo := field.Tag.Get("info")
		tagdoc := field.Tag.Get("doc")
		fmt.Printf("%v %T\n", taginfo, taginfo)
		fmt.Printf("%v %T\n", tagdoc, tagdoc)
	}
}

func hello_structTag() {
	var re resume
	findTag(&re)
}

// =============================================================================
// 25结构体标签with json
type Movie struct {
	Title  string   `json:"title"`
	Year   int      `jason:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"test"`
}

func hello_json() {
	movie := Movie{"my life", 2000, 100, []string{"Nemesis"}}
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal failed", err)
		return
	}
	fmt.Printf("%s %v %T\n", jsonStr, jsonStr, jsonStr) // []uint8

	my_movie := Movie{}
	err2 := json.Unmarshal(jsonStr, &my_movie)
	if err2 != nil {
		fmt.Println("json unmarshal failed", err2)
		return
	}
	fmt.Printf("%v %T\n", my_movie, my_movie)
}

// =============================================================================
// 26co-routine、goroutine(Golang协程)、GMP

// co-routine(协程):
// 协作式(cooperative 非抢占式协程)
// Golang: 不支持非抢占式协程
// Python: async-await 本质是由事件循环(主线程)调度的、通过await传递消息的一个可以挂起可恢复的函数
// Java: 不支持非抢占式协程，提供CompletableFuture进行异步编程

// 抢占式协程(协程调度器)
// Golang: goroutine
// Python: 不支持抢占式协程
// Java: Virtual threads

// GMP(goroutine user_thread processor)
// GMP架构看笔记图

// goroutine 协程 // 由go运行时调度的用户态的并发执行体(并发执行的函数)  // A goroutine is a lightweight thread managed by the Go runtime.

// =============================================================================
// 27goroutine
func newTask() {
	i := 0
	for { // 死循环
		i++
		fmt.Printf("newTask: %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func hello_goroutine() {
	go func() { // go: 创建一个goroutine并发协程并执行
		defer fmt.Println("A.defer")

		func() { // 匿名函数
			defer fmt.Println("B.defer")
			// return           // return当前函数
			runtime.Goexit() // 退出当前goroutine，退出之前会执行defer
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	go func(a int, b int) bool {
		fmt.Println(a, b)
		return true
	}(10, 20)

	go newTask()

	i := 0
	for {
		i++
		fmt.Printf("main: %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// =============================================================================
// 28channel, 用于goroutine间通信和同步

// 无缓冲的同步机制: 发送者和接收者必须同时准备好(上锁)、先到达的操作`channel <- <value>`或`value, ok := <-channel`会阻塞
// 有缓冲的同步机制: 发送者和接收者不需要同时准备好(不上锁)、缓冲区满或空时会阻塞

func hello_channel_without_buffer() {
	defer time.Sleep(1 * time.Second)
	defer fmt.Println("main程结束")
	fmt.Println("main程正在运行...")

	// 定义
	channel := make(chan int, 0)            // make(chan Type, capacity缓冲区容量)
	fmt.Printf("%v %T\n", channel, channel) // 类型: chan int
	fmt.Println(len(channel), cap(channel))

	// go程
	go func() {
		defer fmt.Println("go程结束")
		fmt.Println("go程正在运行...")
		channel <- 100
		close(channel) // data, ok := <-channel ok会返回false
		// channel <- 100 // panic: 发送到关闭的channel
	}()

	// main程
	for data := range channel { // range本质就是语法糖
		fmt.Println("main程:", data)
	}
	// for {
	// 	if data, ok := <-channel; ok { // ok表示检查channel是否已关闭或为空
	// 		fmt.Printf("%v %T\n", ok, ok)
	// 		fmt.Println("main程:", data)
	// 	} else {
	// 		break
	// 	}
	// }
}

func hello_channel_with_buffer() {
	defer time.Sleep(1 * time.Second)
	defer fmt.Println("main程结束")
	fmt.Println("main程正在运行...")
	channel := make(chan int, 2)

	go func() {
		defer fmt.Println("go程结束")
		fmt.Println("go程正在运行...")
		for i := 0; i < 3; i++ {
			channel <- i
			fmt.Printf("go程: 发送%d %d %d\n", i, len(channel), cap(channel))
		}
	}()

	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		num := <-channel
		fmt.Printf("main: 接收%d %d %d\n", num, len(channel), cap(channel))
	}
}

// =============================================================================
// 32select
func fibonacci(channel, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case channel <- x:
			temp := x
			x = y
			y = temp + y
		case <-quit:
			fmt.Println("main程结束")
			return
		default:
			// fmt.Println("main: default")
		}
	}
}

func hello_select() {
	channel := make(chan int, 0)
	quit := make(chan int, 0)

	// go程
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-channel)
		}
		quit <- 0
	}()

	// main程
	fibonacci(channel, quit)
}

// =============================================================================
func hello_num() {
	num2 := 0b1010
	num8 := 0o666
	num16 := 0xFF
	fmt.Printf("%d %d %d\n", num2, num8, num16)
}

// =============================================================================
func main() {
	// hello_print()
	// hello_var()
	// hello_const()
	// hello_func()
	// hello_lib()
	// hello_point()
	// hello_defer()
	// hello_defer_closure()
	// hello_array()
	// hello_slice()
	// hello_map()
	// hello_type()
	// hello_struct()
	// hello_class()
	// hello_inherit()
	// hello_interface()
	// hello_typeassert()
	// hello_pair()
	// hello_reflect()
	// hello_structTag()
	// hello_json()
	// hello_goroutine()
	// hello_channel_without_buffer()
	// hello_channel_with_buffer()
	// hello_select()
	hello_num()
}
