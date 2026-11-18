package iteration

import "strings"

const repeatCount = 5

func Repeat(text string) string {

	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(text)
	}
	return repeated.String()
}
