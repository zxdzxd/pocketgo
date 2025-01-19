package main

import (
	"testing"
)

func Examplemain(){
	main()
	// output:
	// Hello world
}

func TestGreet(t *testing.T){
	type testCase struct{
		lang language
		expectedGreeting string
	}

	var tests = map[string]testCase{
		"English": {
			lang: "en",
			expectedGreeting: "Hello world",
		},
		"French": {
			lang: "fr",
			expectedGreeting: "Bonjour",
		},
		"Hindi": {
			lang: "hin",
			expectedGreeting: "Kaise ho",
		},
	}

	for languageName, tc := range tests{
		t.Run(languageName, func(t *testing.T){
			greeting := greet(tc.lang)
			if greeting != tc.expectedGreeting{
				t.Errorf("Expected: %q, Got: %q",tc.expectedGreeting, greeting)
			}
		})
	}
}

// func TestGreet_English(t *testing.T){
// 	lang := "en"
// 	expectedGreeting := "Hello world"

// 	greeting := greet(language(lang))

// 	if greeting != expectedGreeting{
// 		t.Errorf("Expected: %q, Got: %q",expectedGreeting, greeting)
// 	}
// }

// func TestGreet_French(t *testing.T){
// 	lang := "fr"
// 	expectedGreeting := "Bonjour"

// 	greeting := greet(language(lang))

// 	if greeting != expectedGreeting{
// 		t.Errorf("Expected: %q, Got: %q",expectedGreeting, greeting)
// 	}
// }