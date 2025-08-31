package l1_12

type SetStruct[T comparable] struct {
	set []T
}

func CreateSet[T comparable](arr []T) *SetStruct[T] {
	mapSet := make(map[T]bool)
	for _, v := range arr {
		mapSet[v] = true
	}
	newSet := &SetStruct[T]{set: make([]T, len(mapSet))}
	for v := range mapSet {
		newSet.set = append(newSet.set, v)
	}
	return newSet
}
