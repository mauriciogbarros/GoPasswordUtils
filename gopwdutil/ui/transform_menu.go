package ui

import (
	"fmt"
	"gopwdutil/tools"
	"strings"
)

func TransformMenu() int {
	fmt.Println("====== Transformation Tools ======")
	fmt.Println("1) Reverse password")
	fmt.Println("2) Scramble password")
	fmt.Println("3) Hash password")
	fmt.Println("4) Salt password")
	fmt.Println("0) Return to Main Menu")
	fmt.Println(strings.Repeat("-", 34))
	fmt.Print("Choice: ")

	return tools.GetChoice(4)
}