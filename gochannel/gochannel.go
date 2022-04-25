package gochannel

import "fmt"

var myChan chan int = make(chan int, 2000)
var sigChan chan bool = make(chan bool, 8)
var resChan []int = make([]int, 8)

func WriteData() {
	for i := 0; i < 2000; i++ {
		myChan <- (i + 1)
	}
	close(myChan)
}

func Acc(i int) {
	for {
		v, ok := <-myChan
		if ok {
			resChan[i] = resChan[i] + v
		} else {
			sigChan <- false
			return
		}
	}
}

func Test01() {
	var sum int = 0
	go WriteData()
	for i := 0; i < 8; i++ {
		go Acc(i)
	}

	for {
		if len(sigChan) == 8 {
			close(sigChan)
			break
		}
	}

	for i := 0; i < 8; i++ {
		sum += resChan[i]
	}

	fmt.Println(sum)
}
