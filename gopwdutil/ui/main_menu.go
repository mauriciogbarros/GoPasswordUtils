package ui

import (
	"fmt"
	"gopwdutil/tools"
	"strings"
)

func MainMenu(length int) int {
	fmt.Println("=========== Main Menu ============")
	fmt.Println("Password length:", length)
	fmt.Println("1) Input tools")
	if length > 0 {
		fmt.Println("2) Analysis tools")
		fmt.Println("3) Transformation tools")
		fmt.Println("4) Output options")
		fmt.Println("5) Reset current password")
	}
	fmt.Println("0) Exit")
	fmt.Println(strings.Repeat("-", 34))
	fmt.Print("Choice: ")

	choice := 0
	if length > 0 {
		choice = tools.GetChoice(5)
	} else {
		choice = tools.GetChoice(1)
	}
	fmt.Println()

	return choice
}
