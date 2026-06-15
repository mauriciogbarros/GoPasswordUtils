package output

import (
	"fmt"
)

func ToScreen(ppwd *[]byte) {
	if ppwd == nil { return }
	
	fmt.Println("Password:", string(*ppwd))
}