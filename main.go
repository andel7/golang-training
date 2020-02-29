package main

import "fmt"

var stack []string

func push(str string) {
	stack = append(stack, str)
}

// Pop вернет последний добавленный в стек элемент
func pop() string {
	numOfElements := len(stack)
	// Когда стек будет пустым, он вернет пустую строку
	if numOfElements == 0 {
		return ""
	}
	popElem := stack[0]
	stack = append(stack[1:])
	return popElem
}

type Truck struct {
	companyName       string
	yearCreated       int
	trunkCapacity     int
	engineOn          bool
	windowsOpen       bool
	availableCapacity int
}

type Car struct {
	companyName       string
	yearCreated       int
	trunkCapacity     int
	engineOn          bool
	windowsOpen       bool
	availableCapacity int
}

func main() {
	truck1 := Truck{
		companyName:       "Scania",
		yearCreated:       2015,
		trunkCapacity:     400,
		engineOn:          false,
		windowsOpen:       false,
		availableCapacity: 100,
	}
	car1 := Car{
		companyName:       "mazda",
		yearCreated:       2018,
		trunkCapacity:     40,
		engineOn:          true,
		windowsOpen:       false,
		availableCapacity: 30,
	}
	fmt.Println(car1, car1.companyName)
	fmt.Println(truck1, truck1.windowsOpen)

	push("test1")
	push("test2")
	push("test3")
	fmt.Println(stack)
	pop()
	fmt.Println(stack)
	pop()
	fmt.Println(stack)
	pop()
	fmt.Println(stack)
}
