package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sum int

	for _, day := range birdsPerDay {
		sum += day
	}

	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var birds []int
	var sum int

	if week == 1 {
		birds = birdsPerDay[:7]
	} else {
		start := (week * 6) - 5
		birds = birdsPerDay[start : start+7]
	}

	for _, day := range birds {
		sum += day
	}

	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index := range birdsPerDay {
		if index%2 == 0 {
			birdsPerDay[index]++
		}
	}

	return birdsPerDay
}
