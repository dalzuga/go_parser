package tools_test

import "fmt"

func ExampleSubStringInString() {
	fmt.Println(SubStringInString("ab", "abcdef"))
	fmt.Println(SubStringInString("substring", "string"))
}
