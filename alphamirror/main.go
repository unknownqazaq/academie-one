package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		args := []rune(os.Args[1])
		for i, char := range args {
			if char >= 'a' && char <= 'z' {
				args[i] = 'z' - char + 'a'
			}
			if char >= 'A' && char <= 'Z' {
				args[i] = 'Z' - char + 'A'

			}
			z01.PrintRune(args[i])
		}
		z01.PrintRune('\n')
	}

}
