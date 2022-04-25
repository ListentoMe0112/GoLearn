package cal

func AddUpper(n int) int {
	var res int
	for i := 0; i < n; i++ {
		res += i
	}

	return res
}

func GetSub(n1 int, n2 int) int {
	return n1 - n2
}
