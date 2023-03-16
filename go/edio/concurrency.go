package main

import "fmt"

func deposit(balance *int, amount int) {
	*balance += amount //add amount to balance
}

func withdraw(balance *int, amount int) {
	*balance -= amount //subtract amount from balance
}

func main() {

	balance := 100

	go deposit(&balance, 10) //depositing 10

	withdraw(&balance, 50) //withdrawing 50

	fmt.Println(balance)

}

//package main
//
//import "fmt"
//
//func deposit(balance *int, amount int) {
//	*balance += amount //add amount to balance
//}
//
//func withdraw(balance *int, amount int) {
//	*balance -= amount //subtract amount from balance
//}
//
//func main() {
//
//	balance := 100
//
//	go deposit(&balance, 10) //depositing 10
//
//	withdraw(&balance, 50) //withdrawing 50
//
//	fmt.Println(balance)
//
//}

//package main
//
//import "fmt"
//
//func main() {
//	number := 0
//
//	go func() {
//		number++ //reading and modifying the value of 'number'
//	}()
//
//	fmt.Println(number) //reading the value of 'number'
//
//}
