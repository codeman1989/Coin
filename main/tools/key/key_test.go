package Key

import (
	"fmt"
	"testing"
)

func Test_GenetareKey(t *testing.T) {
	key := MakeNewKey()
	fmt.Printf("%x\n", key.PrivateKey)
	fmt.Printf("%x\n", key.PublicKey)
}
func Test_Verify(t *testing.T) {
	key := MakeNewKey()
	payload := []byte("Golang")
	ret, _ := Sign(payload, key.PrivateKey)
	if Verify(payload, ret, key.PublicKey) == false {
		t.Errorf("Verify Failed")
	}
}
