package utils

func Contains[T comparable](s []T, item T) bool {
	for _, i := range s {
		if i == item {
			return true
		}
	}
	return false
}

func Difference[T comparable](s []T, diffFrom []T) []T {
	var diffValues []T
	for _, i := range s {
		valueExists := false
		for _, j := range diffFrom {
			if i == j {
				valueExists = true
				break
			}
		}
		if !valueExists {
			diffValues = append(diffValues, i)
		}
	}
	return diffValues
}

func UnifySlice[T comparable](s []T) []T {
	keys := make(map[T]bool)
	setSlice := []T{}
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			setSlice = append(setSlice, entry)
		}
	}
	return setSlice
}
