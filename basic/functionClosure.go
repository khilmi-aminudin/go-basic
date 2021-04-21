package main

// import "fmt"

// func main() {
// 	var getMinMax = func(n []int) (int, int) {
// 		var min, max int
// 		for i, e := range n {
// 			switch {
// 			case i == 0:
// 				max, min = e, e
// 			case e > max:
// 				max = e
// 			case e < min:
// 				min = e
// 			}
// 		}
// 		return min, max
// 	}

// 	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
// 	var min, max = getMinMax(numbers)
// 	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)

// 	// Immediately-Invoked Function Expression (IIFE)
// 	var numbers1 = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

//     var newNumbers = func(min int) []int {
//         var r []int
//         for _, e := range numbers1 {
//             if e < min {
//                 continue
//             }
//             r = append(r, e)
//         }
//         return r
//     }(3)

//     fmt.Println("original number :", numbers1)
//     fmt.Println("filtered number :", newNumbers)
// }
