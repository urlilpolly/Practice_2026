package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	str := "LVIII"
	out := 0
	str = strings.ToUpper(str)
	r := []rune(str)

	for i := 0; i < len(r); i++ {
		val, ok := nums[r[i]]
		if !ok {
			fmt.Printf("\nНедопустимый символ: %c", r[i])
		}
		if i < len(r)-1 {
			nextVal, ok := nums[r[i+1]]
			if !ok {
				fmt.Printf("\nНедопустимый символ: %c", r[i+1])
			}
			if val < nextVal {
				out -= val
				continue
			}
		}
		out += val
	}
	fmt.Println(out)
}
