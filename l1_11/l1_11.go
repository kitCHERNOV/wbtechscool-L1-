package l1_11

func SetIntersection(arr1, arr2 []int) (intersection []int) {
	valuesMap := make(map[int]int)
	for i := 0; i < len(arr1); i++ {
		valuesMap[arr1[i]] = 1
	}
	for i := 0; i < len(arr2); i++ {
		if _, ok := valuesMap[arr2[i]]; ok {
			intersection = append(intersection, arr2[i])
		}
	}
	return
}
