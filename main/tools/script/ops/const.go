package ops

// 缺少加密操作
// 如RIPEMD160 SHA1 SHA256 HASH256 CHECKSIG等
const (
	// 操作
	Op0         uint8 = 0x00
	OpFalse     uint8 = 0x00
	OpPushData1 uint8 = 0x4c
	OpPushData2 uint8 = 0x4d
	OpPushData4 uint8 = 0x4e
	Op1Negate   uint8 = 0x4f
	OpTrue      uint8 = 0x51
	Op1         uint8 = 0x51
	Op2         uint8 = 0x52
	Op3         uint8 = 0x53
	Op4         uint8 = 0x54
	Op5         uint8 = 0x55
	Op6         uint8 = 0x56
	Op7         uint8 = 0x57
	Op8         uint8 = 0x58
	Op9         uint8 = 0x59
	Op10        uint8 = 0x5a
	Op11        uint8 = 0x5b
	Op12        uint8 = 0x5c
	Op13        uint8 = 0x5d
	Op14        uint8 = 0x5e
	Op15        uint8 = 0x5f
	Op16        uint8 = 0x60

	// 算术
	Op1Add               uint8 = 0x8b
	Op1Sub               uint8 = 0x8c
	Op2Mul               uint8 = 0x8d
	Op2Div               uint8 = 0x8e
	OpNegate             uint8 = 0x8f
	OpAbs                uint8 = 0x90
	OpNot                uint8 = 0x91
	Op0NotEqual          uint8 = 0x92
	OpAdd                uint8 = 0x93
	OpSub                uint8 = 0x94
	OpMul                uint8 = 0x95
	OpDiv                uint8 = 0x96
	OpMod                uint8 = 0x97
	OpLShift             uint8 = 0x98
	OpRShift             uint8 = 0x99
	OpBoolAnd            uint8 = 0x9a
	OpBoolOr             uint8 = 0x9b
	OpNumEqual           uint8 = 0x9c
	OpNumEqualVerify     uint8 = 0x9d
	OpNumNotEqual        uint8 = 0x9e
	OpLessThan           uint8 = 0x9f
	OpGreaterThan        uint8 = 0xa0
	OpLessThanOrEqual    uint8 = 0xa1
	OpGreaterThanOrEqual uint8 = 0xa2
	OpMin                uint8 = 0xa3
	OpMax                uint8 = 0xa4
	OpWithin             uint8 = 0xa5

	// 位操作

	//  OpInvert 翻转输入的全部位
	// [in] => [out]
	OpInvert uint8 = 0x83

	//  OpAnd 把输入的两个位变成&操作
	// [x1, x2] => [x1 & x2]
	OpAnd uint8 = 0x84

	//  OpOr 或
	// [x1, x2] => [x1 | x2]
	OpOr uint8 = 0x85

	//  OpXor xor
	// [x1, x2] => [x1 ^ x2]
	OpXor uint8 = 0x86

	//  OpEqual 返回1，当输入相等时，否则返回0
	// [x1, x2] => [x1 == x2]
	OpEqual uint8 = 0x87

	// OpEqualVerify 返回1，如果输入相等。如果返回0，抛出ErrVerify错误
	OpEqualVerify uint8 = 0x88
	// OP_RESERVED1 OP_RESERVED2

	// 控制操作操作
	// 这里只有一个OpVerify
	// 缺少 NOP IF NOTIFY ELSE ENDIF RETURN等
	// OpVerify 如果栈顶部事务佳，返回无效
	OpVerify uint8 = 0x69
	// OP_NOP OP_VER OP_IF OP_NOTIF OP_BOTIF OP_VERIF OP_OP_VERNOTIF
	// OP_ELSE OP_ENDIF OP_RETURN

	// 拼接操作

	//  OpCat 拼接文本
	// [x1, x2] => [ x1x2 ]
	OpCat uint8 = 0x7e

	//  OpSubstr 返回中间文本
	// [in, begin, size] => [ in[begin:begin+size] ]
	OpSubstr uint8 = 0x7f

	//  OpLeft 返回文本左边定长文本
	// [in, size] => [ out ]
	OpLeft uint8 = 0x80

	//  OpRight 返回文本右边定长文本
	// [in, size] => [ out ]
	OpRight uint8 = 0x81

	//	OpSize 返回栈顶文本长度而不pop
	// [in] => [in, len(in)]
	OpSize uint8 = 0x82

	// 栈操作
	OpToAltStack   uint8 = 0x6b
	OpFromAltStack uint8 = 0x6c
	OpIfDup        uint8 = 0x73
	OpDepth        uint8 = 0x74
	OpDrop         uint8 = 0x75
	OpDup          uint8 = 0x76
	OpNip          uint8 = 0x77
	OpOver         uint8 = 0x78
	OpPick         uint8 = 0x79
	OpRoll         uint8 = 0x7a
	OpRot          uint8 = 0x7b
	OpSwap         uint8 = 0x7c
	OpTuck         uint8 = 0x7d
	Op2Drop        uint8 = 0x6d
	Op2Dup         uint8 = 0x6e
	Op3Dup         uint8 = 0x6f
	Op2Over        uint8 = 0x70
	Op2Rot         uint8 = 0x71
	Op2Swap        uint8 = 0x72

	// crypto
	//OP_RIPEMD160 = 0xa6,
	//OP_SHA1 = 0xa7,
	//OP_SHA256 = 0xa8,
	//OP_HASH160 = 0xa9,
	//OP_HASH256 = 0xaa,
	//OP_CODESEPARATOR = 0xab,
	//OP_CHECKSIG = 0xac,
	//OP_CHECKSIGVERIFY = 0xad,
	//OP_CHECKMULTISIG = 0xae,
	//OP_CHECKMULTISIGVERIFY = 0xaf,

	// expansion
	//OP_NOP1 = 0xb0,
	//OP_CHECKLOCKTIMEVERIFY = 0xb1,
	//OP_NOP2 = OP_CHECKLOCKTIMEVERIFY,
	//OP_CHECKSEQUENCEVERIFY = 0xb2,
	//OP_NOP3 = OP_CHECKSEQUENCEVERIFY,
	//OP_NOP4 = 0xb3,
	//OP_NOP5 = 0xb4,
	//OP_NOP6 = 0xb5,
	//OP_NOP7 = 0xb6,
	//OP_NOP8 = 0xb7,
	//OP_NOP9 = 0xb8,
	//OP_NOP10 = 0xb9,

	// template matching params
	//OP_SMALLINTEGER = 0xfa,
	//OP_PUBKEYS = 0xfb,
	//OP_PUBKEYHASH = 0xfd,
	//OP_PUBKEY = 0xfe,

	//OP_INVALIDOPCODE = 0xff,

)
