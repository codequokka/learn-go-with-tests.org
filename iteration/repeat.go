package iteration

const repeatCount = 5

// Repeat returns 5 times repeated strings.
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}