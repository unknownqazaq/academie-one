package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) == 4 {
		a := Atoi(os.Args[1])
		b := Atoi(os.Args[3])

		operator := os.Args[2]

		switch operator {
		case "+":
			result := a + b
			if (result > a) == (b > 0) {
				sumStr := Itoa(result)
				for _, char := range sumStr {
					z01.PrintRune(char)
				}
				z01.PrintRune('\n')

			}
		case "-":
			result := a - b
			if (result < a) == (b > 0) {
				sumStr := Itoa(result)
				for _, char := range sumStr {
					z01.PrintRune(char)
				}
				z01.PrintRune('\n')

			}
		case "*":
			result := a * b
			if a == 0 || (result/a == b) {
				sumStr := Itoa(result)
				for _, char := range sumStr {
					z01.PrintRune(char)
				}
				z01.PrintRune('\n')

			}
		case "/":
			if b == 0 {
				fmt.Println("No division by 0")
			} else {
				result := a / b
				sumStr := Itoa(result)
				for _, char := range sumStr {
					z01.PrintRune(char)
				}
				z01.PrintRune('\n')

			}
		case "%":

			if b == 0 {
				fmt.Println("No modulo by 0")
			} else {
				result := a % b
				sumStr := Itoa(result)
				for _, char := range sumStr {
					z01.PrintRune(char)
				}
				z01.PrintRune('\n')

			}

		}
	}
}

func Atoi(s string) int {
	ans := 0
	for _, char := range s {
		if char < '0' || char > '9' {
			return ans
		} else {
			ans *= 10
			ans += int(char) - '0'
		}

	}
	return ans
}

func Itoa(d int) string {
	sign := ""
	result := ""
	if d < 0 {
		sign = "-"
		d = -d // Make d positive
	} else if d == 0 {
		return "0" // Special case for 0
	}

	for d > 0 {
		digit := d % 10
		result = string('0'+digit) + result
		d /= 10
	}
	return sign + result
}
