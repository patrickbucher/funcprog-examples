package main

import "fmt"

// TODO
// bind
// unit: buf -> (buf, error)
// lift: func(buf) (buf, error) -> func(buf, error) (buf, error)

type Buffer struct {
	Data []byte
	Size int
	Tip  int
}

func NewBuffer(size int) *Buffer {
	data := make([]byte, size)
	return &Buffer{data, size, 0}
}

func (b *Buffer) Add(x byte) *Buffer {
	if b.Tip < b.Size {
		b.Data[b.Tip] = x
		b.Tip++
		return b
	} else {
		size := b.Size * 2
		tip := b.Tip
		data := make([]byte, size)
		copy(data, b.Data[0:tip])
		data[tip] = x
		return &Buffer{data, size, tip + 1}
	}
}

func (b *Buffer) String() string {
	return fmt.Sprintf("data: %v, size: %d, tip: %d", b.Data, b.Size, b.Tip)
}

func main() {
	b := NewBuffer(3)
	b = b.Add(1)
	fmt.Println(b)
	b = b.Add(2)
	fmt.Println(b)
	b = b.Add(3)
	fmt.Println(b)
	b = b.Add(4)
	fmt.Println(b)
}
