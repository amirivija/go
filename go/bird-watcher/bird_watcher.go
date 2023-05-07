package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalBirds := 0
	for _, v := range birdsPerDay {
		totalBirds += v
	}

	return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	totalWeeks := len(birdsPerDay) / 7

	if week <= 0 || week > totalWeeks {
		panic("week is less than 0 or greater than total weeks")
	}

	birdsInWeek := 0

	for i := (week - 1) * 7; i < (week-1)*7+7; i++ {
		birdsInWeek += birdsPerDay[i]
	}

	return birdsInWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}

	return birdsPerDay
}
