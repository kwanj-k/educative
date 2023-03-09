package main

import (
	"fmt"
)

func Season(month int) string {
	switch month {
	case 12, 1, 2:
		return "Winter"
	case 3, 4, 5:
		return "Spring"
	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Autumn"
	default:
		return "Season unknown"
	}
}

func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func trace(s string) {
	fmt.Println("entering:", s)
} // entering func.
func untrace(s string) {
	fmt.Println("leaving:", s)
} // leaving func.

func a() {
	trace("a")
	defer untrace("a") // untracing via defer
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b") // untracing via defer
	fmt.Println("in b")
	a()
}

func f(a [3]int) { fmt.Println(a) } // accepts copy

func fp(a *[3]int) { fmt.Println(a) } // accepts pointer

func Filter(s []int, fn func(int) bool) []int {
	var p []int
	for _, i := range s {
		if fn(i) {
			p = append(p, i)
		}
	}
	return p
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = Filter(s, even)
	fmt.Println(s)
	// var b bytes.Buffer

	//b.WriteString("ABC")
	//b.WriteString("DEF")
	//
	//fmt.Fprintf(&b, " A number: %d, a string: %v\n", 10, "bird")
	//b.WriteString("[DONE]")
	//
	//fmt.Println(&b)

	// var ar [3]int
	//f(ar)   // passes a copy of ar
	// fp(&ar) // passes a pointer to ar
	//var arr1 [5]int
	//
	//for i := 0; i < len(arr1); i++ {
	//	arr1[i] = i * 2
	//}
	//
	//fmt.Println(arr1)
	//b()
	//n := 0
	//reply := &n
	//fmt.Println("Before M:", n)
	//fmt.Println("Before M:", *reply)
	//
	//Multiply(10, 5, reply)
	//
	//fmt.Println("Multiply:", *reply) // Multiply: 50
	//fmt.Println("Multiply:", n)      // Multiply: 50

	//i := 0
	//HERE: // adding a label for a location
	//fmt.Printf("%d \n", i)
	//i++
	//if i == 5 {
	//	return
	//}
	//goto HERE // goto HERE

	//LABEL1:                       // adding a label for location
	//for i := 0; i <= 5; i++ { // outer loop
	//	for j := 0; j <= 5; j++ { // inner loop
	//		if j == 4 {
	//			continue LABEL1 // jump to the label
	//		}
	//		fmt.Printf("i is: %d, and j is: %d\n", i, j)
	//	}
	//}
	//for {
	//	fmt.Println("I am a God!!!!!!")
	//}
	//fmt.Println(Season(1))
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//str := "Go is a beautiful language!"
	//for pos, char := range str {
	//	fmt.Printf("Postion: %d and character is: %c \n", pos, char)
	//}
	//fmt.Printf("The length of str is: %d\n", len(str))
	//
	//// for loop to find character at each position
	//for ix := 0; ix < len(str); ix++ {
	//	fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	//}
	//str2 := "Chinese: 日本語"
	//fmt.Printf("The length of str2 is: %d\n", len(str2))
	//
	//// for loop to find character at each position
	//for ix := 0; ix < len(str2); ix++ {
	//	fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	//}
}
