package mixedradix

import (
	"fmt"
	"slices"
)

type MixedRadix struct {
	bits  []uint
	signs [][]string
	value []uint
	max   uint
}

// Return true when reached max value
func (mr *MixedRadix) Add(n uint) bool {
	return mr.Set(mr.Decimal() + n)
}

// Return true when reached max value
func (mr *MixedRadix) Set(value uint) bool {
	for i, bit := range mr.bits {
		n := value / bit
		mr.value[i] = n
		if n >= uint(len(mr.signs[i])) {
			mr.SetMax()
			return true
		}
		value = value % bit
	}
	return false
}

func (mr *MixedRadix) SetMax() {
	for i := range mr.value {
		mr.value[i] = uint(len(mr.signs[i])) - 1
	}
}

func (mr *MixedRadix) Decimal() uint {
	d := uint(0)
	for i, v := range mr.value {
		d += v * mr.bits[i]
	}
	return d
}

func (mr *MixedRadix) String() string {
	str := ""
	for i, v := range mr.value {
		str += mr.signs[i][v]
	}
	return str
}

func (mr *MixedRadix) Format(pattern string) string {
	var signs []any
	for i, v := range mr.value {
		signs = append(signs, mr.signs[i][v])
	}
	return fmt.Sprintf(pattern, signs...)
}

func (mr *MixedRadix) Dump() string {
	return fmt.Sprintf("Value %v\nBits %v\nMax %v", mr.value, mr.bits, mr.max)
}

func NewMixedRadix(signs ...[]string) *MixedRadix {
	mr := &MixedRadix{}
	mr.max = 1

	slices.Reverse(signs)
	for _, sign := range signs {
		n := uint(len(sign))
		mr.bits = append(mr.bits, mr.max)
		mr.signs = append(mr.signs, sign)
		mr.value = append(mr.value, 0)
		mr.max *= n
	}
	slices.Reverse(mr.bits)
	slices.Reverse(mr.signs)
	return mr
}
