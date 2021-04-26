package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	runtime.GOMAXPROCS(2)

// 	var messages = make(chan string)
// 	go sendMessage(messages)
// 	printMessage(messages)

// }

// func sendMessage(ch chan<- string) {
// 	for i := 0; i < 20; i++ {
// 		ch <- fmt.Sprintf("data %d", i)
// 	}
// 	close(ch)
// }

// func printMessage(ch <-chan string) {
// 	for message := range ch {
// 		fmt.Println(message)
// 	}
// }

// /*
// 	+-------------------------------------------------------------------------------------------+
// 	|			Sintaks		|							Penjelasan								|
// 	+-----------------------+-------------------------------------------------------------------+
// 	|	ch chan string		|	Parameter ch bisa digunakan untuk mengirim dan menerima data	|
// 	|	ch chan<- string	|	Parameter ch hanya bisa digunakan untuk mengirim data			|
// 	|	ch <-chan string	|	Parameter ch hanya bisa digunakan untuk menerima data			|
// 	+-------------------------------------------------------------------------------------------+
// */
