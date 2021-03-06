package utils

import "fmt"

func ReverseString(ori string) string {
	runes := []rune(ori)
	for from, to := 0, len(runes) - 1; from < to; from, to = from + 1, to - 1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

func TestVersion(){
	fmt.Println("测试版")
}