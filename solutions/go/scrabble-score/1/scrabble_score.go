package scrabble

import (
	"strings"
)

var scores = map[int][]string{
	1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
	2:  {"D", "G"},
	3:  {"B", "C", "M", "P"},
	4:  {"F", "H", "V", "W", "Y"},
	5:  {"K"},
	8:  {"J", "X"},
	10: {"Q", "Z"},
}

func Score(word string) int {
	points := 0

	for _, char := range word {
		for score, chars := range scores {
			if contains(chars, string(char)) {
				points += score
			}
		}
	}

	return points
}

func contains(values []string, match string) bool {
	for _, value := range values {
		if strings.ToLower(value) == strings.ToLower(match) {
			return true
		}
	}

	return false
}
