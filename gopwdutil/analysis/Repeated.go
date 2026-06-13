package analysis

import (
	"fmt"
)

func Repeated(ppwd *[]rune) {
	if ppwd == nil { return }

	fmt.Println("======> Repeated:")
	
	repeated := getRepeated(ppwd)
	if len(repeated) == 0 {
		fmt.Println("        No repeated characters.")
	} else {
		for r, n := range repeated {
			if n > 1 {
				fmt.Printf("        '%c': %d times\n", r, n)
			}
		}
	}

	fmt.Print("        Press Enter to continue... ")
	fmt.Scanln()
}

func getRepeated(ppwd *[]rune) map[rune]int {
	repeated := map[rune]int{}
	for _, r := range *ppwd {
		if r != 0 {
			repeated[r]++
		}
	}

	for r, n := range repeated {
		if n == 1 {
			delete(repeated, r)
		}
	}

	return repeated
}