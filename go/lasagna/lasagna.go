package lasagna

const OvenTime = 40

func RemainingOvenTime(time int) int {
	return OvenTime - time
}

const LayerPreparationTime = 2

func PreparationTime(layers int) int {
	return layers * LayerPreparationTime
}

func ElapsedTime(layers, time int) int {
	return PreparationTime(layers) + time
}
