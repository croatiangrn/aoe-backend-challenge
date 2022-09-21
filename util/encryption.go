package util

import (
	"fmt"
	"strings"
)

const (
	asciiSpaceCharacter = 32
	asciiLowercaseA     = 97
	asciiLowerCaseZ     = 122
)

func DeeSeeChiffre(input string, key int) (string, error) {
	inputAsRune := []rune(strings.ToLower(input))

	var sb strings.Builder

	for _, singleRune := range inputAsRune {
		runeAsInt := int(singleRune)

		switch {
		case runeAsInt+key > asciiLowerCaseZ:
			diff := (runeAsInt + key) - asciiLowerCaseZ
			actualKey := (asciiLowercaseA + diff) - 1
			sb.WriteString(fmt.Sprintf("%c", actualKey))
		case runeAsInt == asciiSpaceCharacter:
			sb.WriteString(" ")
		default:
			sb.WriteString(fmt.Sprintf("%c", runeAsInt+key))
		}
	}

	return sb.String(), nil
}
