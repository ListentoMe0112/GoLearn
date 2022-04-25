package goroutine

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var (
	m    map[int]int = make(map[int]int, 10)
	lock sync.Mutex
)

func Test01() {
	for i := 1; i < 10; i++ {
		fmt.Println("[routine] hello world " + strconv.Itoa(i))
	}
}

func Test02() {
	go Test01()
	time.Sleep(time.Second)
	for i := 1; i < 10; i++ {
		fmt.Println("[main]hello world " + strconv.Itoa(i))
	}
}

func Test03() {
	num := runtime.NumCPU()
	fmt.Println("num = ", num)
	// 设置使用多个CPU
	runtime.GOMAXPROCS(num)
	Test02()
	// fmt.Println("num = ", num)
}

func calin(m *map[int]int, num int) {
	res := 1
	for i := 1; i <= num; i++ {
		res *= i
	}
	lock.Lock()
	(*m)[num] = res
	lock.Unlock()
}

func Test04() {
	for i := 1; i <= 10; i++ {
		go calin(&m, i)
	}

	time.Sleep(time.Second * 10)

	lock.Lock()
	for i, v := range m {
		fmt.Printf("map[%d] = %d\n", i, v)
	}
	lock.Unlock()
}
