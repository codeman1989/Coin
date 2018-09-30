package Key

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"math/big"
)

const (
	SignRLen     = 28
	SignSLen     = 28
	PublicKeyLen = 230
)

type CKey struct {
	PublicKey  []byte
	PrivateKey []byte
}

func MakeNewKey() *CKey {
	key := &CKey{}
	pk, err := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	if err != nil {
		return nil
	}
	x509encoded, _ := x509.MarshalECPrivateKey(pk)
	key.PrivateKey = pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: x509encoded})

	x509encodedpub, _ := x509.MarshalPKIXPublicKey(&pk.PublicKey)
	key.PublicKey = pem.EncodeToMemory(&pem.Block{Type: "EC PUBLIC KEY", Bytes: x509encodedpub})

	return key
}

//Sign 通过私钥对数据进行签名
func Sign(payload []byte, privateKey []byte) ([]byte, error) {
	pk := privateKey
	block, _ := pem.Decode(pk)
	x509encoded := block.Bytes
	realPrivateKey, err := x509.ParseECPrivateKey(x509encoded)
	if err != nil {
		return nil, err
	}
	r, s, err := ecdsa.Sign(rand.Reader, realPrivateKey, payload)
	rBytes := FillBytesToFront(r.Bytes(), SignRLen)
	sBytes := FillBytesToFront(s.Bytes(), SignSLen)
	signature := append(rBytes, sBytes...)

	return signature, err
}

//Verify 验证签名对否
func Verify(payload []byte, signature []byte, publicKey []byte) bool {
	pk := publicKey

	blockPub, _ := pem.Decode(pk)
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	genericPublicKeyA := genericPublicKey.(*ecdsa.PublicKey)

	sign := signature
	buf := bytes.NewBuffer(sign)
	rBytes := new(big.Int).SetBytes(buf.Next(SignRLen))
	sBytes := new(big.Int).SetBytes(buf.Next(SignSLen))
	verifystatus := ecdsa.Verify(genericPublicKeyA, payload, rBytes, sBytes)
	return verifystatus
}

//FillBytesToFront 把数据截取/填充到指定长度
func FillBytesToFront(data []byte, totalLen int) []byte {
	if len(data) < totalLen {
		delta := totalLen - len(data)
		appendByte := []byte{}
		for delta != 0 {
			appendByte = append(appendByte, 0)
			delta--
		}
		return append(appendByte, data...)
	}
	return data[:totalLen]
}
