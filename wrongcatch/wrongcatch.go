package wrongcatch

import (
	"errors"
	"fmt"
)

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err = ", err)
		}
	}()
	var num1 int = 10
	var num2 int = 0
	fmt.Println(num1 / num2)
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("reading file fail")
	}
}

func Test01() {
	test()
	fmt.Println("Continue")
}

func Test02() {
	err := readConf("config.init")
	if err != nil {
		panic(err)
	}
}
