package analysis

import (
	"fmt"
)

func Repeated(ppwd *[]byte) {
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

func getRepeated(ppwd *[]byte) map[byte]int {
	if ppwd == nil { return nil }

	repeated := map[byte]int{}
	for _, b := range *ppwd {
		if b != 0 {	repeated[b]++	}
	}

	for b, n := range repeated {
		if n == 1 {
			delete(repeated, b)
		}
	}

	return repeated
}