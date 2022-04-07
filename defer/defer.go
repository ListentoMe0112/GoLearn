package defer_test

import "fmt"

func sum(n1 int, n2 int) int {
	// 执行到defer时会将后面的语句压入栈中，相应的值也会入栈
	// 函数执行完毕后从defer栈中依次弹出执行
	defer fmt.Println("ok1 n1 = ", n1)
	defer fmt.Println("ok2 n2 = ", n2)
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res = ", res)
	return res
}

func Test01() {
	fmt.Println(sum(10, 20))
}

// defer主要用于函数退出后释放资源

/*
func openFile() {
	file = openfile("name")
	defer file.close()
	// normal code
}

func openConnect() {
	connect = openDatabase()
	defer connect.close()
	// normal code
}
*/
