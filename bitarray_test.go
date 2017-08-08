package bitarray

import (
	"testing"
)

const (
	_0 = false
	_1 = true
)

func TestNewBitArray(t *testing.T) {

	tests := []struct {
		size     int
		val      bool
		expected string
	}{
		{0, _0, "BitArray: size=0, data=[]"},
		{1, _0, "BitArray: size=1, data=[0]"},
		{1, _1, "BitArray: size=1, data=[1]"},
		{2, _0, "BitArray: size=2, data=[0, 0]"},
		{32, _0, "BitArray: size=32, data=[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]"},
		{33, _0, "BitArray: size=33, data=[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]"},
	}

	for _, test := range tests {
		actual := NewBitArray(test.size, test.val).String()
		if actual != test.expected {
			t.Errorf("NewBitArray(%v, %v) => '%v', want '%v'",
				test.size, test.val, actual, test.expected)
		}
	}
}

func TestNewBitArrayWithInit(t *testing.T) {

	tests := []struct {
		bits     []bool
		expected string
	}{
		{[]bool{}, "BitArray: size=0, data=[]"},
		{[]bool{_0}, "BitArray: size=1, data=[0]"},
		{[]bool{_1}, "BitArray: size=1, data=[1]"},
		{[]bool{_0, _1}, "BitArray: size=2, data=[0, 1]"},
	}

	for _, test := range tests {
		actual := NewBitArrayWithInit(test.bits).String()
		if actual != test.expected {
			t.Errorf("NewBitArrayWithInit(%v) => '%v', want '%v'",
				test.bits, actual, test.expected)
		}
	}
}

func TestEqual(t *testing.T) {

	var bits0, bits1 *BitArray

	bits0 = NewBitArrayWithInit([]bool{_1, _1, _0, _0, _1})
	bits1 = NewBitArray(5, _0)
	bits1.Set(0, _1)
	bits1.Set(1, _1)
	bits1.Set(4, _1)
	if !bits0.Equal(bits1) {
		t.Errorf("Equal(%v), wants %v", bits1, bits0)
	}

	bits0 = NewBitArrayWithInit([]bool{_1, _1, _0, _0, _1})
	bits1 = NewBitArray(5, _0)
	if bits0.Equal(bits1) {
		t.Errorf("Equal(%v) must be false", bits1)
	}

	bits0 = NewBitArrayWithInit([]bool{_1, _1, _0, _0, _1})
	bits1 = NewBitArray(5, _1)
	if bits0.Equal(bits1) {
		t.Errorf("Equal(%v) must be false", bits1)
	}
}

func TestGet(t *testing.T) {

	bits := NewBitArrayWithInit([]bool{_1, _1, _0, _0, _1})
	expected := []bool{_1, _1, _0, _0, _1}

	for i, b := range expected {
		actual, err := bits.Get(i)
		if err != nil {
			t.Errorf(err.Error())
		}
		if actual != b {
			t.Errorf("Get(%v) => %v", i, actual)
		}
	}

	_, err := bits.Get(5)
	if err.Error() != "out of index: 5 >= 5" {
		t.Errorf("Get(5) must return error")
	}
}

func TestSet(t *testing.T) {

	bits := NewBitArrayWithInit([]bool{_1, _1, _0, _0, _1})
	bits.Set(0, _0)
	bits.Set(3, _1)
	bits.Set(4, _1)
	expected := []bool{_0, _1, _0, _1, _1}

	for i, b := range expected {
		actual, err := bits.Get(i)
		if err != nil {
			t.Errorf(err.Error())
		}
		if actual != b {
			t.Errorf("Set(%v) => %v", i, actual)
		}
	}
}

func TestToggle(t *testing.T) {

	bits := NewBitArrayWithInit([]bool{_1, _1, _0, _0, _1})
	bits.Toggle(0)
	bits.Toggle(1)
	bits.Toggle(2)
	expected := []bool{_0, _0, _1, _0, _1}

	for i, b := range expected {
		actual, err := bits.Get(i)
		if err != nil {
			t.Errorf(err.Error())
		}
		if actual != b {
			t.Errorf("Toggle(%v) => %v", i, actual)
		}
	}
}

