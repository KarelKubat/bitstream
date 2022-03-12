package bitstream

// Bit represents a binary digit, with the values `Zero` or `One`.
type Bit int

// `Bit` values are `Zero` or `One`.
const (
	Zero Bit = iota
	One
)

func (b Bit) String() string {
	return []string{"0", "1"}[b]
}
