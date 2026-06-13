package input

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func Random(ppwd *[]rune, minLength int, maxLength int) {
	if ppwd == nil { return }

	fmt.Println("======== Generate Random =========")
	fmt.Printf("min %d, max %d characters\n", minLength, maxLength)
	fmt.Print("Enter desired length: ")

	line, _ := reader.ReadString('\n')
	raw := strings.TrimRight(line, "\r\n")
	length, err := strconv.Atoi(raw)
	if err != nil {
		fmt.Printf("Error: Please enter a number, got %q.\n", line)
		return
	}

	if length < minLength || length > maxLength {
		fmt.Println("Error: invalid length")
		return
	}

	for i := range *ppwd { (*ppwd)[i] = 0 }
	for i := 0; i < length; i++ {
		*ppwd = append(*ppwd, characters[rand.Intn(len(characters))])
	}
}