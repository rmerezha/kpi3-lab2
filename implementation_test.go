package lab2

import (
	"fmt"
	"testing"
)

func TestPrefixToPostfix(t *testing.T) {
	//res, err := PrefixToPostfix("+ 5 * - 4 2 3")
	//if assert.Nil(t, err) {
	//	assert.Equal(t, "4 2 - 3 * 5 +", res)
	//}
}

func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 2 +
}
