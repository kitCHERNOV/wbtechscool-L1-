package l1_17

func BinarySearch(arr []int, number int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == number {
			return mid
		}
		if arr[mid] < number {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
