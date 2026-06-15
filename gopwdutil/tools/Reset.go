package tools

func Reset(ppwd *[]byte) {
	if ppwd == nil { return }

	// Zero out each byte before shrinking to prevent lingering sensitive data in memory	
	for i := range *ppwd {
		(*ppwd)[i] = 0
	}

	// Shrink slice length without freeing the underlying array	
	*ppwd = (*ppwd)[:0]

}
