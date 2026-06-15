package ui

import (
	"fmt"
	"gopwdutil/tools"
	"strings"
)

func OutputMenu() int {
		fmt.Println("========== Output Utils ==========")
		fmt.Println("1) To screen")
		fmt.Println("2) To clipboard")
		fmt.Println("3) To file")
		fmt.Println("0) Return to Main Menu")
		fmt.Println(strings.Repeat("-", 34))
		fmt.Print("Choice: ")

		return tools.GetChoice(3)
}
