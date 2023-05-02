package next_letter

// Find the next letter for a menu option
func FindNext(l string) string {
	bits := []byte(l)
	currentLetter := bits[len(l)-1]

	switch {
	case currentLetter >= 'a' && currentLetter < 'z':
		next := currentLetter + 1
		if len(bits) > 1 {
			for i := 0; i < len(bits); i++ {
				bits[i] = next
			}
			break
		}
		bits[0] = next

	case currentLetter == 'z':
		next := currentLetter - 25

		for i := 0; i < len(bits); i++ {
			bits[i] = next
		}
		bits = append(bits, 'a')

	case currentLetter >= 'A' && currentLetter < 'Z':
		next := currentLetter + 1
		if len(bits) > 1 {
			for i := 0; i < len(bits); i++ {
				bits[i] = next
			}
			break
		}
		bits[0] = next

	case currentLetter == 'Z':
		next := currentLetter - 25

		for i := 0; i < len(bits); i++ {
			bits[i] = next
		}
		bits = append(bits, 'Z')
	}

	return string(bits)
}
