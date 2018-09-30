package script

import (
	"bytes"
	"fmt"
	"io"

	"./ops"
)

// Evaluate 执行字节集的所代表的脚本
func Evaluate(input io.Reader) error {
	// 1字节大小的slice，用于读取字节命令
	op := make([]byte, 1)
	// 上下文环境
	c := &context{&stack{}, &stack{}, input}

	// 循环读取命令
	for {
		_, err := input.Read(op)
		switch err {
		// 读取到了命令
		case nil:
			opCode := uint8(op[0])
			op, ok := ops.Default.GetOp(opCode)

			if !ok {
				return fmt.Errorf("unknown op for code %d", opCode)
			}
			// 执行操作
			err := op(c)
			if err != nil {
				return fmt.Errorf("op (%d) failed: %s", opCode, err)
			}
		// 读到了命令末尾
		case io.EOF:
			if bytes.Equal(c.stack.Pop(), []byte{0x00, 0x00, 0x00, 0x00}) {
				return fmt.Errorf("top value of stack is false")
			}
			return nil
		default:
			return err
		}
	}
}
