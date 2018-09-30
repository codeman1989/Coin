package script

import "io"

type context struct {
	*stack
	alt    *stack
	reader io.Reader
}

func (c *context) Pop() []byte                 { return c.stack.Pop() }
func (c *context) PopAlt() []byte              { return c.alt.Pop() }
func (c *context) Push(value []byte)           { c.stack.Push(value) }
func (c *context) PushAlt(value []byte)        { c.alt.Push(value) }
func (c *context) Read(bs []byte) (int, error) { return c.reader.Read(bs) }
func (c *context) Size() int                   { return c.stack.Size() }
func (c *context) SizeAlt() int                { return c.alt.Size() }
