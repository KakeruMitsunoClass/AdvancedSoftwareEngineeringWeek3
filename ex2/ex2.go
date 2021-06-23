package myPackage

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func romaToInt(s string) int {

	a := [...]string{"IV", "IX", "XL", "XC", "CD", "CM"}
	b := [...]string{"IIII", "VIIII", "XXXX", "LXXXX", "CCCC", "DCCCC"}

	for j := 0; j < 6; j++ {
		s = strings.Replace(s, a[j], b[j], -1)

	}
	n := 0
	number := [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roma := [...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < utf8.RuneCountInString(s); i++ {
		for j := 0; j < 13; j++ {
			if s[i:i+1] == roma[j] {
				n += number[j]
			}
		}
	}
	if n > 4000 {
		return -1
	} else {
		return n
	}
}

func main() {
	var s string
	fmt.Scan(&s)
	num := romaToInt(s)
	fmt.Printf("%d\n", num)

}
