package dateTime

import (
	"fmt"
	"strconv"
	"time"
)

func Test01() {
	var i int
	for {
		if i < 10 {
			i++
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		} else {
			break
		}
	}
}

func Test02() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
}

func Test03() {
	now := time.Now().Unix()
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
	fin := time.Now().Unix()
	fmt.Println(fin - now)
}
