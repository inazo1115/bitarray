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

func newBitArray(size int) *BitArray {
	bitIdx, _ := getOffsets(size)
	data := make([]uint32, bitIdx+1)
	return &BitArray{size, data}
}

func NewBitArray(size int, val bool) *BitArray {
	ret := newBitArray(size)
	if val {
		ret.Invert()
	}
	return ret
}

func NewBitArrayWithInit(init []bool) *BitArray {
	ret := newBitArray(len(init))
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
	bitIdx, subIdx := getOffsets(bits.size)
	for i := 0; i < bitIdx; i++ {
		if bits.data[i]^other.data[i] != 0 {
			return false
		}
	}
	mask := uint32((1 << uint(subIdx)) - 1)
	return bits.data[bitIdx]^other.data[bitIdx]&mask == 0
}

func (bits *BitArray) Get(idx int) (bool, error) {
	if idx >= bits.size {
		return false, fmt.Errorf("out of index: %d >= %d", idx, bits.size)
	}
	bitIdx, subIdx := getOffsets(idx)
	return bits.data[bitIdx]&(1<<uint(subIdx)) > 0, nil
}

func (bits *BitArray) SubArray(from, to int) (*BitArray, error) {

	if from > to {
		return nil, fmt.Errorf("'from' must be smaller than 'to': %v > %v", from, to)
	}

	if to > bits.size {
		to = bits.size
	}

	fromBitIdx, fromSubIdx := getOffsets(from)
	newSize := to - from
	newBitIdx, _ := getOffsets(newSize)
	newData := make([]uint32, newBitIdx+1)

	if from/internalBitSize == to/internalBitSize {
		newData[0] = bits.data[fromBitIdx] >> uint(fromSubIdx)
		return &BitArray{newSize, newData}, nil
	}

	for i := 0; i < len(newData); i++ {
		lower := bits.data[fromBitIdx+i] >> uint(fromSubIdx)
		upper := bits.data[fromBitIdx+i+1] << uint(internalBitSize-fromSubIdx)
		newData[i] = lower | upper
	}

	return &BitArray{newSize, newData}, nil
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

func (bits *BitArray) Access(idx int) (bool, error) {
	return bits.Get(idx)
}

func (bits *BitArray) Rank(val bool, idx int) (int, error) {
	if idx >= bits.size {
		return 0, fmt.Errorf("out of index: %d >= %d", idx, bits.size)
	}
	ret := 0
	for i := 0; i < idx; i++ {
		if b, _ := bits.Access(i); b == val {
			ret++
		}
	}
	return ret, nil
}

func (bits *BitArray) Select(val bool, ith int) (int, error) {
	count := 0
	for i := 0; i < bits.size; i++ {
		b, _ := bits.Access(i)
		if b == val {
			count++
		}
		if count == ith+1 {
			return i, nil
		}
	}
	return 0, fmt.Errorf("bits doesn't have %d + 1 %t", ith, val)
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

func (bits *BitArray) Data() []uint32 {
	return bits.data
}

func (bits *BitArray) Int() int {
	ret := 0
	x := 1
	for i := bits.size - 1; i >= 0; i-- {
		b, _ := bits.Get(i)
		if b {
			ret += x
		}
		x *= 2
	}
	return ret
}

func getOffsets(idx int) (int, int) {
	bitIdx := idx / internalBitSize
	subIdx := idx % internalBitSize
	return bitIdx, subIdx
}
