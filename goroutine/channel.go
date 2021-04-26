package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	runtime.GOMAXPROCS(2)

// 	var messages = make(chan string)

// 	var sayHelloTo = func(who string) {
// 		var data = fmt.Sprintf("hello %s", who)
// 		messages <- data
// 	}

// 	go sayHelloTo("john wick")
// 	go sayHelloTo("ethan hunt")
// 	go sayHelloTo("jason bourne")

// 	var message1 = <-messages
// 	fmt.Println(message1)

// 	var message2 = <-messages
// 	fmt.Println(message2)

// 	var message3 = <-messages
// 	fmt.Println(message3)

// 	fmt.Println()

// 	// // channel sebagai parameter
// 	var messages1 = make(chan string)

// 	for _, each := range []string{"wick", "hunt", "bourne"} {
// 		go func(who string) {
// 			var data = fmt.Sprintf("hello %s", who)
// 			messages1 <- data
// 		}(each)
// 	}

// 	for i := 0; i < 3; i++ {
// 		printMessage(messages1)
// 	}
// }

// func printMessage(what chan string) {
// 	fmt.Println(<-what)
// }
