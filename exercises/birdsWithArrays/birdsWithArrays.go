package main

// Sauce: https://exercism.org/tracks/go/exercises/bird-watcher/edit

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var count int
	for i := range birdsPerDay {
		count += birdsPerDay[i]
	}
	return count
}

func getStartAndEndIndeces(week int) (int, int) {
	var start, end int
	if week == 1 {
		start = 0
		end = 7
	} else {
		start = (week * 7) - 7
		end = (week * 7)
	}
	return start, end
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var count int
	start, end := getStartAndEndIndeces(week)

	for i := start; i < end; i++ {
		count += birdsPerDay[i]
	}
	return count
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i = i + 2 {
		birdsPerDay[i] = birdsPerDay[i] + 1
	}
	return birdsPerDay
}
