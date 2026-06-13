package analysis

import "fmt"

func Strength(ppwd *[]rune) {
	if ppwd == nil { return }

	score := 0
	length := count_char(ppwd)

	scoreLength := 0
	switch {
		case length >= 128: scoreLength += 32
		case length >= 64:  scoreLength += 16
		case length >= 32:  scoreLength += 8
		case length >= 16:  scoreLength += 4
		case length >= 8:   scoreLength += 2
	}
	score += scoreLength

	scoreUpper 		:= count_upper(ppwd)
	scoreLower 		:= count_lower(ppwd)
	scoreNumeric 	:= count_numeric(ppwd)
	scoreSpecial 	:= count_special(ppwd)

	score += scoreUpper + scoreLower + scoreNumeric + scoreSpecial

	repeatedCount := len(getRepeated(ppwd))
	scoreRepeated := 0
	switch {
		case repeatedCount >= 4: scoreRepeated += 0
		case repeatedCount == 3: scoreRepeated += 1
		case repeatedCount == 2: scoreRepeated += 2
		case repeatedCount == 1: scoreRepeated += 4
		case repeatedCount == 0: scoreRepeated += 8
	}

	var label string
	switch {
		case score >= 16: label = "Very Strong"
		case score >= 10: label = "Strong"
		case score >= 4:  label = "Fair"
		default:          label = "Weak"
	}

	fmt.Printf("======> Strength analysis:\n")
	fmt.Printf("        Length score:    %d\n", scoreLength)
	fmt.Printf("        Uppercase score: %d\n", scoreUpper)
	fmt.Printf("        Lowercase score: %d\n", scoreLower)
	fmt.Printf("        Numeric score:   %d\n", scoreNumeric)
	fmt.Printf("        Special score:   %d\n", scoreSpecial)
	fmt.Printf("        --------------------\n")
	fmt.Printf("        Total score:     %d\n", score)
	fmt.Printf("======> Password strength: %s \n", label)
	fmt.Print("        Press Enter to continue... ")
	fmt.Scanln()
}