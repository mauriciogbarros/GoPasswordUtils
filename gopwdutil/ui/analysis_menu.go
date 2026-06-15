package ui

import (
	"fmt"
	"gopwdutil/tools"
	"strings"
)

func AnalysisMenu() int {
	fmt.Println("======== Analysis Utils ==========")
	fmt.Println("1) Character count")
	fmt.Println("2) Word count")
	fmt.Println("3) Letter count")
	fmt.Println("4) Upper count")
	fmt.Println("5) Lower count")
	fmt.Println("6) Numeric count")
	fmt.Println("7) Special count")
	fmt.Println("8) Repeated counts")
	fmt.Println("9) Strength analysis")
	fmt.Println("0) Return to Main Menu")
	fmt.Println(strings.Repeat("-", 34))
	fmt.Print("Choice: ")
	
	return tools.GetChoice(9)
}
