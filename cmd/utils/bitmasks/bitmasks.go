package bitmasks

type BitMask uint8

const (
	READ BitMask = 1 << iota
	WRITE
	EXECUTE
)

func (b *BitMask) Set(flag BitMask) {
	*b = *b | flag
}

func (b *BitMask) Clear(flag BitMask) {
	*b = *b &^ flag
}

func (b *BitMask) Toggle(flag BitMask) {
	*b = *b ^ flag
}

func (b *BitMask) Has(flag BitMask) bool {
	return *b&flag != 0
}
