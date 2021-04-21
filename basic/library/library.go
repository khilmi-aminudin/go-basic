package library

// exported
func HelloWorld() string {
	return "Hello World"
}

// Unexported
func helloWorld() string {
	return "Hello World"
}


// Export Import package
// Exported : variabe/struct/method/fungsi yang di awali dengan Uppercase
// Unexported : variabe/struct/method/fungsi yang di awali dengan Lowercase
// 