package main

// import (
// 	"flag"
// 	"fmt"
// )

// // Arguments
// // func main() {
// // 	var argsRaw = os.Args
// // 	fmt.Printf("-> %#v\n", argsRaw)
// // 	// -> []string{".../bab45", "banana", "potato", "ice cream"}

// // 	var args = argsRaw[1:]
// // 	fmt.Printf("-> %#v\n", args)
// // 	// -> []string{"banana", "potatao", "ice cream"}
// // }

// // flag
// func main() {

// 	// var name = flag.String("name", "anonymous", "type your name")
// 	// var age = flag.Int64("age", 25, "type your age")

// 	// flag.Parse()
// 	// fmt.Printf("name\t: %s\n", *name)
// 	// fmt.Printf("age\t: %d\n", *age)

// 	// cara ke-1
// 	var data1 = flag.String("name", "anonymous", "type your name")
// 	fmt.Println(*data1)

// 	// cara ke-2
// 	var data2 string
// 	flag.StringVar(&data2, "gender", "male", "type your gender")
// 	fmt.Println(data2)
// }

// /*
// +-------------------------------------------+-----------------------+
// |		Nama Fungsi							|		Return Value	|
// +-------------------------------------------+-----------------------+
// | flag.Bool(name, defaultValue, usage)		|	*bool				|
// | flag.Duration(name, defaultValue, usage)	|	*time.Duration		|
// | flag.Float64(name, defaultValue, usage)	|	*float64			|
// | flag.Int(name, defaultValue, usage)		|	*int				|
// | flag.Int64(name, defaultValue, usage)		|	*int64				|
// | flag.String(name, defaultValue, usage)	|	*string				|
// | flag.Uint(name, defaultValue, usage)		|	*uint				|
// | flag.Uint64(name, defaultValue, usage)	|	*uint64				|
// +-------------------------------------------+-----------------------+
// */
