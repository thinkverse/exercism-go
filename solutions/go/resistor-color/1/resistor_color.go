package resistorcolor

import "sort"

var colors = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Colors should return the list of all colors.
func Colors() []string {
	keys := make([]string, 0, len(colors))

	for key := range colors {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return colors[keys[i]] < colors[keys[j]]
	})

	return keys
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	if resistance, ok := colors[color]; ok {
		return resistance
	}

	return 0
}
