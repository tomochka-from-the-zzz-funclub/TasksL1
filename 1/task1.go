package main

import "fmt"

type Human struct {
	name string
}

func (h *Human) PrintName() {
	fmt.Println(h.name)
}

type Action struct {
	Human
}

func main() {
	h := Action{
		Human: Human{
			name: "fox",
		},
	}
	h.PrintName()
	h.Human.PrintName()

}
