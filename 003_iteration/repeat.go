package iteration

import "strings"

func Repeat(character string, repetitions int) string {
	var repeated strings.Builder
	for range repetitions {
		repeated.WriteString(character)
	}
	return repeated.String()
}
