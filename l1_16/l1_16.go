package l1_16

import "math/rand"

// Реализация quick sort
func QuickSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	left, right := 0, len(ar)-1
	pivotIndex := rand.Int() % len(ar)

	ar[pivotIndex], ar[right] = ar[right], ar[pivotIndex]

	for i := 0; i < len(ar); i++ {
		if ar[i] < ar[right] {
			ar[i], ar[left] = ar[left], ar[i]
			left++
		}
	}

	ar[left], ar[right] = ar[right], ar[left]

	QuickSort(ar[:left])
	QuickSort(ar[left+1:])

	return
}
