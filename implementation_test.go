package lab2

import (
	"fmt"
	"log"
	"testing"
)

func TestPrefixToInfix(t *testing.T) {
	data := map[string]string{
		"+ 5 * + 4 2 ^ 3 2":             "5 + (4 + 2) * (3 ^ 2)",
		"+ / * + 3 5 - 2 8 4 * 6 + 1 2": "(3 + 5) * (2 - 8) / 4 + 6 * (1 + 2)",
	}
	for input, expected := range data {
		t.Run(input, func(t *testing.T) {
			res, err := PrefixToInfix(input)

			if err != nil {
				t.Error(err)
				return
			}
			if res != expected {
				t.Errorf("Expected %s, but got %s", expected, res)
			}
		})
	}
}

func TestPrefixToInfixWithError(t *testing.T) {
	inputs := []string{"", "foo", "foo bar foo", "+", "5 + 5", "+ 5", "+ 5 foo"}
	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			res, err := PrefixToInfix(input)

			if err == nil {
				t.Errorf("Expected error, but got %s", res)
			}
		})

	}

}

func ExamplePrefixToInfix() {
	input := "+ 3 * 5 4"
	result, err := PrefixToInfix(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result) // Outputs: 3 + 5 * 4
}
