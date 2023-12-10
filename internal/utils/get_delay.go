package utils

import "math"

// Calculate the delay between each job base on the number of elements
func GetDelay(numberOfFeed int) int {
	return int(math.Round(float64(60 / numberOfFeed)))
}
