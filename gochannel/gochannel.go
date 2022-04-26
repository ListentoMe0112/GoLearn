package gochannel

import "fmt"

func Acc(n int) int {
	var sum int = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func GoRoutineAcc(sigChan chan bool, numChan chan int, resChan chan int) {
	for {
		num, ok := <-numChan
		if !ok {
			break
		}
		resChan <- Acc(num)
	}
	sigChan <- true
}

func Test01() {
	var sigChan chan bool = make(chan bool, 4)
	var numChan chan int = make(chan int, 8000)
	var resChan chan int = make(chan int, 2000)

	go func() {
		for i := 1; i <= 2000; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	for i := 0; i < 8; i++ {
		go GoRoutineAcc(sigChan, numChan, resChan)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-sigChan
		}
		close(sigChan)
		close(resChan)
	}()

	for {
		num, ok := <-resChan
		if !ok {
			break
		} else {
			fmt.Println(num)
		}
	}
}
