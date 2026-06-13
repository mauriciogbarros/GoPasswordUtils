package input

import (
	"fmt"
	"strings"
)

func Manual(ppwd *[]rune, minLength int, maxLength int) {
	if ppwd == nil { return }
	
	fmt.Println("========== Manual Input ==========")
	fmt.Printf("From %d to %d characters long\n", minLength, maxLength)
	fmt.Print("New password: ")

	var new_pwd []rune
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println("Error: failed to capture")
			return
		}
		if r == '\n' || r =='\r' {
			break
		}

		new_pwd = append(new_pwd, r)
	}

	if len(new_pwd) < minLength || len(new_pwd) > maxLength {
		fmt.Println("Error: invalid password length.")
		return
	}

	if strings.TrimSpace(string(new_pwd)) == "" {
		fmt.Println("Error: white space only.")
		return
	}

	for i := range *ppwd { (*ppwd)[i] = 0 }

	for i := 0; i < len(new_pwd); i++ {
		*ppwd = append(*ppwd, new_pwd[i])
	}

	for i := range(new_pwd) {
		new_pwd[i] = 0
	}
}