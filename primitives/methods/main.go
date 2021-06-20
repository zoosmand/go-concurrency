package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func (p *Person) Listen(what string) {
	fmt.Println("I'm listening to", what)
}

type Robot struct {
	Name string
}

func (r *Robot) Talk() {
	fmt.Println("I am a Robot, my name is", r.Name)
}

func main() {
	a := new(Android)
	a.Person.Name = "Leila"
	a.Talk()
	a.Listen("the Radio")

	r := new(Robot)
	r.Name = "Jack"
	r.Talk()
}
