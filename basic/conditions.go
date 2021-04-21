package main

// import "fmt"

// func main() {
// 	var angka = 2

// 	// regular if else condition
// 	if angka > 3 {
// 		fmt.Println("percabangan 1")
// 	} else if angka > 2 {
// 		fmt.Println("percabangan 2")
// 	} else {
// 		fmt.Println("block else")
// 	}

// 	// if else dengan variabel temporary
// 	var point = 8840.0

// 	if percent := point / 100; percent >= 100 {
// 		fmt.Printf("%.1f%s perfect!\n", percent, "%")
// 	} else if percent >= 70 {
// 		fmt.Printf("%.1f%s good\n", percent, "%")
// 	} else {
// 		fmt.Printf("%.1f%s not bad\n", percent, "%")
// 	}

// 	// switch case

// 	// Regular
// 	var point1 = 6

// 	switch point1 {
// 	case 8:
// 		fmt.Println("perfect")
// 	case 7:
// 		fmt.Println("awesome")
// 	default:
// 		fmt.Println("not bad")
// 	}

// 	// case codition
// 	var point2 = 6

// 	switch point2 {
// 	case 8:
// 		fmt.Println("perfect")
// 	case 7, 6, 5, 4:
// 		fmt.Println("awesome")
// 	default:
// 		fmt.Println("not bad")
// 	}

// 	// wiith block "{}" on case or default
// 	var point3 = 6

// 	switch point3 {
// 	case 8:
// 		fmt.Println("perfect")
// 	case 7, 6, 5, 4:
// 		fmt.Println("awesome")
// 	default:
// 		// block {}
// 		{
// 			fmt.Println("not bad")
// 			fmt.Println("you can be better!")
// 		}
// 	}

// 	// like if-else
// 	// kondisi tidak pada switch
// 	// tetapi langsung pada case
// 	var point4 = 6

// 	switch {
// 	case point4 == 8:
// 		fmt.Println("perfect")
// 	case (point4 < 8) && (point4 > 3):
// 		fmt.Println("awesome")
// 	default:
// 		{
// 			fmt.Println("not bad")
// 			fmt.Println("you need to learn more")
// 		}
// 	}

// 	// with fallthrough
// 	// melanjutkan ke case selanjutnya walau case sudah terpenuhi
// 	var point5 = 6

// 	switch {
// 	case point5 == 8:
// 		fmt.Println("perfect")
// 	case (point5 < 8) && (point5 > 3):
// 		fmt.Println("awesome")
// 		fallthrough
// 	case point5 < 5:
// 		fmt.Println("you need to learn more")
// 	default:
// 		{
// 			fmt.Println("not bad")
// 			fmt.Println("you need to learn more")
// 		}
// 	}

// 	// nested condition
// 	var pointx = 10

// 	if pointx > 7 {
// 		switch pointx {
// 		case 10:
// 			fmt.Println("perfect!")
// 		default:
// 			fmt.Println("nice!")
// 		}
// 	} else {
// 		if pointx == 5 {
// 			fmt.Println("not bad")
// 		} else if pointx == 3 {
// 			fmt.Println("keep trying")
// 		} else {
// 			fmt.Println("you can do it")
// 			if pointx == 0 {
// 				fmt.Println("try harder!")
// 			}
// 		}
// 	}
// }
