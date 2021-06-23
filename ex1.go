package main

import (
	"fmt"
)

func intToRome(n int) string {
	// nは入力された数字

	var rome string // ローマ数字
	var quo int     // 商
	//  var arr[2] string = [2]string {"Golang", "Java"}
	// var 変数名 [長さ]型 = [大きさ]型{初期値1, 初期値n}
	// 変数名 := [...]型{初期値１, 初期値n}

	intnum := [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romnum := [...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < 13; i++ {
		quo = n / intnum[i]
		for j := quo; j > 0; j-- {
			rome += romnum[i]
		}
		n = n % intnum[i]
	}

	return rome

}

func main() {
	var num int
	fmt.Scan(&num)

	intToRome(num)
}
