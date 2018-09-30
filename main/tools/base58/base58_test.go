package base58

import (
	"testing"
)

func TestDecode(t *testing.T) {
	if EncodeBase58(DecodeBase58("1PrMRMbBhGhDKw1C2Zm92nY4uwtQViaeyd")) != "1PrMRMbBhGhDKw1C2Zm92nY4uwtQViaeyd" {
		t.Error("Base58 Failed")
	}

}

func TestP2H(t *testing.T) {
	a := []byte{94, 82, 254, 228, 126, 107, 7, 5, 101, 247, 67, 114, 70, 140, 220, 105, 157, 232, 145, 7, 0, 65, 32, 180, 81}
	if PubKeyToAddress([]byte("test")) != EncodeBase58(a) {
		t.Error("P2H Failed")
	}
}

func TestIsValidBitcoinAddress(t *testing.T) {
	if IsValidBitcoinAddress("exLYKBakCrYAe6qKkJbCxVXZPnDxMATzo2") != true {
		t.Error("IsValidBitcoinAddress Failed")
	}
}
