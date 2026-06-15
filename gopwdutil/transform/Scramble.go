package transform

import (
	"crypto/rand"
	"fmt"
	"gopwdutil/tools"
	"math/big"
)

func Scramble(ppwd *[]byte) {
	if ppwd == nil { return }

	fmt.Println("Scrambling ...")
	length := len(*ppwd)
	scrambled := []byte{}
	for i := 0; i < length; i++ {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(length - i)))
		if err != nil {
			fmt.Println("Error: randomizing error")
			return
		}

		// Pick a random index from the remaining unselected bytes, then remove it from the pool
		idx := j.Int64()
		scrambled = append(scrambled, (*ppwd)[idx])

		// remove selected byte from pool
		*ppwd = append((*ppwd)[:idx], (*ppwd)[idx+1:]...)
	}

	// Copy scrambled to ppwd
	*ppwd = append(*ppwd, scrambled...)

	// Zero out scrambled
	tools.Reset(&scrambled)
	fmt.Println("Scrambled password:", string(*ppwd))
	fmt.Print("Press Enter to continue ...")
	fmt.Scanln()
}