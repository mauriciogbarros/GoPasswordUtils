package analysis

import (
	"fmt"
	"gopwdutil/tools"
	"strings"
)

func Count(ppwd *[]byte, choice int) {
	if ppwd == nil { return }

	count := getCount(ppwd, choice)
	if count == -1 {
		fmt.Println("Error: password is nil")
		return
	}

	countType := getCountType(choice)
	countUnit := getCountUnit(count, choice)

	fmt.Printf("======> %s count: %d %s\n", countType, count, countUnit)
	fmt.Print("        Press Enter to continue... ")
	fmt.Scanln()
}

func getCountType(choice int) string {
	countType := ""
	switch choice {
		case 1: countType = "Character"
		case 2: countType = "Word"
		case 3: countType = "Letter character"
		case 4: countType = "Uppercase character"
		case 5: countType = "Lowercase character"
		case 6: countType = "Numeric character"
		case 7: countType = "Special character"
	}

	return countType
}

func getCount(ppwd *[]byte, choice int) int {
	if ppwd == nil { return -1 }
	count := 0
	switch choice {
		case 1: count = countChar(ppwd)
		case 2: count = countWord(ppwd)
		case 3: count = countLetter(ppwd)
		case 4: count = countUpper(ppwd)
		case 5: count = countLower(ppwd)
		case 6: count = countNumeric(ppwd)
		case 7: count = countSpecial(ppwd)
	}

	return count
}

func getCountUnit(count int, choice int) string {
	unit := ""
	if count > 0 {
		switch choice {
			case 1: unit = "character"
			case 2: unit = "word"
			case 3: unit = "letter"
			case 4: unit = "uppercase letter"
			case 5: unit = "lowercase letter"
			case 6: unit = "numeric character"
			case 7: unit = "special character"
		}

		if count > 1 {
			unit += "s"
		}
	}

	return unit
}

func countChar(ppwd *[]byte) int {
	if ppwd == nil { return -1 }
	return len(strings.ReplaceAll(string(*ppwd), " ", ""))
}

func countWord(ppwd *[]byte) int {
	if ppwd == nil { return -1 }

	return len(strings.Fields(string(*ppwd)))
}

func countLetter(ppwd *[]byte) int {
	if ppwd == nil { return -1 }

	var letters = append(lowercase, uppercase...)

	count := 0
	for i := range *ppwd {
		if tools.ContainsByte((*ppwd)[i], letters) { count += 1 }
	}

	return count
}

func countUpper(ppwd *[]byte) int {
	if ppwd == nil { return -1 }

	count := 0
	for i := range *ppwd {
		if tools.ContainsByte((*ppwd)[i], uppercase) { count += 1 }
	}

	return count
}

func countLower(ppwd *[]byte) int {
	if ppwd == nil { return -1 }

	count := 0
	for i := range *ppwd {
		if tools.ContainsByte((*ppwd)[i], lowercase) { count += 1 }
	}

	return count
}

func countNumeric(ppwd *[]byte) int {
	if ppwd == nil { return -1 }

	count := 0
	for i := range *ppwd {
		if tools.ContainsByte((*ppwd)[i], numeric) { count += 1 }
	}

	return count
}

func countSpecial(ppwd *[]byte) int {
	if ppwd == nil { return -1 }

	count := 0
	for i := range *ppwd {
		if tools.ContainsByte((*ppwd)[i], special) { count += 1 }
	}

	return count
}