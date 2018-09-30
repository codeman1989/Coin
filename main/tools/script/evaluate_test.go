package script

import (
	"fmt"
	"testing"

	"./ops"
)

func Test_evaluate(t *testing.T) {
	num, _ := ops.Default.GetCode("OP_PUSHDATA1")
	if num != 76 {
		t.Errorf("GetCode Failed")
	}

	buf, _ := Parse("33 33 OP_EQUAL")
	err := Evaluate(buf)
	fmt.Printf("%s\n", err)

}
