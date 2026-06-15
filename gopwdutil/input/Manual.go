package input

import (
	"fmt"
	"gopwdutil/tools"
	"strings"
)

func Manual(ppwd *[]byte, minLength int, maxLength int) {
	if ppwd == nil { return }
	
	fmt.Println("========== Manual Input ==========")
	fmt.Printf("From %d to %d characters long\n", minLength, maxLength)
	fmt.Print("New password: ")

	var newPwd []byte
	for {
		b, err := tools.Reader.ReadByte()
		if err != nil {
			fmt.Println("Error: failed to capture")
			return
		}
		if b == '\n' || b =='\r' {
			break
		}

		newPwd = append(newPwd, b)
	}

	if len(newPwd) < minLength || len(newPwd) > maxLength {
		fmt.Println("Error: invalid password length.")
		for i := range newPwd { newPwd[i] = 0 }
		return
	}

	if strings.TrimSpace(string(newPwd)) == "" {
		fmt.Println("Error: white space only.")
		for i := range newPwd { newPwd[i] = 0 }		
		return
	}

	tools.Reset(ppwd)

	for i := 0; i < len(newPwd); i++ {
		*ppwd = append(*ppwd, newPwd[i])
	}

	for i := range(newPwd) {
		newPwd[i] = 0
	}
}