package unittest

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m *Monster) Store() {
	// 序列化数据
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Serilize Fail, err = ", err)
	}
	fmt.Println("After Serilize: ", string(data))

	// 打开文件
	file, err := os.OpenFile("monster.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	defer file.Close()

	// 写入内容
	writer := bufio.NewWriter(file)
	writer.WriteString(string(data) + "\n")
	writer.Flush()
}

func (m *Monster) ReStore() {

	// 打开文件
	file, err := os.OpenFile("monster.txt", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	defer file.Close()

	// 写入内容
	reader := bufio.NewReader(file)
	str, _ := reader.ReadString('\n')
	json.Unmarshal([]byte(str), m)
}
