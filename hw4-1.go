package main

import "fmt"

type Dog struct {
	breed  string
	age    int
	weight int
}

type Cat struct {
	breed  string
	age    int
	weight int
}

func (c Cat) talk() {
	fmt.Println("Myaw")
}

func (d Dog) talk() {
	fmt.Println("Woof")
}

type Animal interface {
	talk()
}

func animalTalk(animal Animal) {
	animal.talk()
}

func main() {
	dingo := Dog{breed: "k9", age: 3, weight: 20}
	murka := Cat{breed: "sphinx", age: 2, weight: 5}
	dingo.talk()
	murka.talk()
	animalTalk(dingo)
	animalTalk(murka)
}
