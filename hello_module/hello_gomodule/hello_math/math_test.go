package hello_math_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestConst(t *testing.T) {
	num := math.MaxInt
	fmt.Printf("%T %v\n", num, num)
	num = math.MaxInt64
	fmt.Printf("%T %v\n", num, num)
	f := math.Pi
	fmt.Printf("%T %v\n", f, f)
}

// Not a Number
func TestIsNaN(t *testing.T) {
	// false
	var ff float64
	fmt.Printf("%T %v\n", ff, ff)
	fmt.Printf("%v\n", math.IsNaN(ff))
	fmt.Printf("%v\n", math.IsNaN(math.Pi))

	// true  // 无效的浮点数操作结果
	f := math.NaN()
	fmt.Printf("%T %v\n", f, f)
	fmt.Printf("%v\n", math.IsNaN(f))
	f = math.Inf(1) - math.Inf(1)
	fmt.Printf("%T %v\n", f, f)
	fmt.Printf("%v\n", math.IsNaN(f))
}

// 正负无穷值
func TestInf(t *testing.T) {
	f := math.Inf(1)
	fmt.Printf("%T %v\n", f, f)
	f = math.Inf(-1)
	fmt.Printf("%T %v\n", f, f)
}

// 取整
func TestCeilFloor(t *testing.T) {
	f := 3.14159
	// 天花板取整、地板取整
	ret := math.Ceil(f)
	fmt.Printf("%T %v\n", ret, ret)
	fmt.Printf("%v\n", math.Ceil(4)) // >=

	fmt.Printf("%v\n", math.Floor(f))
	fmt.Printf("%v\n", math.Floor(3)) // <=

	// 截取取整、四舍五入取整
	fmt.Printf("%v\n", math.Trunc(f))
	fmt.Printf("%v\n", math.Round(f))
}

// 返回最值
func TestMaxMin(t *testing.T) {
	ret := math.Max(20, 10)
	fmt.Printf("%T %v\n", ret, ret)

	ret = math.Min(20, 10)
	fmt.Printf("%T %v\n", ret, ret)

	ret = math.Dim(20, 10)
	fmt.Printf("%T %v\n", ret, ret)
	ret = math.Dim(10, 20)
	fmt.Printf("%T %v\n", ret, ret)
}

// 其他
// math.Abs()
func TestMain(t *testing.T) {
	fmt.Printf("%v\n", math.Abs(-3.14))
	fmt.Printf("%v\n", math.Mod(11.1, 3)) // 取模
}

func TestRoot(t *testing.T) {
	fmt.Printf("%v\n", math.Sqrt(144))
	fmt.Printf("%v\n", math.Cbrt(27))
	fmt.Printf("%v\n", math.Hypot(3, 4)) // 求平方和的平方根
}

func TestPow(t *testing.T) {
	fmt.Printf("%v\n", math.Pow(2, 3))
	fmt.Printf("%v\n", math.Pow10(3))
}

// 三角函数
func TestTrigonometric(t *testing.T) {
	fmt.Printf("%v\n", math.Sin(math.Pi/2)) // 1
	fmt.Printf("%v\n", math.Cos(math.Pi))   // -1
	fmt.Printf("%v\n", math.Tan(math.Pi/4)) // 1
}

func TestLog(t *testing.T) {
	fmt.Printf("%v\n", math.Log(math.E))
	fmt.Printf("%v\n", math.Log10(100))
	fmt.Printf("%v\n", math.Log2(8))
}

func TestSignbit(t *testing.T) {
	fmt.Printf("%v\n", math.Signbit(-1)) // true
	fmt.Printf("%v\n", math.Signbit(-0)) // false
	fmt.Printf("%v\n", math.Signbit(0))  // false
	fmt.Printf("%v\n", math.Signbit(+0)) // false
	fmt.Printf("%v\n", math.Signbit(1))  // false
}

func TestRand(t *testing.T) {
	fmt.Printf("%v\n", rand.Int())    // [0, 2^63)
	fmt.Printf("%v\n", rand.Int31())  // [0, 2^31)
	fmt.Printf("%v\n", rand.Intn(5))  // [0, n)
	fmt.Printf("%v\n", rand.Uint32()) // [0, 2^32)

	slice := rand.Perm(5)
	fmt.Printf("%T %v\n", slice, slice) // [n]int [0,n)

	rs := rand.NewSource(time.Now().UnixNano())
	fmt.Printf("%T\n", rs)
	localRand := rand.New(rs)
	fmt.Printf("%T %v\n", localRand, localRand)
	fmt.Printf("%v %v\n", localRand.Int(), localRand.Int())
}
