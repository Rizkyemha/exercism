package lasagna

// TODO: define the 'OvenTime' constant

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	const OvenTime = 20
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	const PreparationTimePerLayer = 2
	return numberOfLayers * PreparationTimePerLayer
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
