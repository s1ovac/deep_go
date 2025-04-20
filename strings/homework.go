package strings

import "unsafe"

type COWBuffer struct {
	data []byte
	refs *int
}

func NewCOWBuffer(data []byte) COWBuffer {
	return COWBuffer{
		data: data,
		refs: new(int),
	}
}

// Clone - создать новую копию буфера.
//
// Предположим, что все будут производить копирование
// буффера только с использованием метода Clone()
func (b *COWBuffer) Clone() COWBuffer {
	if b.refs != nil {
		*b.refs++
	}

	return COWBuffer{
		data: b.data,
	}
}

// Close - перестать использовать копию буффера.
func (b *COWBuffer) Close() {
	if b.refs != nil && *b.refs > 0 {
		*b.refs--
	}
}

// Update - изменить определенный байт в буффере.
func (b *COWBuffer) Update(index int, value byte) bool {
	if len(b.data) == 0 || index < 0 || index >= len(b.data) {
		return false
	}

	if b.refs != nil && *b.refs > 0 {
		*b.refs--

		newData := make([]byte, len(b.data))
		copy(newData, b.data)
		b.data = newData
	}

	b.data[index] = value

	return true
}

// String - сконвертировать буффер в строку.
func (b *COWBuffer) String() string {
	return unsafe.String(unsafe.SliceData(b.data), len(b.data))
}
