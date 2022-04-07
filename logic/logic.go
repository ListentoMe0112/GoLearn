package logic

import (
	"fmt"
	"math"
)

func Test01() {
	var age byte
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("Your age is below 18, please be responsible for yourself")
	} else if age > 12 {
		fmt.Println("Your age is below 12, please be responsible ")
	} else {
		fmt.Println("Your are too small ")
	}
}

func Test02() {
	var num1 int
	var num2 int
	fmt.Scanf("%d %d\n", &num1, &num2)
	if num1+num2 >= 50 {
		fmt.Println("Hello World")
	}
}

func Test03() {
	var num1 float64
	var num2 float64
	fmt.Scanf("%f %f\n", &num1, &num2)
	if num1 > 10.0 && num2 < 20.0 {
		fmt.Println(num1 + num2)
	}
}

func Test04() {
	var num1 int
	var num2 int
	fmt.Scanf("%d %d\n", &num1, &num2)
	if (num1+num2)%3 == 0 && (num1+num2)%5 == 0 {
		fmt.Println("ok")
	}
}

func Test05() {
	var num1 int
	fmt.Scanf("%d\n", &num1)
	if (num1%4 == 0 && num1%100 != 0) || num1%400 == 0 {
		fmt.Println("ok")
	}
}

func Test06() {
	var score int
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("BMW")
	} else if score > 80 && score <= 99 {
		fmt.Println("iphone 7p")
	} else if score >= 60 && score <= 80 {
		fmt.Println("ipad")
	} else {
		fmt.Println("Nothing")
	}
}

func Test07() {
	var b bool = true
	if b == false {
		fmt.Println("a")
	} else if !b {
		fmt.Println("b")
	} else if b {
		fmt.Println("C")
	}
}

func Test08() {
	var a float64 = 3.0
	var b float64 = 1
	var c float64 = 6.0
	m := b*b - 4*a*c
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / (2 * a)
		x2 := (-b - math.Sqrt(m)) / (2 * a)
		fmt.Println("Two: ", x1, x2)
	} else if m == 0 {
		x1 := -b / (2 * a)
		fmt.Println("One: ", x1)
	} else {
		fmt.Println("No solution")
	}

}

func Test09() {
	var height float64
	var fortune float64
	var handsome bool
	fmt.Scanf("%f %f %t\n", &height, &fortune, &handsome)
	if height > 180.0 && fortune > 10000000 && handsome {
		fmt.Println("ok1")
	} else if height > 180.0 || fortune > 10000000 || handsome {
		fmt.Println("ok2")
	} else {
		fmt.Println("ok3")
	}
}
