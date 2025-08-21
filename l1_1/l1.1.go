package l1_1

import "fmt"

type Human struct {
	Name string
	Age  int
}
type Action struct {
	Human
}

func (h Human) Talk(word string) {
	fmt.Println(word)
}

func (h Human) IntroduceYourself() {
	fmt.Printf("I am %s and i have %d\n", h.Name, h.Age)
}
