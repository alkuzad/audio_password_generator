package nato

import (
	"strings"

	"log"
)

// natoMap maps each letter to its NATO phonetic equivalent.
var NatoMap = map[rune]string{
	'a': "Alfa", 'b': "Bravo", 'c': "Charlie", 'd': "Delta",
	'e': "Echo", 'f': "Foxtrot", 'g': "Golf", 'h': "Hotel",
	'i': "India", 'j': "Juliett", 'k': "Kilo", 'l': "Lima",
	'm': "Mike", 'n': "November", 'o': "Oscar", 'p': "Papa",
	'q': "Quebec", 'r': "Romeo", 's': "Sierra", 't': "Tango",
	'u': "Uniform", 'v': "Victor", 'w': "Whiskey", 'x': "X-ray",
	'y': "Yankee", 'z': "Zulu",
	'0': "Zero", '1': "One", '2': "Two", '3': "Three",
	'4': "Four", '5': "Five", '6': "Six", '7': "Seven",
	'8': "Ait", '9': "Niner",
}

// ToNato takes a string and returns a slice of strings with each letter
// converted to its NATO phonetic alphabet equivalent.
func ToNato(input string) []string {
	var result []string
	for _, char := range strings.ToLower(input) {
		if code, ok := NatoMap[char]; ok {
			result = append(result, code)
		} else {
			log.Fatalf("Missing mapping for %v", char)
		}
	}
	return result
}
