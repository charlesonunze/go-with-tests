package iteration

// Repeat function
func Repeat(character string) string {
	var repeated string

	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}

	return repeated
}
