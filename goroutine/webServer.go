package main

// import (
// 	"fmt"
// 	"net/http"
// 	"text/template"
// )

// func index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "apa kabar?")
// }

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		var data = map[string]string{
// 			"Name":    "john wick",
// 			"Message": "have a nice day",
// 		}

// 		var t, err = template.ParseFiles("template.html")
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return
// 		}

// 		t.Execute(w, data)
// 	})

// 	http.HandleFunc("/index", index)

// 	fmt.Println("starting web server at http://localhost:9000/")
// 	http.ListenAndServe(":9000", nil)
// }
