package ops

import (
	"io"
)

// Context
type Context interface {
	io.Reader

	Pop() []byte
	PopAlt() []byte

	Push([]byte)
	PushAlt([]byte)

	Size() int
	SizeAlt() int
}

// Op是将被执行的指令，需要一个上下文参数
type Op func(Context) error
