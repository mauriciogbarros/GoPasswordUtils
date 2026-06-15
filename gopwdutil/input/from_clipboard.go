package input

import (
	"fmt"
	"gopwdutil/tools"
	"strings"

	"golang.design/x/clipboard"
)

func FromClipboard(ppwd *[]byte, minLength int, maxLength int) {
	if ppwd == nil { return }
	
	fmt.Println("====== Load from Clipboard =======")
	if err := clipboard.Init(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Clear clipboard first to ensure we read only what the user explicitly copies next	
	clipboard.Write(clipboard.FmtText, []byte(""))
	fmt.Print("Copy string to clipboard and press Enter to continue... ")
	fmt.Scanln()

	fmt.Println("Reading from clipboard...")
	data := clipboard.Read(clipboard.FmtText)

	if data == nil {
		fmt.Println("Error: clipboard is empty.")
		return
	}

	if strings.TrimSpace(string(data)) == "" {
		fmt.Println("Error: white space only.")
		clipboard.Write(clipboard.FmtText, []byte(""))
		return
	}

	if len(data) < minLength || len(data) > maxLength {
		fmt.Println("Error: invalid password length")
		clipboard.Write(clipboard.FmtText, []byte(""))
		return
	}

	tools.Reset(ppwd)

	for i := range data {
		*ppwd = append(*ppwd, data[i])
	}
	
	clipboard.Write(clipboard.FmtText, []byte(""))
	for i := range(data) { data[i] = 0 }
}