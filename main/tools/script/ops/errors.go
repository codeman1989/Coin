package ops

import (
	"errors"
)

var (
	// ErrInvalidStackOperation 无效操作
	ErrInvalidStackOperation = errors.New("operation not valid with the current stack size")

	// ErrVerify 表示OpVerify操作失败
	ErrVerify = errors.New("script failed an OP_VERIFY operation")

	// ErrEqualVerify 表示 OpEqualVerify 操作失败
	ErrEqualVerify = errors.New("script failed an OP_EQUALVERIFY operation")

	// ErrNumEqualVerify 表示 OpNumEqualVerify 失败
	ErrNumEqualVerify = errors.New("script failed an OP_NUMEQUALVERIFY operation")

	// ErrInsufficientNumberOfBytes 表示push n字节失败 由于reader没有读出n字节
	ErrInsufficientNumberOfBytes = errors.New("insufficient number of bytes available")
)
