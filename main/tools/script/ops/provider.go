package ops

// Provider接口
type Provider interface {
	// 由字节码获得操作符
	GetOp(uint8) (Op, bool)
	// 由名称获得字节码
	GetCode(string) (uint8, bool)
}

type provider struct {
	ops   map[uint8]Op
	codes map[string]uint8
}

func (p *provider) GetOp(code uint8) (Op, bool) {
	op, status := p.ops[code]
	return op, status
}
func (p *provider) GetCode(str string) (uint8, bool) {
	code, status := p.codes[str]
	return code, status
}
