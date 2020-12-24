package iteration

import "strings"

func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}

func RepeatV2(character string, count int) (repeated string) {
	repeated = strings.Repeat(character, count)
	return
}
