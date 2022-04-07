package array

import (
	"fmt"
	"math/rand"
	"time"
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

func help(arr *[3]int) {
	(*arr)[0] = 88
	fmt.Println("In help  ", *arr)
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

	for index, value := range strArr05 {
		fmt.Println(index, value)
	}

	for _, value := range strArr05 {
		fmt.Println(value)
	}

	var arr [3]int = [3]int{1, 2, 3}
	help(&arr)
	fmt.Println(arr)
}

func Test04() {
	var alphas [26]byte
	for i := 0; i < 26; i++ {
		alphas[i] = byte(int('A') + i)
	}

	for _, v := range alphas {
		fmt.Printf("%c", v)
	}
	fmt.Printf("\n")
}

func Test05() {
	rand.Seed(time.Now().Unix())
	var nums [5]int
	for i := 0; i < 5; i++ {
		nums[i] = rand.Intn(10) + 1
	}

	for _, v := range nums {
		fmt.Println(v)
	}

	for i := 4; i >= 0; i-- {
		fmt.Println(nums[i])
	}
}
