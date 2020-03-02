package main

import (
	"fmt"
	"sort"
)

type phoneRecord struct {
	name  string
	phone int
}
type phonesBook []phoneRecord

func (p phonesBook) ViewList() {
	for i, phone := range p {
		fmt.Printf("\t %v) %v \n", i, phone)
	}
}

func (p phonesBook) Len() int {
	return len(p)
}
func (p phonesBook) Less(i, j int) bool {
	return p[i].name < p[j].name
}
func (p phonesBook) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var myPhoneBook phonesBook
	myPhoneBook = append(myPhoneBook, phoneRecord{"Petya", 27128282910})
	myPhoneBook = append(myPhoneBook, phoneRecord{"Vasya", 12344543123})
	myPhoneBook = append(myPhoneBook, phoneRecord{"Alex", 1231231232})
	fmt.Println(myPhoneBook)
	sort.Sort(myPhoneBook)
	fmt.Println(myPhoneBook)
}
