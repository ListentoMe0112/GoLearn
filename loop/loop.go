package loop

import (
	"fmt"
	"math/rand"
	"time"
)

func Test01() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Hello world")
	}
}

func Test02() {
	var str string = "hello world， 北京"

	str_r := []rune(str)

	for i := 0; i < len(str_r); i++ {
		fmt.Printf("%c\n", str_r[i])
	}

	for index, val := range str_r {
		fmt.Printf("index = %d, val = %c\n", index, val)
	}
}

func Test03() {
	var cnt int = 0
	var sum int = 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			cnt++
			sum += i
		}
	}
	fmt.Println(cnt, sum)
}

func Test04() {
	var a int = 0
	var b int = 6
	for b >= 0 {
		fmt.Printf("%d + %d = 6\n", a, b)
		a++
		b--
	}
}

func Test05() {
	var i int = 1
	for i < 10 {
		fmt.Println("hello world", i)
		i++
	}
}

func Test06() {
	var n1 float64
	for i := 1; i <= 3; i++ {
		var class_val float64
		var cnt int
		for j := 1; j <= 5; j++ {
			fmt.Scanln(&n1)
			class_val += n1
			if n1 >= 60 {
				cnt++
			}
		}
		fmt.Println(class_val/5, cnt)
	}
}

func Test07() {
	for i := 1; i <= 6; i++ {
		for j := 6 - i; j >= 0; j-- {
			fmt.Printf(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Printf("*")
		}
		// for j := 6 - i; j >= 0; j-- {
		// 	fmt.Printf(" ")
		// }
		fmt.Println("\n")
	}
}

func Test08() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Printf("\n")
	}
}

func Test09() {
	rand.Seed(time.Now().Unix())
	var cnt int = 0
	var n int
	for {
		cnt++
		n = rand.Intn(100) + 1
		if n == 99 {
			break
		}
	}
	fmt.Println(cnt)
}

func Test10() {
	for i := 0; i < 4; i++ {
	lable1:
		for j := 0; j < 10; j++ {
			if j == 2 {
				break lable1
			}
			fmt.Println("j = ", j)
		}
	}
}

func Test11() {
	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println(i)
			break
		}
	}
}

func Test12() {
	for i := 0; i < 3; i++ {
		var username string
		var password string
		fmt.Scanln(&username)
		fmt.Scanln(&password)
		if username == "Xieyibo" && password == "123" {
			fmt.Println("Login Success")
			break
		} else {
			fmt.Printf("%d times remained\n", 3-i-1)
		}
	}
}

func Test13() {
here:
	for i := 0; i < 2; i++ {
		for j := 1; j < 4; j++ {
			if j == 2 {
				continue here
			}

			fmt.Println("i = ", i, ", j = ", j)
		}
	}
}

func Test14() {
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		} else {
			continue
		}
	}
}

func Test15() {
	var n float64
	for {
		fmt.Scanln(&n)
		if n != 0 {
			if n > 0 {
				fmt.Println("positive")
				continue
			} else {
				fmt.Println("negtive")
				continue
			}
		}
		break
	}
}

func Test16() {
	fmt.Println("ok1")
	goto label1
	fmt.Println("ok2")
	fmt.Println("ok3")
label1:
	fmt.Println("ok4")
	fmt.Println("ok5")
}
