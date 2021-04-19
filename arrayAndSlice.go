package main

// import "fmt"

// func main() {

// 	// // contoh array
// 	// var names [4]string
// 	// names[0] = "Khilmi"
// 	// names[1] = "Aminudin"
// 	// names[2] = "Fathi"
// 	// names[3] = "Fazza"
// 	// // names[4] = "hai" // error out of bound

// 	// for i := 0; i < len(names); i++{
// 	// 	fmt.Print(names[i]+", ")
// 	// }

// 	// // initialitation with value
// 	// var fruits [4]string= [4]string{"apple", "banana", "strawberry", "orange"}
// 	// for i := 0; i < len(fruits); i++{
// 	// 	fmt.Print(fruits[i]+", ")
// 	// }

// 	// fmt.Println("")

// 	// // vertical initialitation
// 	// var fruits1 [4]string

// 	// // cara horizontal
// 	// fruits1  = [4]string{"apple", "grape", "banana", "melon"}

// 	// // cara vertikal
// 	// fruits1  = [4]string{
// 	// 	"apple",
// 	// 	"grape",
// 	// 	"banana",
// 	// 	"melon",
// 	// }
// 	// var t = 0
// 	// for {
// 	// 	fmt.Println(fruits1[t], "nilai t : ",t)
// 	// 	t++
// 	// 	if t == len(fruits1){
// 	// 		break
// 	// 	}
// 	// }
// 	// fmt.Println(len(fruits1))

// 	// // Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
// 	// var numbers = [...]int{1,2,3,4,5,6,7,8,9}
// 	// fmt.Println("data array \t:", numbers)
// 	// fmt.Println("array length\t:", len(numbers))

// 	// // Array multidimensi
// 	// var numbers1 = [2][3]int{
// 	// 	[3]int{3, 2, 3},
// 	// 	[3]int{3, 4, 5},
// 	// }
// 	// var numbers2 = [2][3]int{
// 	// 	{3, 2, 3},
// 	// 	{3, 4, 5},
// 	// }

// 	// fmt.Println("numbers1", numbers1)
// 	// fmt.Println("numbers2", numbers2)

// 	// // for loop untuk menampilkan nilai array
// 	// var fruits3 = [4]string{"apple", "grape", "banana", "melon"}

// 	// for i := 0; i < len(fruits3); i++ {
// 	// 	fmt.Printf("elemen %d : %s\n", i, fruits[i])
// 	// }

// 	// // for-range
// 	// /*
// 	// 	for indexArray, valueArray := range namaArray {
// 	// 			block code //
// 	// 	}
// 	//  */
// 	// for i, fruit := range fruits3 {
// 	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
// 	// }

// 	// // underscore variable untuk menampung index array
// 	// // jika index array nya tidak akan digunakan ataupun seba;iknya

// 	// // underscore untuk index
// 	// for _, fruit := range fruits {
// 	// 	fmt.Printf("nama buah : %s\n", fruit)
// 	// }
// 	// // underscore untuk value
// 	// for i, _ := range fruits {
// 	// 	fmt.Println(i)
// 	// }
// 	// // Deklarasi sekaligus alokasi data array juga bisa dilakukan lewat keyword make
// 	// var fruits4 = make([]string, 2)
// 	// fruits4[0] = "apple"
// 	// fruits4[1] = "manggo"

// 	// fmt.Println(fruits4)  // [apple manggo]

// 	// // Perbedaan array dan slice
// 	// var fruitsA = []string{"apple", "grape"}      // slice
// 	// var fruitsB = [2]string{"banana", "melon"}    // array
// 	// var fruitsC = [...]string{"papaya", "grape"}  // array

// 	// fmt.Println(fruitsA,"\n", fruitsB,"\n", fruitsC)

// 	// // membuat dari array
// 	// var fruitsD = []string{"apple", "grape", "banana", "melon"} // Array
// 	// var newFruits = fruitsD[0:2]								//Slice
// 	// // var newFruits = fruitsD[indexAwal:indexsesudahterakhir]
// 	// fmt.Println(newFruits) // ["apple", "grape"]

