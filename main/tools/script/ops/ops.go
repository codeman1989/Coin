package ops

// 创建默认provider结构
var (
	defaultProvider = &provider{
		map[uint8]Op{
			// 常量
			OpFalse:     opFalse,
			OpPushData1: createOpPushNBytes(1),
			OpPushData2: createOpPushNBytes(2),
			OpPushData4: createOpPushNBytes(4),
			Op1Negate:   op1Negate,
			OpTrue:      opTrue,

			// 逻辑
			OpEqual:       opEqual,
			OpInvert:      opInvert,
			OpAnd:         opAnd,
			OpOr:          opOr,
			OpXor:         opXor,
			OpEqualVerify: opEqualVerify,

			// 栈
			OpToAltStack:   opToAltStack,
			OpFromAltStack: opFromAltStack,
			OpIfDup:        opIfDup,
			OpDepth:        opDepth,
			OpDup:          opDup,
			OpNip:          opNip,
			OpOver:         opOver,
			OpPick:         opPick,
			OpRoll:         opRoll,
			OpRot:          opRot,
			OpSwap:         opSwap,
			OpTuck:         opTuck,
			Op2Drop:        op2Drop,
			Op2Dup:         op2Dup,
			Op3Dup:         op3Dup,
			Op2Over:        op2Over,
			Op2Rot:         op2Rot,
			Op2Swap:        op2Swap,

			// 拼接
			OpCat:    opCat,
			OpSubstr: opSubstr,
			OpLeft:   opLeft,
			OpRight:  opRight,
			OpSize:   opSize,

			// 算术
			Op1Add:               op1Add,
			Op1Sub:               op1Sub,
			Op2Mul:               op2Mul,
			Op2Div:               op2Div,
			OpNegate:             opNegate,
			OpAbs:                opAbs,
			OpNot:                opNot,
			Op0NotEqual:          op0NotEqual,
			OpAdd:                opAdd,
			OpSub:                opSub,
			OpMul:                opMul,
			OpDiv:                opDiv,
			OpMod:                opMod,
			OpLShift:             opLShift,
			OpRShift:             opRShift,
			OpBoolAnd:            opBoolAnd,
			OpBoolOr:             opBoolOr,
			OpNumEqual:           opNumEqual,
			OpNumEqualVerify:     opNumEqualVerify,
			OpNumNotEqual:        opNumNotEqual,
			OpLessThan:           opLessThan,
			OpGreaterThan:        opGreaterThan,
			OpLessThanOrEqual:    opLessThanOrEqual,
			OpGreaterThanOrEqual: opGreaterThanOrEqual,
			OpMin:                opMin,
			OpMax:                opMax,
			OpWithin:             opWithin,
		},
		map[string]uint8{
			// Constants
			"OP_FALSE":     OpFalse,
			"OP_PUSHDATA1": OpPushData1,
			"OP_PUSHDATA2": OpPushData2,
			"OP_PUSHDATA4": OpPushData4,
			"OP_1NEGATE":   Op1Negate,
			"OP_TRUE":      OpTrue,

			// Logic
			"OP_EQUAL": OpEqual,

			// Stack
			"OP_TOALTSTACK":   OpToAltStack,
			"OP_FROMALTSTACK": OpFromAltStack,
			"OP_IFDUP":        OpIfDup,
			"OP_DEPTH":        OpDepth,
			"OP_DUP":          OpDup,
			"OP_NIP":          OpNip,
			"OP_OVER":         OpOver,
			"OP_PICK":         OpPick,
			"OP_ROLL":         OpRoll,
			"OP_ROT":          OpRot,
			"OP_SWAP":         OpSwap,
			"OP_TUCK":         OpTuck,
			"OP_2DROP":        Op2Drop,
			"OP_2DUP":         Op2Dup,
			"OP_3DUP":         Op3Dup,
			"OP_2OVER":        Op2Over,
			"OP_2ROT":         Op2Rot,
			"OP_2SWAP":        Op2Swap,

			// Splice

			"OP_CAT":    OpCat,
			"OP_SUBSTR": OpSubstr,
			"OP_LEFT":   OpLeft,
			"OP_RIGHT":  OpRight,
			"OP_SIZE":   OpSize,

			// Arithmetic
			"OP_1ADD":               Op1Add,
			"OP_1SUB":               Op1Sub,
			"OP_2MUL":               Op2Mul,
			"OP_2DIV":               Op2Div,
			"OP_NEGATE":             OpNegate,
			"OP_ABS":                OpAbs,
			"OP_NOT":                OpNot,
			"OP_0NOTEQUAL":          Op0NotEqual,
			"OP_ADD":                OpAdd,
			"OP_SUB":                OpSub,
			"OP_MUL":                OpMul,
			"OP_DIV":                OpDiv,
			"OP_MOD":                OpMod,
			"OP_LSHIFT":             OpLShift,
			"OP_RSHIFT":             OpRShift,
			"OP_BOOLAND":            OpBoolAnd,
			"OP_BOOLOR":             OpBoolOr,
			"OP_NUMEQUAL":           OpNumEqual,
			"OP_NUMEQUALVERIFY":     OpNumEqualVerify,
			"OP_NUMNOTEQUAL":        OpNumNotEqual,
			"OP_LESSTHAN":           OpLessThan,
			"OP_GREATERTHAN":        OpGreaterThan,
			"OP_LESSTHANOREQUAL":    OpLessThanOrEqual,
			"OP_GREATERTHANOREQUAL": OpGreaterThanOrEqual,
			"OP_MIN":                OpMin,
			"OP_MAX":                OpMax,
			"OP_WITHIN":             OpWithin,
		},
	}
	//
	Default Provider = defaultProvider
)

// 创建压入n个数字的函数
func createOpPushN(n uint8) Op {
	return func(c Context) error {
		return writeInt(c, int32(n))
	}
}

// 创建压入n个字符的函数
func createOpPushNBytes(n uint8) Op {
	return func(c Context) error {
		if n == 0 {
			return nil
		}
		bs := make([]byte, n)
		cnt, err := c.Read(bs)
		if err != nil {
			return err
		}

		if cnt != int(n) {
			return ErrInsufficientNumberOfBytes
		}

		c.Push(bs)

		return nil
	}

}

// ops 包的初始化函数
// OpPushData1 0x4c即76
func init() {
	// <=75都不是操作码，是要push的字符串长度
	for i := uint8(1); i <= 75; i++ {
		defaultProvider.ops[i] = createOpPushNBytes(i)
	}
	// 从2开始，因为Op1就是push
	for i := Op2; i <= Op16; i++ {
		defaultProvider.ops[i] = createOpPushN(i)
	}
}

// 以下为操作函数
func opFalse(c Context) error {
	return writeBool(c, false)
}

func op1Negate(c Context) error {
	return writeInt(c, -1)
}

func opTrue(c Context) error {
	return writeBool(c, true)
}
