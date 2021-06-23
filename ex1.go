package main

import (
	"fmt"
)

func intToRome(n int) string {
	// nは入力された数字

	var rome string // ローマ数字
	var quo int     // 商

	if n > 0 {
		quo = n / 1000
	}

	for i := 1; i <= quo; i++ {
		rome += "M"
	}

	return rome

}

func main() {
	var num int
	fmt.Scan(&num)

	intToRome(num)
}
