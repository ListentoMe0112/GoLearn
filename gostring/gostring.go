package gostring

import (
	"fmt"
	"strconv"
	"strings"
)

func LenTest() {
	// 统计字符串的长度，按字节len
	// golang统一采用utf-8辩吗，ascii的字符占用1个字节，汉字占用3个字节
	str := "hello背景"
	fmt.Println("str len = ", len(str))
}

func RuneTest() {
	// 统计字符串的长度，按字节len
	// golang统一采用utf-8辩吗，ascii的字符占用1个字节，汉字占用3个字节
	str := "hello背景"
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c, %d\n", r[i], i)
	}
}

func AtoiTest() {
	str := "-123das"
	n, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("[error]", err)
	} else {
		fmt.Println(n)
	}
}

func ItoaTest() {
	fmt.Println(strconv.Itoa(-312789))
}

func ByteTest() {
	var bytes = []byte("hello go")
	fmt.Println(bytes)
	fmt.Printf("str = %v\n", string(bytes))
}

func IntTest() {
	var x int64 = 1023
	fmt.Println(strconv.FormatInt(x, 16))
}

func ContainsTest() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "mary"))
}

func CountTest() {
	fmt.Println(strings.Count("seafood", "foofd"))
	fmt.Println(strings.Count("seafood", "o"))
}

func EqualFoldTest() {
	fmt.Println(strings.EqualFold("abc", "ABCD"))
}

func IndexTest() {
	fmt.Println(strings.Index("abcABC", "ABCD"))
}

func LastIndexTest() {
	fmt.Println(strings.LastIndex("abcABCABCDABCD", "ABCD"))
}

func ReplaceTest() {
	fmt.Println(strings.Replace("abcABCABCDABCD", "ABCD", "dhjask", -1))
}

func SplitTest() {
	fmt.Println(strings.Split("Hello, Test, Test, Test", ","))
}

func ToLowerTest() {
	fmt.Println(strings.ToLower("Hello, Test, Test, Test"))
}

func ToUpperTest() {
	fmt.Println(strings.ToUpper("Hello, Test, Test, Test"))
}

func TrimSpaceTest() {
	fmt.Println(strings.TrimSpace("   dhasjk   "))
}

func TrimTest() {
	fmt.Println(strings.Trim("---!dha sjk   ---", "- "))
}

func HasPrefixTest() {
	fmt.Println(strings.HasPrefix("https://---!dha sjk   ---", "https://"))
}
