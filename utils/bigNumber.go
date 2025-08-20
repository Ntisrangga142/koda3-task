package utils

func BigNumber(arr []uint) uint {

	var bigNumber uint = 0

	for _, s := range arr {
		if bigNumber < s {
			bigNumber = s
		}
	}

	return bigNumber
}
