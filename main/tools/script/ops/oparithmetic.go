package ops

func unaryOp(c Context, f func(int32) int32) error {
	d, err := readInt(c)
	if err != nil {
		return err
	}
	return writeInt(c, f(d))
}

func binaryOp(c Context, f func(int32, int32) int32) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}

	b, err := readInt(c)
	if err != nil {
		return err
	}

	a, err := readInt(c)
	if err != nil {
		return err
	}

	return writeInt(c, f(a, b))
}

func op1Add(c Context) error {
	return unaryOp(c, func(d int32) int32 { return d + 1 })
}

func op1Sub(c Context) error {
	return unaryOp(c, func(d int32) int32 { return d - 1 })
}

func op2Mul(c Context) error {
	return unaryOp(c, func(d int32) int32 { return d * 2 })
}

func op2Div(c Context) error {
	return unaryOp(c, func(d int32) int32 { return d / 2 })
}

func opNegate(c Context) error {
	return unaryOp(c, func(d int32) int32 { return d * -1 })
}

func opAbs(c Context) error {
	return unaryOp(c, func(d int32) int32 {
		if d < 0 {
			return d * -1
		}
		return d
	})
}

func opNot(c Context) error {
	return unaryOp(c, func(d int32) int32 {
		if d < 0 {
			return 1
		}
		return 0
	})
}

func op0NotEqual(c Context) error {
	return unaryOp(c, func(d int32) int32 {
		if d == 0 {
			return 0
		}
		return 1
	})
}

func opAdd(c Context) error {
	return binaryOp(c, func(a, b int32) int32 { return a + b })
}

func opSub(c Context) error {
	return binaryOp(c, func(a, b int32) int32 { return a - b })
}

func opMul(c Context) error {
	return binaryOp(c, func(a, b int32) int32 { return a * b })
}

func opDiv(c Context) error {
	return binaryOp(c, func(a, b int32) int32 { return a / b })
}

func opMod(c Context) error {
	return binaryOp(c, func(a, b int32) int32 { return a % b })
}

func opLShift(c Context) error {
	return binaryOp(c, func(a, b int32) int32 {
		am := 0x8000 & a
		return am | (a << uint(b))
	})
}

func opRShift(c Context) error {
	return binaryOp(c, func(a, b int32) int32 {
		am := 0x8000 & a
		return am | (a >> uint(b))
	})
}

func opBoolAnd(c Context) error {
	return binaryOp(c, func(a, b int32) int32 {
		if a != 0 && b != 0 {
			return 1
		}
		return 0
	})
}

func opBoolOr(c Context) error {
	return binaryOp(c, func(a, b int32) int32 {
		if a != 0 || b != 0 {
			return 1
		}
		return 0
	})
}

func opNumEqual(c Context) error {
	return binaryOp(c, func(a, b int32) int32 {
		if a == b {
			return 1
		}
		return 0
	})
}

func opNumEqualVerify(c Context) error {
	err := binaryOp(c, func(a, b int32) int32 {
		if a == b {
			return 1
		}

		return 0
	})
	if err != nil {
		return err
	}

	b, err := readBool(c)
	if err != nil {
		return err
	}

	if !b {
		return ErrNumEqualVerify
	}

	return nil
}

func opNumNotEqual(c Context) error {
	return binaryOp(c, func(a, b int32) int32 {
		if a != b {
			return 1
		}
		return 0
	})
}

func opLessThan(c Context) error {
	return binaryOp(c, func(a, b int32) int32 {
		if a < b {
			return 1
		}
		return 0
	})
}

func opGreaterThan(c Context) error {
	return binaryOp(c, func(a, b int32) int32 {
		if a > b {
			return 1
		}
		return 0
	})
}

func opLessThanOrEqual(c Context) error {
	return binaryOp(c, func(a, b int32) int32 {
		if a <= b {
			return 1
		}
		return 0
	})
}

func opGreaterThanOrEqual(c Context) error {
	return binaryOp(c, func(a, b int32) int32 {
		if a >= b {
			return 1
		}
		return 0
	})
}

func opMin(c Context) error {
	return binaryOp(c, func(a, b int32) int32 {
		if a < b {
			return a
		}
		return b
	})
}

func opMax(c Context) error {
	return binaryOp(c, func(a, b int32) int32 {
		if a > b {
			return a
		}
		return b
	})
}

func opWithin(c Context) error {
	if c.Size() < 3 {
		return ErrInvalidStackOperation
	}

	max, err := readInt(c)
	if err != nil {
		return err
	}

	min, err := readInt(c)
	if err != nil {
		return err
	}

	v, err := readInt(c)
	if err != nil {
		return err
	}

	if v >= min && v < max {
		return writeInt(c, 1)
	}

	return writeInt(c, 0)
}
