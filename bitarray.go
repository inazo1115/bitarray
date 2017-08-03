package bitarray

import (
	"fmt"
	"math"
	"strings"
)

const internalBitSize = 32

type BitArray struct {
	size int
	data []uint32
}

func NewBitArray(size int, val bool) *BitArray {
	bitIdx, _ := getOffsets(size)
	data := make([]uint32, bitIdx+1)
	ret := &BitArray{size, data}
	if val {
		ret.Invert()
	}
	return ret
}

func NewBitArrayWithInit(init []bool) *BitArray {
	bitIdx, _ := getOffsets(len(init))
	data := make([]uint32, bitIdx+1)
	ret := &BitArray{len(init), data}
	for i, b := range init {
		if b {
			ret.Set(i, b)
		}
	}
	return ret
}

func (bits *BitArray) Equal(other *BitArray) bool {
	if bits.size != other.size {
		return false
	}
	bitIdx, _ := getOffsets(bits.size)
	for i := 0; i < bitIdx+1; i++ {
		if bits.data[i]^other.data[i] != 0 {
			return false
		}
	}
	return true
}

func (bits *BitArray) Get(idx int) (bool, error) {
	if idx >= bits.size {
		return false, fmt.Errorf("out of index: %d >= %d", idx, bits.size)
	}
	bitIdx, subIdx := getOffsets(idx)
	return bits.data[bitIdx]&(1<<uint(subIdx)) > 0, nil
}

func (bits *BitArray) Set(idx int, val bool) {
	bitIdx, subIdx := getOffsets(idx)
	if val {
		bits.data[bitIdx] |= 1 << uint(subIdx)
	} else {
		bits.data[bitIdx] &= (1 << uint(subIdx)) ^ math.MaxUint32
	}
}

func (bits *BitArray) Toggle(idx int) {
	bitIdx, subIdx := getOffsets(idx)
	bits.data[bitIdx] ^= 1 << uint(subIdx)
}

func (bits *BitArray) Invert() {
	bitIdx, _ := getOffsets(bits.size)
	for i := 0; i < bitIdx+1; i++ {
		bits.data[i] ^= math.MaxUint32
	}
}

func (bits *BitArray) String() string {
	tmp := make([]string, bits.size)
	for i := 0; i < bits.size; i++ {
		b, _ := bits.Get(i)
		if b {
			tmp[i] = "1"
		} else {
			tmp[i] = "0"
		}
	}
	data := strings.Join(tmp, ", ")
	return fmt.Sprintf("BitArray: size=%d, data=[%s]", bits.size, data)
}

func (bits *BitArray) Size() int {
	return bits.size
}

func getOffsets(idx int) (int, int) {
	bitIdx := idx / internalBitSize
	subIdx := idx % internalBitSize
	return bitIdx, subIdx
}
