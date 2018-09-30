package ops

import (
	"bytes"
	"encoding/binary"
)

// duplicate 创建给定字节切片的副本
func duplicate(v []byte) []byte {
	c := make([]byte, len(v))
	copy(c, v)

	return c
}

func writeInt(c Context, num int32) error {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, &num); err != nil {
		return err
	}

	c.Push(buf.Bytes())

	return nil
}

func readInt(c Context) (d int32, err error) {
	if c.Size() < 1 {
		return d, ErrInvalidStackOperation
	}

	buf := bytes.NewBuffer(c.Pop())
	err = binary.Read(buf, binary.LittleEndian, &d)

	return d, err
}

func readBool(c Context) (bool, error) {
	if c.Size() < 1 {
		return false, ErrInvalidStackOperation
	}

	for _, b := range c.Pop() {
		if b != 0x00 {
			return true, nil
		}
	}

	return false, nil
}

func writeBool(c Context, b bool) error {
	if b {
		return writeInt(c, 1)
	}

	return writeInt(c, 0)
}
