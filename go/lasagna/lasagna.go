package lasagna

// TODO: define the 'OvenTime' constant

const (
	OvenTime                = 40
	preparationTimePerLayer = 2
)

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	if actualMinutesInOven > OvenTime {
		panic("The food has burnt")
	}
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return preparationTimePerLayer * numberOfLayers
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	if numberOfLayers <= 0 {
		panic("cant bake an empty lasagna")
	}

	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
