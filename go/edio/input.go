package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}
	fmt.Printf("Your name is %s", input)

	// For Unix: test with delimiter "\n", for Windows: test with "\r\n"
	switch input {
	case "Philip\n":
		fmt.Println("Welcome Philip!")
	case "Chris\n":
		fmt.Println("Welcome Chris!")
	case "Ivo\n":
		fmt.Println("Welcome Ivo!")
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}
	// version 2:
	switch input {
	case "Philip\n":
		fallthrough
	case "Ivo\n":
		fallthrough
	case "Chris\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	// version 3:
	switch input {
	case "Philip\n", "Ivo\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}
}

//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//var inputReader *bufio.Reader
//var input string
//var err error
//
//func main() {
//	inputReader = bufio.NewReader(os.Stdin)
//	fmt.Println("Please enter some input: ")
//
//	input, err = inputReader.ReadString('\n')
//	if err == nil {
//		fmt.Printf("The input was: %s\n", input)
//	}
//}

//package main
//
//import "fmt"
//
//var (
//	firstName, lastName, s string
//	i                      int
//	g                      float32
//	input                  = "56.12 / 5212 / Go"
//	format                 = "%f / %d / %s"
//)
//
//func main() {
//	fmt.Println("Please enter your full name: ")
//	fmt.Scanln(&firstName, &lastName)
//	// fmt.Scanf("%s %s", &firstName, &lastName)
//	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
//	fmt.Sscanf(input, format, &g, &i, &s)
//	fmt.Println("From the string we read: ", g, i, s)
//}