// 	// // len() pada slice
// 	// var fruits = []string{"apple", "grape", "banana", "melon"}
// 	// fmt.Println(len(fruits)) // 4

// 	// // cap() pada slice
// 	// var fruits = []string{"apple", "grape", "banana", "melon"}
// 	// fmt.Println(len(fruits)) // len: 4
// 	// fmt.Println(cap(fruits)) // cap: 4

// 	// var aFruits = fruits[0:3]
// 	// fmt.Println(len(aFruits)) // len: 3
// 	// fmt.Println(cap(aFruits)) // cap: 4

// 	// var bFruits = fruits[1:4]
// 	// fmt.Println(len(bFruits)) // len: 3
// 	// fmt.Println(cap(bFruits)) // cap: 3
// 	/*
// 		+--------------------------------------------------------+
// 		|		Kode	| 			Output		| len() |  cap() |
// 		+---------------+-----------------------+-------+--------|
// 		| fruits[0:4]	| [buah buah buah buah]	|	4	|	4	 |
// 		| aFruits[0:3]	| [buah buah buah ----]	|	3	|	4	 |
// 		| bFruits[1:3]	| [---- [buah buah buah]|	3	|	3	 |
// 		+---------------+-----------------------+-------+--------+
// 	*/

// 	// append() pada slice
// 	var fruitsA = []string{"apple", "grape", "banana"}
// 	var cFruitsA = append(fruitsA, "papaya")

// 	fmt.Println(fruitsA)  // ["apple", "grape", "banana"]
// 	fmt.Println(cFruitsA) // ["apple", "grape", "banana", "papaya"]

// 	// ------------------------------------------------------
// 	var fruits = []string{"apple", "grape", "banana"}
// 	var bFruits = fruits[0:2]

// 	fmt.Println(cap(bFruits)) // 3
// 	fmt.Println(len(bFruits)) // 2

// 	fmt.Println(fruits)  // ["apple", "grape", "banana"]
// 	fmt.Println(bFruits) // ["apple", "grape"]

// 	var cFruits = append(bFruits, "papaya")

// 	fmt.Println(fruits)  // ["apple", "grape", "papaya"]
// 	fmt.Println(bFruits) // ["apple", "grape"]
// 	fmt.Println(cFruits) // ["apple", "grape", "papaya"]

// 	// copy() pada slice
// 	dst := make([]string, 3)
// 	src := []string{"watermelon", "pinnaple", "apple", "orange"}
// 	n := copy(dst, src)

// 	fmt.Println(dst) // watermelon pinnaple apple
// 	fmt.Println(src) // watermelon pinnaple apple orange
// 	fmt.Println(n)   // 3

// }

// // slice s&k
// // slice adalah tipe data references sehingga jika data di slah satu field di rubah maka nilai nya akan ikut terubah
// // var fruits = []string{"apple", "grape", "banana", "melon"}

// /*
// 	+-----------------------------------------------+---------------------------------------------------------------------------------------+
// 	|	Kode		|			Output				|										Penjelasan										|
// 	+---------------+-------------------------------+---------------------------------------------------------------------------------------+
// 	|fruits[0:2]	| [apple, grape]				|	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-2							|
// 	|fruits[0:4]	| [apple, grape, banana, melon]	|	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-4							|
// 	|fruits[0:0]	| []							|	menghasilkan slice kosong, karena tidak ada elemen sebelum indeks ke-0				|
// 	|fruits[4:4]	| []							|	menghasilkan slice kosong, karena tidak ada elemen yang dimulai dari indeks ke-4	|
// 	|fruits[4:0]	| []							|	error, pada penulisan fruits[a,b] nilai a harus lebih besar atau sama dengan b		|
// 	|fruits[:]		| [apple, grape, banana, melon]	|	semua elemen																		|
// 	|fruits[2:]		| [banana, melon]				|	semua elemen mulai indeks ke-2														|
// 	|fruits[:2]		| [apple, grape]				|	semua elemen hingga sebelum indeks ke-2												|
// 	+---------------+-------------------------------+---------------------------------------------------------------------------------------+
// */
