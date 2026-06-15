package transform

import (
	"fmt"
	"gopwdutil/tools"

	"golang.org/x/crypto/bcrypt"
)

func Hash(ppwd *[]byte) {
	if ppwd == nil { return }
	hash, err := bcrypt.GenerateFromPassword(*ppwd, bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error: failed to hash.")
		return
	}
	fmt.Println("Password hashed")
	// Zero out the original password bytes before overwriting with the hash
	tools.Reset(ppwd)

	// Copy hash into ppwd
	*ppwd = append(*ppwd, hash...)

	// Zero out hash
	tools.Reset(&hash)
	fmt.Println("Hash erased")
}