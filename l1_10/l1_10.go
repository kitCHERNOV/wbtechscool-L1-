package l1_10

import (
	"fmt"
	"math"
)

/*
Реализация программы разбиения температурных значений по группам с шагом 10
решение представляет собой такой подход:
	- 0, 10, 20 и т.д. есть центры групп (класетров)
	- к группе по условию задачи относятся значения центр + 9.(9) [т.е. 9 и 9 в периоде]
*/

func groupCompositor(temperatureMap map[int][]float64, gettingArray []float64) {
	for _, value := range gettingArray {
		groupNumber := int(value - math.Mod(value, 10))
		temperatureMap[groupNumber] = append(temperatureMap[groupNumber], value)
	}
}

func MainFunc() {
	var temperature map[int][]float64 = make(map[int][]float64)
	tArr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groupCompositor(temperature, tArr)
	fmt.Println(temperature)
}
