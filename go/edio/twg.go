package main

import (
	"fmt"
	"strings"
)

type Person struct { // struct definition
	firstName string
	lastName  string
}

func upPerson(p *Person) { // function using struct as a parameter
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// 1- struct as a value type:
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	// 2 - struct as a pointer:
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward" // this is also valid
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	// 3 - struct as a literal:
	pers3 := &Person{"Chris", "Woodward"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}

//type struct1 struct { // struct definition
//	i1  int
//	f1  float32
//	str string
//}
//
//func main() {
//	ms := new(struct1) // making a struct1 type variable
//	// Filling fields of the struct of struct1 type
//	ms.i1 = 10
//	ms.f1 = 15.5
//	ms.str = "Chris"
//	fmt.Printf("The int is: %d\n", ms.i1)
//	fmt.Printf("The float is: %f\n", ms.f1)
//	fmt.Printf("The string is: %s\n", ms.str)
//	fmt.Println(ms)
//}
