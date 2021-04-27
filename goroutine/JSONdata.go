package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type User struct {
// 	FullName string `json:"Name"`
// 	Age      int
// }

// // Decode JSON
// // json.Unmarshal()
// // merubah data json string ke object di Golang

// // Decode JSON
// // json.Marshal()
// // merubah object di Golang ke data json string
// func main() {
// 	var jsonString = `{"Name": "john wick", "Age": 27}`
// 	var jsonData = []byte(jsonString)

// 	// JSON dengan STRUCT
// 	var data User
// 	var err = json.Unmarshal(jsonData, &data)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Println("user :", data.FullName)
// 	fmt.Println("age  :", data.Age)
// 	fmt.Println(data)

// 	// JSON dengan MAP
// 	var data1 map[string]interface{}
// 	json.Unmarshal(jsonData, &data1)

// 	fmt.Println("user :", data1["Name"])
// 	fmt.Println("age  :", data1["Age"])
// 	fmt.Println(data1)

// 	var jsonString1 = `[
// 		{"Name": "john wick", "Age": 27},
// 		{"Name": "ethan hunt", "Age": 32}
// 	]`

// 	var data2 []User

// 	var errr = json.Unmarshal([]byte(jsonString1), &data2)
// 	if errr != nil {
// 		fmt.Println(errr.Error())
// 		return
// 	}

// 	fmt.Println("user 1:", data2[0].FullName)
// 	fmt.Println("user 2:", data2[1].FullName)

// 	// json.Marshall()
// 	// []User
// 	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
// 	var jsonData0, err0 = json.Marshal(object)
// 	if err0 != nil {
// 		fmt.Println(err0.Error())
// 		return
// 	}

// 	var jsonString0 = string(jsonData0)
// 	fmt.Println(jsonString0)

// }
