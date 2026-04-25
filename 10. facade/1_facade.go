package main

type Buffer struct {
	width, height int
	buffer        []rune
}

func NewBuffer(width, height int) *Buffer {
	return &Buffer{width, height,
		make([]rune, width*height)}
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

type Viewport struct {
	buffer *Buffer
	offset int
}

func NewViewport(buffer *Buffer) *Viewport {
	return &Viewport{buffer: buffer}
}

func (v *Viewport) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

// a facade over buffers and viewports
type Console struct {
	buffers   []*Buffer
	viewports []*Viewport
	offset    int
}

func NewConsole() *Console {
	b := NewBuffer(10, 10)
	v := NewViewport(b)
	return &Console{[]*Buffer{b}, []*Viewport{v}, 0}
}

func (c *Console) GetCharacterAt(index int) rune { //
	return c.viewports[0].GetCharacterAt(index)
}

func main() {
	c := NewConsole()
	u := c.GetCharacterAt(1)
}

// NewConsole() builds the buffer and viewport internally —
// callers do not need to know the order or how they are wired together.

// GetCharacterAt does not force you to pick which viewport,
// which buffer, or what offset to use; in this example it simply forwards to viewports[0].

// What the Facade is doing here

// It simplifies access to a subsystem that could be fiddly (many objects, many indices).
// The client only needs NewConsole() and c.GetCharacterAt(i) —
// that is the Facade idea: one entry point in front of several types behind the scenes.

// A Facade is a high-level, thin, easy-to-use API built on top of a more complex subsystem
// (many types, many steps, many parameters). Clients talk to a single entry point
// (for example Console, an API gateway, or OrderService.placeOrder()),
// instead of having to know the call order for Buffer, Viewport, PaymentClient, and so on.
