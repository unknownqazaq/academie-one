package main

import (
	"github.com/01-edu/z01"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		number, _ := strconv.Atoi(os.Args[1])
		if number > 0 {
			for i := 1; i <= 9; i++ {
				z01.PrintRune(rune(i + '0'))
				z01.PrintRune('x')
				z01.PrintRune(rune(number + '0'))
				z01.PrintRune('=')
				sum := i * number
				sumStr := strconv.Itoa(sum) // Convert sum to a string
				for _, digit := range sumStr {
					z01.PrintRune(digit)
				}
				z01.PrintRune('\n')
			}
		}
	}
}
