package recur

func Test01(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Test01(n-1) + Test01(n-2)
	}
}

func Test02(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*Test02(n-1) + 1
	}
}

func Test03(n int) int {
	if n == 1 {
		return 1
	} else {
		return (1 + Test03(n-1)) * 2
	}
}
