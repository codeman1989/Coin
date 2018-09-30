package ops

import (
	"bytes"
)

func opCat(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}
	x2 := c.Pop()
	x1 := c.Pop()

	var buf bytes.Buffer
	_, err := buf.Write(x1)
	if err != nil {
		return err
	}
	_, err = buf.Write(x2)
	if err != nil {
		return err
	}

	c.Push(buf.Bytes())

	return nil
}

func opSubstr(c Context) error {
	if c.Size() < 3 {
		return ErrInvalidStackOperation
	}

	size, err := readInt(c)
	if err != nil {
		return err
	}

	begin, err := readInt(c)
	if err != nil {
		return err
	}

	in := string(duplicate(c.Pop()))
	c.Push([]byte(in[begin : begin+size]))

	return nil
}

func opLeft(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}
	size, err := readInt(c)
	if err != nil {
		return err
	}
	in := string(duplicate(c.Pop()))

	if size >= int32(len(in)) {
		return ErrInvalidStackOperation
	}

	c.Push([]byte(in[0:size]))
	return nil
}

func opRight(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}
	size, err := readInt(c)
	if err != nil {
		return err
	}

	in := string(duplicate(c.Pop()))

	if size >= int32(len(in)) {
		return ErrInvalidStackOperation
	}

	c.Push([]byte(in[size:]))
	return nil
}

func opSize(c Context) error {
	if c.Size() < 1 {
		return ErrInvalidStackOperation
	}

	in := c.Pop()
	c.Push(in)

	return writeInt(c, int32(len(string(in))))
}
