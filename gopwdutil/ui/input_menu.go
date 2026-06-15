package ui

import (
	"fmt"
	"gopwdutil/tools"
	"strings"
)

func InputMenu() int {
	fmt.Println("========== Input Utils ===========")
	fmt.Println("1) Manual input")
	fmt.Println("2) Load from clipboard")
	fmt.Println("3) Generate random string")
	fmt.Println("0) Return to Main Menu")
	fmt.Println(strings.Repeat("-", 34))
	fmt.Print("Choice: ")
	
	return tools.GetChoice(3)
}

