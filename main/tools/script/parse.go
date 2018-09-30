package script

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	"./ops"
)

// Parse 生成字节集，用于执行（evaluated）
func Parse(in string) (*bytes.Buffer, error) {
	// 定义byte slice
	buf := []byte{}
	// 用空格分割字符串，如 ‘123a 222a’分为 {123a,222a}
	for _, token := range strings.Fields(in) {
		op, status := ops.Default.GetCode(token)
		if status {
			buf = append(buf, byte(op))
		} else {
			str, err := hex.DecodeString(token)
			if err != nil {
				return new(bytes.Buffer), err
			}
			// 75以上是操作符,字符串长度不能超过75
			if len(str) > 75 {
				return new(bytes.Buffer), fmt.Errorf("data token too large")
			}
			buf = append(buf, byte(uint(len(str))))
			buf = append(buf, str...)
		}
	}
	return bytes.NewBuffer(buf), nil
}
