package main

// import "fmt"

// // Iterasi atau perulangan di Go hanya ada for

// func main() {

// 	// regular for`
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("Angka", i)
// 	}

// 	// for just condition
// 	var i = 0

// 	for i < 5 {
// 		fmt.Println("Angka", i)
// 		i++
// 	}

// 	// for with no condition
// 	var it = 0

// 	for {
// 		fmt.Println("Angka", it)

// 		it++
// 		if it == 5 {
// 			break
// 		}
// 	}

// 	// for with break / continue
// 	// break : untuk menghentikan perulangan
// 	// continue : untuk melewati perulangan

// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 1 {
// 			continue
// 		}

// 		if i > 8 {
// 			break
// 		}

// 		fmt.Println("Angka", i)
// 	}

// 	// nested for
// 	for i := 0; i < 5; i++ {
// 		for j := i; j < 5; j++ {
// 			fmt.Print(j, " ")
// 		}

// 		fmt.Println()
// 	}

// 	// labeling untuk break / continue
// 	// outerLoop:
// 	for i := 0; i < 5; i++ {
// 		for j := 0; j < 5; j++ {
// 			if i == 3 {
// 				// break outerLoop
// 				fmt.Println("Break di comment")
// 			}
// 			fmt.Print("matriks [", i, "][", j, "]", "\n")
// 		}
// 	}

// }
