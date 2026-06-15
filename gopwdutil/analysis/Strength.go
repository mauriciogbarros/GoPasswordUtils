package analysis

import "fmt"

func Strength(ppwd *[]byte) {
	if ppwd == nil { return }

	score := 0
	length := countChar(ppwd)

	// Score based on length tiers; longer passwords earn exponentially more points	
	scoreLength := 0
	switch {
		case length >= 72:  scoreLength += 32
		case length >= 36:  scoreLength += 16
		case length >= 18:  scoreLength += 8
		case length >= 14:  scoreLength += 4
		case length >= 8:   scoreLength += 2
	}
	score += scoreLength

	scoreUpper 		:= countUpper(ppwd)
	scoreLower 		:= countLower(ppwd)
	scoreNumeric 	:= countNumeric(ppwd)
	scoreSpecial 	:= countSpecial(ppwd)

	score += scoreUpper + scoreLower + scoreNumeric + scoreSpecial

	// Penalize passwords with many repeated characters; fewer repeats = higher score	
	repeatedCount := len(getRepeated(ppwd))
	scoreRepeated := 0
	switch {
		case repeatedCount >= 4: scoreRepeated += 0
		case repeatedCount == 3: scoreRepeated += 1
		case repeatedCount == 2: scoreRepeated += 2
		case repeatedCount == 1: scoreRepeated += 4
		case repeatedCount == 0: scoreRepeated += 8
	}
	score += scoreRepeated

	// Strength label thresholds are relative to the max achievable score	
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