package cal

import (
	"fmt"
	"testing"
)

func TestAddUpper(t *testing.T) {
	fmt.Println("Test01")
	res := AddUpper(10)
	if res != 55 {
		t.Fatalf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
	}

	// 如果正确，输出日志
	t.Logf("AddUpper(10) 执行正确。。。")
}

func TestAddUpper1(t *testing.T) {
	fmt.Println("Test02")
	res := AddUpper(10)
	if res != 55 {
		t.Fatalf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
	}

	// 如果正确，输出日志
	t.Logf("AddUpper(10) 执行正确。。。")
}

func TestSub(t *testing.T) {
	fmt.Println("Test03")
	res := GetSub(10, 20)
	if res != -10 {
		t.Fatalf("GetSub(10, 20)执行错误，期望值=%v 实际值=%v\n", -10, res)
	}

	// 如果正确，输出日志
	t.Logf("GetSub(10, 20) 执行正确。。。")
}
