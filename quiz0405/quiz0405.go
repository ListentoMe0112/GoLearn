package quiz0405

import "fmt"

// Name := "tom" 不能赋值，只能声明

var Name string = "xieyibo"

func Test01() {
	fmt.Println(Name)
}

func Test02() {
	Name := "jack"
	fmt.Println(Name)
}

func Test03(x int) {
	for i := 1; i <= x; i++ {
		for j := x - i; j >= 0; j-- {
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

func Test04(x int) {
	for i := 1; i <= x; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Printf("\n")
	}
}
