package main

// import "fmt"

// type person struct {
// 	name string
// 	age  int
// }

// type student struct {
// 	person
// 	age   int
// 	grade int
// }

// // Nested struct
// type siswa struct {
// 	person struct {
// 		name string
// 		age  int
// 	}
// 	grade   int
// 	hobbies []string
// }

// // Tag property pada struct
// type orang struct {
// 	name string `tag1`
// 	age  int    `tag2`
// }

// func main() {
// 	// var s1 = student{}
// 	// s1.name = "wick"
// 	// s1.grade = 9
// 	// s1.age = 21        // age of student
// 	// s1.person.age = 22 // age of person

// 	// fmt.Println(s1.name)
// 	// fmt.Println(s1.age)
// 	// fmt.Println(s1.person.age)

// 	// var s1 = student{
// 	// 	person: person{name: "Khilmi aminudin", age: 23},
// 	// 	age:    21,
// 	// 	grade:  6,
// 	// }

// 	// fmt.Println(s1)

// 	// var s1 = struct {
// 	// 	person
// 	// 	age int
// 	// }{
// 	// 	person: person{name: "khilmi aminudin", age: 23},
// 	// 	age:    23,
// 	// }

// 	// fmt.Println(s1)

// 	// // Slice dan struct
// 	// var allStudents = []person{
// 	// 	{name: "Wick", age: 23},
// 	// 	{name: "Ethan", age: 23},
// 	// 	{name: "Bourne", age: 22},
// 	// }

// 	// for _, student := range allStudents {
// 	// 	fmt.Println(student.name, "age is", student.age)
// 	// }

// 	// var allStudents1 = []struct {
// 	// 	person
// 	// 	grade int
// 	// }{
// 	// 	{person: person{"wick", 21}, grade: 2},
// 	// 	{person: person{"ethan", 22}, grade: 3},
// 	// 	{person: person{"bond", 21}, grade: 3},
// 	// }

// 	// for _, student := range allStudents1 {
// 	// 	fmt.Println(student)
// 	// }

// 	// // struct menggunakan var
// 	// // hanya deklarasi
// 	// var student struct {
// 	// 	grade int
// 	// }

// 	// // deklarasi sekaligus inisialisasi
// 	// var student1 = struct {
// 	// 	grade int
// 	// } {
// 	// 	12,
// 	// }

// 	// Type Alias
// 	type Person struct {
// 		name string
// 		age  int
// 	}
// 	type People = Person

// 	var p1 = Person{"wick", 21}
// 	fmt.Println(p1)

// 	var p2 = People{"wick", 21}
// 	fmt.Println(p2)

// }
