package main

import (
	"fmt"
)

func main() {
	greeting := greet("fr")
	fmt.Println(greeting)
}

type language string

var pharsebook = map[language]string{
	"en":  "Hello world",
	"fr":  "Bonjour",
	"hin": "Kaise ho",
}

func greet(l language) string {
	if greeting, ok := pharsebook[l]; ok {
		return greeting

	}
	return fmt.Sprintf("unsupported language %q.", l)
}

// func greet(l language) string{

// 	switch l{
// 	case "en":
// 		return "Hello world"
// 	case "fr":
// 		return "Bonjour"
// 	default:
// 		return ""
// 	}

// 	// return "Hello world"
// }
