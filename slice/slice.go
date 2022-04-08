package slice

import "fmt"

func Test01() {
	var intArr [5]int = [...]int{1, 22, 3, 6, 9}
	// [begin, end)
	slice := intArr[1:3]
	fmt.Println(intArr)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Printf("%p %p\n", &slice, &intArr[1])
	slice[1] = 33
	fmt.Println(intArr)

	var slice1 []int = make([]int, 4, 10)
	fmt.Println(slice1, len(slice1), cap(slice1))
	var slicef []float64 = make([]float64, 10, 20)
	fmt.Println(slicef, len(slicef), cap(slicef))
	fmt.Printf("%T\n", slicef)

	for k, v := range slice1 {
		fmt.Println(k, v)
	}

	sliceNew := append(slice1, 200, 300)
	sliceNew = append(sliceNew, sliceNew...)
	fmt.Println(sliceNew)

	var a []int = []int{1, 2, 3, 4, 5}

	var sliceA = make([]int, 10)

	copy(sliceA, a)

	fmt.Println(sliceA)
	fmt.Printf("%T\n", a)
}

func fbn(n int) []int {
	ret := make([]int, n, 2*n)
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			ret[i] = 1
		} else {
			ret[i] = ret[i-1] + ret[i-2]
		}
	}

	return ret
}

func Test02() {
	fmt.Println(fbn(10))
}
