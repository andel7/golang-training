package main

import (
	"fmt"
	"math/big"
)

func checkEven(num int) bool {
	return checkMod(num, 2)
}
func checkModThree(num int) bool {
	return checkMod(num, 3)
}
func checkMod(num, del int) bool {
	return num%del == 0
}
func fibonachi() {
	temp := big.NewInt(0)
	last := big.NewInt(1)
	previous := big.NewInt(0)
	for total := 2; total < 100; total++ {
		fmt.Println("current:", total, last)
		temp = last
		last = last.Add(previous, last)
		previous = temp
	}
}
func main() {
	fmt.Println("2", checkEven(2))
	fmt.Println("3", checkEven(3))
	fmt.Println("3", checkModThree(3))
	fmt.Println("4", checkModThree(4))
	fibonachi()
}
