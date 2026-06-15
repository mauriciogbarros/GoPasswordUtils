package input

import (
	"crypto/rand"
	"fmt"
	"gopwdutil/tools"
	"math/big"
	"strconv"
	"strings"
)

func Random(ppwd *[]byte, minLength int, maxLength int) {
	if ppwd == nil { return }

	fmt.Println("======== Generate Random =========")
	fmt.Printf("min %d, max %d characters\n", minLength, maxLength)
	fmt.Print("Enter desired length: ")

	line, _ := tools.Reader.ReadString('\n')
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

	tools.Reset(ppwd)
	for i := 0; i < length; i++ {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			fmt.Println("Error: randomizing error")
			return
		}
		*ppwd = append(*ppwd, characters[j.Int64()])
	}
}