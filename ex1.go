package main

import (
	"fmt"
)

// #include <iostream>

// using namespace std;

// string introm(int a) {
//   int intnum[] = {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
//   const char *romnum[] = {"M",  "CM", "D",  "CD", "C",  "XC", "L",
//                           "XL", "X",  "IX", "V",  "IV", "I"};
//   string s;
//   int i = 0;
//   while (a) {
//     while (a / intnum[i]) {
//       s += romnum[i];
//       a -= intnum[i];
//     }
//     i++;
//   }
//   return s;
// }

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

	// 	// 千の位を扱う
	// 	if n > 0 {
	// 		quo = n / 1000
	// 	}

	// 	for i := 1; i <= quo; i++ {
	// 		rome += "M"
	// 	}

	// 	// 百の位を扱う
	// 	n = n % 1000
	// 	if n > 0 {
	// 		quo = n / 1000
	// 	}

	// 	for i := 1; i <= quo; i++ {
	// 		rome += "M"
	// 	}

	return rome

}

func main() {
	var num int
	fmt.Scan(&num)

	intToRome(num)
}
