package l1_21

import "fmt"

// Буду адаптирвоать машину к рельсовому траспорту
// И есть Адаптер
type TrainCarPlatformAdapter interface {
	RailMovement()
}

type Car struct {
}

func (car *Car) DrivingForward() {
	fmt.Println("Автомобиль едет вперед")
}

type CarAdapter struct {
	*Car
}

func (ca *CarAdapter) RailMovement() {
	fmt.Println("Автомобиль поставили на жд платформу")
	ca.DrivingForward()

}

func NewCarAdapter(ca *Car) TrainCarPlatformAdapter {
	return &CarAdapter{ca}
}

// По логике поезд и так способен передвигаться по ж/д путям
type Train struct {
}

func (train *Train) RailMovement() {
	fmt.Println("Поезд двигается по железнодорожным путям")
}

func MainFunc() {
	train := new(Train)
	train.RailMovement()

	car := NewCarAdapter(&Car{})
	car.RailMovement()
}
