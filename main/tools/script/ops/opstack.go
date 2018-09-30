package ops

import (
	"bytes"
)

func opToAltStack(c Context) error {
	if c.Size() < 1 {
		return ErrInvalidStackOperation
	}

	c.PushAlt(c.Pop())

	return nil
}

func opFromAltStack(c Context) error {
	if c.SizeAlt() < 1 {
		return ErrInvalidStackOperation
	}

	c.Push(c.PopAlt())
	return nil
}

func opIfDup(c Context) error {
	if c.Size() < 1 {
		return ErrInvalidStackOperation
	}

	v := c.Pop()

	if !bytes.Equal(v, []byte{0x00, 0x00, 0x00, 0x00}) {
		v2 := make([]byte, len(v))
		copy(v2, v)
		c.Push(v2)
	}

	c.Push(v)

	return nil
}

func opDepth(c Context) error {
	return writeInt(c, int32(c.Size()))
}

func opDrop(c Context) error {
	if c.Size() < 1 {
		return ErrInvalidStackOperation
	}
	c.Pop()
	return nil
}

func opDup(c Context) error {
	if c.Size() < 1 {
		return ErrInvalidStackOperation
	}

	v := c.Pop()
	c.Push(duplicate(v))
	c.Push(v)

	return nil
}

func opNip(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}

	v := c.Pop()
	c.Pop()
	c.Push(v)
	return nil
}

func opOver(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}

	v2 := c.Pop()
	v1 := c.Pop()
	c.Push(v1)
	c.Push(v2)

	if v1 != nil {
		c.Push(duplicate(v1))
	}

	return nil
}

func opPick(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}

	n, err := readInt(c)
	if err != nil {
		return err
	}

	if n < 0 || n >= int32(c.Size()) {
		return ErrInvalidStackOperation
	}

	tmp := make([][]byte, n+1, n+1)

	for i := range tmp {
		tmp[i] = c.Pop()
	}

	for i := len(tmp) - 1; i >= 0; i-- {
		c.Push(tmp[i])
	}

	c.Push(duplicate(tmp[len(tmp)-1]))

	return nil
}

func opRoll(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}

	n, err := readInt(c)
	if err != nil {
		return err
	}

	if n < 0 || n >= int32(c.Size()) {
		return ErrInvalidStackOperation
	}

	tmp := make([][]byte, n+1, n+1)

	for i := range tmp {
		tmp[i] = c.Pop()
	}

	for i := len(tmp) - 2; i >= 0; i-- {
		c.Push(tmp[i])
	}

	c.Push(duplicate(tmp[len(tmp)-1]))

	return nil
}

func opRot(c Context) error {
	if c.Size() < 3 {
		return ErrInvalidStackOperation
	}

	v1 := c.Pop()
	v2 := c.Pop()
	v3 := c.Pop()

	c.Push(v2)
	c.Push(v3)
	c.Push(v1)

	return nil
}

func opSwap(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}

	v1 := c.Pop()
	v2 := c.Pop()

	c.Push(v1)
	c.Push(v2)

	return nil
}

func opTuck(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}

	v1 := c.Pop()
	v2 := c.Pop()

	c.Push(duplicate(v1))
	c.Push(v2)
	c.Push(v1)

	return nil
}

func op2Drop(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}

	c.Pop()
	c.Pop()

	return nil
}

func op2Dup(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}

	v1 := c.Pop()
	v2 := c.Pop()
	c.Push(v2)
	c.Push(v1)

	c.Push(duplicate(v2))
	c.Push(duplicate(v1))

	return nil
}

func op3Dup(c Context) error {
	if c.Size() < 3 {
		return ErrInvalidStackOperation
	}

	v1 := c.Pop()
	v2 := c.Pop()
	v3 := c.Pop()
	c.Push(v3)
	c.Push(v2)
	c.Push(v1)

	c.Push(duplicate(v3))
	c.Push(duplicate(v2))
	c.Push(duplicate(v1))

	return nil
}

func op2Over(c Context) error {
	if c.Size() < 4 {
		return ErrInvalidStackOperation
	}
	v1 := c.Pop()
	v2 := c.Pop()
	v3 := c.Pop()
	v4 := c.Pop()

	c.Push(v4)
	c.Push(v3)
	c.Push(v2)
	c.Push(v1)

	c.Push(duplicate(v4))
	c.Push(duplicate(v3))

	return nil
}

func op2Rot(c Context) error {
	if c.Size() < 6 {
		return ErrInvalidStackOperation
	}
	v1 := c.Pop()
	v2 := c.Pop()
	v3 := c.Pop()
	v4 := c.Pop()
	v5 := c.Pop()
	v6 := c.Pop()

	c.Push(v2)
	c.Push(v1)

	c.Push(v6)
	c.Push(v5)
	c.Push(v4)
	c.Push(v3)

	return nil
}

func op2Swap(c Context) error {
	if c.Size() < 4 {
		return ErrInvalidStackOperation
	}
	v1 := c.Pop()
	v2 := c.Pop()
	v3 := c.Pop()
	v4 := c.Pop()

	c.Push(v3)
	c.Push(v4)
	c.Push(v1)
	c.Push(v2)

	return nil
}
