package test02

import "fmt"

func Test() {
	question1()
	question2()
}

func question1() {
	fmt.Printf("%d weeks and %d days\n", 97/7, 97%7)
}

func question2() {
	fmt.Printf("%f\n", (97.7-32)/(9.0/5))
}
