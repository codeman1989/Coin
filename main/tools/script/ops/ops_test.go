package ops

import (
	"fmt"
	"testing"
)

func Test_provider(t *testing.T) {
	code, _ := defaultProvider.GetCode("OP_FALSE")
	fmt.Printf("%x\n", code)
	code1, _ := defaultProvider.GetCode("OP_MAX")
	fmt.Printf("%x\n", code1)
	_, err := defaultProvider.GetCode("OPMAX")
	fmt.Printf("%t\n", err)

	_, bo := defaultProvider.GetOp(uint8(76))
	fmt.Printf("%t\n", bo)
}
