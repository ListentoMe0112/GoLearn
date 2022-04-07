package func_test

import "fmt"

var Age int
var Name string

func init() {
	Age = 18
	Name = "Xieyibo"
	fmt.Println("func_test INIT()")
}

func Cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("operator wrong")
	}

	return res
}

func Test(n1 int) {
	n1 = n1 + 1
	fmt.Println("[In func] n1 = ", n1, " The location is ", &n1)
}

func Test01(pi *int) {
	*pi = 2020
}

type myFunc func(float64, float64, byte) float64

func CalTest(funcvar myFunc, num1 float64) float64 {
	var num float64 = funcvar(10.0, 20.0, '+')
	return num1 + num
}

func Test02(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func Test03(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func Test04() {
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}

	fmt.Println(res1(10, 20))
	fmt.Println(res1(3, 4))
}
