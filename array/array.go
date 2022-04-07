package array

import (
	"fmt"
)

func Test01() {
	var hens [6]float64
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	var total float64
	var i int
	for i = 0; i < len(hens); i++ {
		total += hens[i]
		fmt.Println(hens[i])
	}
	var avg float64 = total / float64(i)
	fmt.Println(total, avg)

}

func Test02() {
	var scores [6]float64
	for i := 0; i < len(scores); i++ {
		fmt.Scanln(&scores[i])
	}

	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}
}

func Test03() {
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println(numArr01)

	numArr02 := [3]int{1, 2, 3}
	fmt.Println(numArr02)

	var numArr03 [3]int = [3]int{1, 2, 3}
	fmt.Println(numArr03)

	numArr04 := [3]int{1: 800}
	fmt.Println(numArr04)

	strArr05 := [...]string{1: "tom", 0: "jcak", 2: "mary"}
	fmt.Println(strArr05)
}
