package base58

import (
	"bytes"
	"crypto/sha256"
	"math/big"

	"golang.org/x/crypto/ripemd160"
)

const (
	// alphabet is the modified base58 alphabet used by Bitcoin.
	alphabet     = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	alphabetIdx0 = '1'
)
const ADDRESSVERSION = byte(0x00) //地址版本号

var b58 = [256]byte{
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 0, 1, 2, 3, 4, 5, 6,
	7, 8, 255, 255, 255, 255, 255, 255,
	255, 9, 10, 11, 12, 13, 14, 15,
	16, 255, 17, 18, 19, 20, 21, 255,
	22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 255, 255, 255, 255, 255,
	255, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 43, 255, 44, 45, 46,
	47, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
}

//go:generate go run genalphabet.go

var bigRadix = big.NewInt(58)
var bigZero = big.NewInt(0)

// Decode decodes a modified base58 string to a byte slice.
func DecodeBase58(b string) []byte {
	answer := big.NewInt(0)
	j := big.NewInt(1)

	scratch := new(big.Int)
	for i := len(b) - 1; i >= 0; i-- {
		//字符，ascii码表的简版-->得到字符代表的值(0，1,2，..57)
		tmp := b58[b[i]]
		//出现不该出现的字符
		if tmp == 255 {
			return []byte("")
		}
		scratch.SetInt64(int64(tmp))

		//scratch = j*scratch
		scratch.Mul(j, scratch)

		answer.Add(answer, scratch)
		//每次进位都要乘上58
		j.Mul(j, bigRadix)
	}

	//得到大端的字节序
	tmpval := answer.Bytes()

	var numZeros int
	for numZeros = 0; numZeros < len(b); numZeros++ {
		//得到高位0的位数
		if b[numZeros] != alphabetIdx0 {
			break
		}
	}
	//得到原来数字的长度
	flen := numZeros + len(tmpval)
	//构造一个新地存放结果的空间
	val := make([]byte, flen, flen)
	copy(val[numZeros:], tmpval)

	return val
}

// Encode encodes a byte slice to a modified base58 string.
func EncodeBase58(b []byte) string {
	x := new(big.Int)
	//将b解释为大端存储
	x.SetBytes(b)

	//Base58编码可以表示的比特位数为Log258 {\displaystyle \approx } \approx5.858bit。经过Base58编码的数据为原始的数据长度的1.37倍
	answer := make([]byte, 0, len(b)*136/100)

	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		//x除于58的余数mod，并将商赋值给x
		x.DivMod(x, bigRadix, mod)
		answer = append(answer, alphabet[mod.Int64()])
	}

	// leading zero bytes
	//因为如果高位为0，0除任何数为0，可以直接设置为‘1’
	for _, i := range b {
		if i != 0 {
			break
		}
		answer = append(answer, alphabetIdx0)
	}

	// reverse
	//因为之前先附加低位的，后附加高位的，所以需要翻转
	alen := len(answer)
	for i := 0; i < alen/2; i++ {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}

	return string(answer)
}

func EncodeBase58Check(vchIn []byte) string {
	h := sha256.New()
	h.Write(vchIn[:])
	b := h.Sum(nil)
	loc := append(vchIn, b[0], b[1], b[2], b[3])
	return EncodeBase58(loc[:])
}

func Hash160ToAddress(hash160 []byte) string {
	apx_v := append(hash160, ADDRESSVERSION)

	return EncodeBase58Check(apx_v[:])
}

func DecodeBase58Check(psz string) bool {
	ret := DecodeBase58(psz)
	if ret == nil {
		return false
	}
	if len(ret) != 25 {

		return false
	}

	if ret[len(ret)-5] > ADDRESSVERSION {
		return false
	}
	loc := ret[0 : len(ret)-4]
	h := sha256.New()
	h.Write(loc[:])
	hash := h.Sum(nil)

	flag := false

	if bytes.Compare(hash[0:3], ret[len(ret)-4:len(ret)-1]) == 0 {
		flag = true
	}

	return flag
}
func IsValidBitcoinAddress(psz string) bool {
	return DecodeBase58Check(psz)
}

// 公钥转地址
// base58encode(hash（公钥）+ 去前四字节（hash(hash（公钥）+ 版本)))
func PubKeyToAddress(PubKey []byte) string {
	rim := ripemd160.New()
	rim.Write(PubKey[:])

	return Hash160ToAddress(rim.Sum(nil))
}
