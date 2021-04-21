package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	// isi manual
// 	fmt.Println(calculate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15))

// 	// isi dengan slice
// 	number := []int{11, 22, 33, 44, 55, 66, 77, 88, 99, 101}
// 	var avg = calculate(number...)
// 	var msg = fmt.Sprint("Rata-Rata nilai : ", avg)

// 	fmt.Println(msg)

// 	var hobbies = []string{"sleeping", "eating"}
// 	yourHobbies("wick", hobbies...)
// 	yourHobbies("wick", "sleeping", "eating")
// }

// func calculate(numbers ...int) (avg float64) {
// 	var total int = 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	avg = float64(total) / float64(len(numbers))
// 	return
// }

// func yourHobbies(name string, hobbies ...string) {
// 	var hobbiesAsString = strings.Join(hobbies, ", ")

// 	fmt.Printf("Hello, my name is: %s\n", name)
// 	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
// }
