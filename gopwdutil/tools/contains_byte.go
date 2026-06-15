package tools

func ContainsByte(bValue byte, refBytes []byte) bool {
	for _, bRef := range refBytes {
		if bValue == bRef { return true }
	}
	return false
}
