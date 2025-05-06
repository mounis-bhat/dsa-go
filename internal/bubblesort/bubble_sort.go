package bubblesort

func BubbleSort(arr *[]int) {
	for i := range len(*arr) - 1 {
		swapped := false
		for j := range len(*arr) - 1 - i {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
