package clousure

import "strings"

func Adapter() func(int) (int, string) {
	var acc int = 10
	str := "hello"
	return func(x int) (int, string) {
		str += "$"
		acc = acc + x
		return acc, str
	}
}

func MakeSuffix(suffix string) func(string) string {
	m_suffix := suffix
	return func(fileName string) string {
		if strings.HasSuffix(fileName, m_suffix) {
			return fileName
		} else {
			var fullName string = fileName + m_suffix
			return fullName
		}
	}
}
