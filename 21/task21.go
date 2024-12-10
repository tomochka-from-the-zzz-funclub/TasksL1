package main

import (
	"fmt"
)

type Vehicle interface {
	Move() string
}

type CarAdapter struct {
	*Car
}

func (c CarAdapter) Move() string {
	return c.Drive()
}

type Car struct {
}

func (c Car) Drive() string {
	return "vroom"
}

func MakeMove(v Vehicle) {
	fmt.Print(v.Move())
}

func main() {
	car := &Car{}
	adapter := CarAdapter{car}
	MakeMove(adapter)
}
