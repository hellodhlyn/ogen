package internal

import (
	"strings"
	"unicode"
)

func splitText(text string, size int) (split string, left string) {
	splits := strings.Split(text, " ")
	lineSize := 0
	for idx, s := range splits {
		for _, c := range s {
			if unicode.Is(unicode.Hangul, c) {
				lineSize += 2
			} else {
				lineSize += 1
			}
		}
		if lineSize > size {
			return split, strings.Join(splits[idx:], " ")
		}
		split = split + " " + s
	}
	return
}
