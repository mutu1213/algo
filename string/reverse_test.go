package string

import (
	"fmt"
	"testing"
)

func Test_Reverse(t *testing.T) {
	fmt.Println(Reverse("123abc"))
}

func Reverse(s string) string {
	result := []byte(s)
	if len(result) != 0 {
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[j], result[i] = result[i], result[j]
		}
	}
	return string(result)
}
