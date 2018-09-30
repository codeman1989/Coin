package ops

import (
	"bytes"
)

func opInvert(c Context) error {
	if c.Size() < 1 {
		return ErrInvalidStackOperation
	}

	in := duplicate(c.Pop())

	for i := range in {
		in[i] ^= 0xFF
	}

	c.Push(in)

	return nil
}

func opAnd(c Context) error {
	return applyOp(c, func(a, b byte) byte { return a & b })
}

func opOr(c Context) error {
	return applyOp(c, func(a, b byte) byte { return a | b })
}

func opXor(c Context) error {
	return applyOp(c, func(a, b byte) byte { return a ^ b })
}

func opEqual(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}

	return writeBool(c, bytes.Equal(c.Pop(), c.Pop()))

}

func opEqualVerify(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}

	e := bytes.Equal(c.Pop(), c.Pop())
	writeBool(c, e)

	if !e {
		return ErrEqualVerify
	}

	return nil
}

func applyOp(c Context, binaryOp func(a, b byte) byte) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}

	x2 := c.Pop()
	x1 := c.Pop()

	// Pick the smaller of the two
	x := x2
	if len(x1) < len(x) {
		x = x1
	}

	out := make([]byte, len(x))
	for i := range x {
		out[i] = binaryOp(x1[i], x2[i])
	}

	c.Push(out)

	return nil
}
