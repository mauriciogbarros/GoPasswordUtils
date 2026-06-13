package ui

import (
	"fmt"
	"gopwdutil/tools"
	"strings"
)

func MainMenu() {
	const minLength = 8
	const maxLength = 256

	fmt.Printf("Min length: %d, Max length: %d\n", minLength, maxLength)
	pwd := make([]rune, 0, maxLength)
	for i := range pwd { pwd[i] = 0 }
	exit := false
	choice := 0
	length := 0
	
	for !exit {
		length = len(pwd)
		
		fmt.Println("=========== Main Menu ============")
		tools.DisplayCurrentPassword(&pwd)
		fmt.Println("1) Input tools")
		if length > 0 {
			fmt.Println("2) Analysis tools")
			fmt.Println("3) Transformation tools")
			fmt.Println("4) Output options")
			fmt.Println("5) Reset current password")
		}
		fmt.Println("0) Exit")
		fmt.Println(strings.Repeat("-", 34))

		if length > 0 {
			choice = tools.GetMenuChoice(5)
		} else {
			choice = tools.GetMenuChoice(1)
		}
		fmt.Println()

		if choice == 0 {
			exit = true
			fmt.Println("Erasing password...")
			for i := range pwd { pwd[i] = 0}
			fmt.Println()
			fmt.Println("Password erased.")
		} else {
			if length > 0 {
				switch choice {
					case 1:	Input(&pwd, minLength, maxLength)
					case 2:	Analysis(&pwd)
					case 3: Transform(&pwd)
					case 4: Output(&pwd)
					case 5: tools.Reset(&pwd)
					default:
				}			
			} else {
				switch choice {
					case 1: Input(&pwd, minLength, maxLength)
					default:
				}
			}

		}
	}
}
