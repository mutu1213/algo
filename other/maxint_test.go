package other

import (
	"fmt"
	"testing"
)

func Test_MaxInt(t *testing.T) {
	MinUint := uint(0)
	MaxUint := ^MinUint
	MaxInt := int(MaxUint >> 1)
	MinInt := ^MaxInt
	fmt.Printf("max uint:%v,min uint:%v,max int:%v,min int:%v", MaxUint, MinUint, MaxInt, MinInt)
}
