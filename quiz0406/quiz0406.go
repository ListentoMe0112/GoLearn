package quiz0406

import (
	"fmt"
	"math/rand"
	"time"
)

func Test01() {
	var i int
	for {
		fmt.Scanln(&i)
		if i <= 0 || i >= 2022 {
			fmt.Println("Wrong Year")
			continue
		}
		fmt.Scanln(&i)
		if i <= 0 || i >= 12 {
			fmt.Println("Wrong Month")
			continue
		}
		fmt.Scanln(&i)
		if i <= 0 || i >= 31 {
			fmt.Println("Wrong Day")
			continue
		}
		break
	}
}

func Test02() {
	rand.Seed(time.Now().Unix())
	var i int = rand.Intn(10) + 1
	fmt.Println(i)
	var x int
	var times int = 0
	for {
		times++
		fmt.Scanln(&x)
		if x == i {
			if times == 1 {
				fmt.Println("ok1")
			} else if times == 2 {
				fmt.Println("ok2")
			} else {
				fmt.Println("ok3")
			}
			break
		}
	}
}
