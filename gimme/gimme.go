package gimme

import "sort"

func Gimme(array [3]int) int {
	if (array[0] > array[1] && array[0] < array[2]) || (array[0] < array[1] && array[0] > array[2]) {
		return 0
	}
	if (array[1] > array[0] && array[1] < array[2]) || (array[1] < array[0] && array[1] > array[2]) {
		return 1
	}
	return 2
}

func GimmeV2(array [3]int) int {
	temp := array
	sort.Ints(temp[:])
	middle := temp[1]
	for i, v := range array {
		if middle == v {
			return i
		}
	}
	return -1
}
