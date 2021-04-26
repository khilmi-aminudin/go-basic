package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// }

// func main() {
// 	var khilmi Person = Person{"khilmi", 23}
// 	var fathi = &khilmi
// 	var faza = &khilmi
// 	fathi.age = 9
// 	fmt.Println(&khilmi)
// 	fmt.Println(fathi)
// 	fmt.Println(&faza)
// }

// // func main() {
// // 	var numberA int = 4
// // 	var numberB *int = &numberA

// // 	fmt.Println("numberA (value)   :", numberA)  // 4
// // 	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

// // 	fmt.Println("numberB (value)   :", *numberB) // 4
// // 	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

// // 	fmt.Println("")

// // 	numberA = 5

// // 	fmt.Println("numberA (value)   :", numberA)
// // 	fmt.Println("numberA (address) :", &numberA)
// // 	fmt.Println("numberB (value)   :", *numberB)
// // 	fmt.Println("numberB (address) :", numberB)

// // 	var number = 4
// // 	fmt.Println("before :", number) // 4

// // 	change(&number, 10)
// // 	fmt.Println("after  :", number) // 10
// // }

// // func change(origin *int, value int) {
// // 	*origin = value
// // }

// //  memory
// //    |
// //   nilai
// //    |
// // referencer