func TestInvert(t *testing.T) {

	bits := NewBitArrayWithInit([]bool{_1, _1, _0, _0, _1})
	bits.Invert()
	expected := []bool{_0, _0, _1, _1, _0}

	for i, b := range expected {
		actual, err := bits.Get(i)
		if err != nil {
			t.Errorf(err.Error())
		}
		if actual != b {
			t.Errorf("Invert(%v) => %v", i, actual)
		}
	}
}

func TestBigSize(t *testing.T) {

	bits := NewBitArray(500, _0)

	bits.Set(100, _1)
	bits.Set(200, _1)
	bits.Set(300, _1)
	bits.Set(400, _1)

	b, _ := bits.Get(99)
	if b != _0 {
		t.Errorf("Get(%v) => %v", 99, _1)
	}
	b, _ = bits.Get(100)
	if b != _1 {
		t.Errorf("Get(%v) => %v", 100, _0)
	}
	b, _ = bits.Get(101)
	if b != _0 {
		t.Errorf("Get(%v) => %v", 101, _1)
	}

	count := 0
	for i := 0; i < bits.Size(); i++ {
		b, _ = bits.Get(i)
		if b {
			count++
		}
	}
	if count != 4 {
		t.Errorf("the number of true is expected 4, but %d", count)
	}

	bits.Toggle(150)
	bits.Toggle(250)
	bits.Toggle(350)
	bits.Toggle(450)

	count = 0
	for i := 0; i < bits.Size(); i++ {
		b, _ = bits.Get(i)
		if b {
			count++
		}
	}
	if count != 8 {
		t.Errorf("the number of true is expected 8, but %d", count)
	}

	bits.Invert()

	count = 0
	for i := 0; i < bits.Size(); i++ {
		b, _ = bits.Get(i)
		if b {
			count++
		}
	}
	if count != 492 {
		t.Errorf("the number of true is expected 492, but %d", count)
	}
}

func TestRank(t *testing.T) {

	tests := []struct {
		bits     []bool
		val      bool
		idx      int
		expected int
	}{
		{[]bool{_0, _1, _1, _0, _1, _1, _0, _0, _1, _0, _0}, _0, 6, 2},
		{[]bool{_0, _1, _1, _0, _1, _1, _0, _0, _1, _0, _0}, _1, 6, 4},
		{[]bool{_0, _1, _1, _0, _1, _1, _0, _0, _1, _0, _0}, _1, 10, 5},
	}

	for _, test := range tests {
		actual, err := NewBitArrayWithInit(test.bits).Rank(test.val, test.idx)
		if err != nil {
			t.Errorf(err.Error())
		}
		if actual != test.expected {
			t.Errorf("Rank(%v, %v) => '%v', want '%v'",
				test.val, test.idx, actual, test.expected)
		}
	}
}

func TestSelect(t *testing.T) {

	tests := []struct {
		bits     []bool
		val      bool
		ith      int
		expected int
	}{
		{[]bool{_0, _1, _1, _0, _1, _1, _0, _0, _1, _0, _0}, _0, 1, 3},
		{[]bool{_0, _1, _1, _0, _1, _1, _0, _0, _1, _0, _0}, _0, 3, 7},
		{[]bool{_0, _1, _1, _0, _1, _1, _0, _0, _1, _0, _0}, _1, 3, 5},
	}

	for _, test := range tests {
		actual, err := NewBitArrayWithInit(test.bits).Select(test.val, test.ith)
		if err != nil {
			t.Errorf(err.Error())
		}
		if actual != test.expected {
			t.Errorf("Select(%v, %v) => '%v', want '%v'",
				test.val, test.ith, actual, test.expected)
		}
	}
}
