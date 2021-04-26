package main

// import (
// 	"fmt"
// 	"os"
// )

// /*
// Defer digunakan untuk mengakhirkan eksekusi sebuah statement tepat sebelum blok fungsi selesai.
// Sedangkan Exit digunakan untuk menghentikan program secara paksa (ingat, menghentikan program, tidak seperti return yang hanya menghentikan blok kode).
// */

// func main() {
// 	// kode error tapi logging() tetap di eksekusi
// 	// hitungNilai(2,"/",0)

// 	// logging di eksekusi di akhir
// 	hitungNilai(7, "*", 5)

// 	// defer selalu di eksekusi di akhir block fungsi (main)
// 	number := 3

// 	if number == 3 {
// 		fmt.Println("halo 1")
// 		// IIFE membuat defer di eksekusi terlebih dahulu
// 		// result jadi halo 1,2,3
// 		// func ()  {
// 		// defer fmt.Println("halo 3")
// 		// }()

// 		// tanpa IIFE defer akan di eksekusi di akhir fungsi main
// 		defer fmt.Println("halo 3")

// 		// Exit menghentikan aplikasi
// 		// defer pun tidak akan tereksekusi jika ada os.Exit()
// 		os.Exit(1)

// 	}

// 	fmt.Println("halo 2")

// }

// func logging() {
// 	fmt.Println("defer keyword di gunakan")
// }

// func hitungNilai(val1 int8, operan string, val2 int8) {
// 	// pemanggilan defer
// 	defer logging()

// 	var resutl string
// 	var total interface{}

// 	switch operan {
// 	case "+":
// 		total = val1 + val2
// 	case "-":
// 		total = val1 - val2
// 	case "/":
// 		total = val1 / val2
// 	case "*":
// 		total = val1 * val2
// 	case "%":
// 		total = val1 % val2
// 	}

// 	switch v := total.(type) {
// 	case float32:
// 		total = float32(v)
// 	case int8:
// 		total = int8(v)
// 	}

// 	resutl = fmt.Sprint(val1, operan, val2, "=", total)

// 	fmt.Println(resutl)
// }
