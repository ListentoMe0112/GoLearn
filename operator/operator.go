package op_test

import "fmt"

func Test01() {
	var n1 int = 9
	var n2 int = 8
	fmt.Printf("n1 == n2 : %t\n", n1 == n2)
}

func Test02() {
	var age int = 40
	if age > 10 && age < 50 {
		fmt.Println("ok1")
	}

	if age > 30 && age <= 40 {
		fmt.Println("ok2")
	}
}

func Test03() {
	a := 9
	b := 2
	fmt.Println(a, b)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}

func Test04() {
	var a int = 10
	var b int = 11
	var c int = 12
	var max int
	if a < b {
		max = b
	} else {
		max = a
	}

	if max < c {
		max = c
	}

	fmt.Printf("%d\n", max)
}
