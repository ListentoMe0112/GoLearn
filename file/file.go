package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func Test01() {
	file, err := os.OpenFile("test1_go.txt", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("open file err = \n", err)
	}

	// fmt.Printf("file=%v\n", file)
	defer file.Close()

	str := "This line is written by [go]\n"
	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}

func Test02() {
	file, err := os.OpenFile("test1_go.txt", os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	defer file.Close()

	str := "Hello, go file io system\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}

func Test03() {
	file, err := os.OpenFile("test1_go.txt", os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	defer file.Close()

	str := "ABC!ENGLISH\n"
	writer := bufio.NewWriter(file)
	writer.WriteString(str)

	writer.Flush()
}

func Test04() {
	file, err := os.OpenFile("test1_go.txt", os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	defer file.Close()

	// 读取所有内容
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}

	// 写入内容
	str := "This line is written by go\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

func Test05() {
	file, err := os.OpenFile("test1_go.txt", os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	file1, err := os.OpenFile("test1_go_1.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	defer file.Close()
	defer file1.Close()
	// 读取所有内容
	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(file1)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		writer.WriteString(str)
	}
	writer.Flush()
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func Test06() {
	file, err := os.OpenFile("test1_cpp.txt", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("Open fail, err = ", err)
		return
	}

	reader := bufio.NewReader(file)

	var alphas, nums, spaces, others int

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for i := 0; i < len(str); i++ {
			if unicode.IsLetter(rune(str[i])) {
				alphas++
			} else if unicode.IsDigit(rune(str[i])) {
				nums++
			} else if str[i] == ' ' {
				spaces++
			} else {
				others++
			}
		}
	}

	fmt.Println(alphas, nums, spaces, others)
}
