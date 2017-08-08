package bitarray

import (
	"testing"
)

const (
	_0 = false
	_1 = true
)

func TestNewBitArray(t *testing.T) {

	var tests = []struct {
		size int
		val  bool
		out  string
	}{
		{0, false, "BitArray: size=0, data=[]"},
		{1, false, "BitArray: size=1, data=[0]"},
		{1, true, "BitArray: size=1, data=[1]"},
		{2, false, "BitArray: size=2, data=[0, 0]"},
		{32, false, "BitArray: size=32, data=[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]"},
		{33, false, "BitArray: size=33, data=[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]"},
	}

	for _, test := range tests {
		actual := NewBitArray(test.size, test.val).String()
		expected := test.out
		if actual != expected {
			t.Errorf("NewBitArray(%v, %v) => '%v', want '%v'",
				test.size, test.val, actual, expected)
		}
	}
}

func TestNewBitArrayWithInit(t *testing.T) {

	var tests = []struct {
		in  []bool
		out string
	}{
		{[]bool{}, "BitArray: size=0, data=[]"},
		{[]bool{false}, "BitArray: size=1, data=[0]"},
		{[]bool{true}, "BitArray: size=1, data=[1]"},
		{[]bool{false, true}, "BitArray: size=2, data=[0, 1]"},
	}

	for _, test := range tests {
		actual := NewBitArrayWithInit(test.in).String()
		expected := test.out
		if actual != expected {
			t.Errorf("NewBitArrayWithInit(%v) => '%v', want '%v'",
				test.in, actual, expected)
		}
	}
}

func TestEqual(t *testing.T) {

	var bits0, bits1 *BitArray

	bits0 = NewBitArrayWithInit([]bool{true, true, false, false, true})
	bits1 = NewBitArray(5, false)
	bits1.Set(0, true)
	bits1.Set(1, true)
	bits1.Set(4, true)
	if !bits0.Equal(bits1) {
		t.Errorf("Equal(%v), wants %v", bits1, bits0)
	}

	bits0 = NewBitArrayWithInit([]bool{true, true, false, false, true})
	bits1 = NewBitArray(5, false)
	if bits0.Equal(bits1) {
		t.Errorf("Equal(%v) must be false", bits1)
	}

	bits0 = NewBitArrayWithInit([]bool{true, true, false, false, true})
	bits1 = NewBitArray(5, true)
	if bits0.Equal(bits1) {
		t.Errorf("Equal(%v) must be false", bits1)
	}
}

func TestGet(t *testing.T) {

	bits := NewBitArrayWithInit([]bool{true, true, false, false, true})
	expected := []bool{true, true, false, false, true}

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

	bits := NewBitArrayWithInit([]bool{true, true, false, false, true})
	bits.Set(0, false)
	bits.Set(3, true)
	bits.Set(4, true)
	expected := []bool{false, true, false, true, true}

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

	bits := NewBitArrayWithInit([]bool{true, true, false, false, true})
	bits.Toggle(0)
	bits.Toggle(1)
	bits.Toggle(2)
	expected := []bool{false, false, true, false, true}

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

	bits := NewBitArrayWithInit([]bool{true, true, false, false, true})
	bits.Invert()
	expected := []bool{false, false, true, true, false}

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

	bits := NewBitArray(500, false)

	bits.Set(100, true)
	bits.Set(200, true)
	bits.Set(300, true)
	bits.Set(400, true)

	b, _ := bits.Get(99)
	if b != false {
		t.Errorf("Get(%v) => %v", 99, true)
	}
	b, _ = bits.Get(100)
	if b != true {
		t.Errorf("Get(%v) => %v", 100, false)
	}
	b, _ = bits.Get(101)
	if b != false {
		t.Errorf("Get(%v) => %v", 101, true)
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

	var tests = []struct {
		in  []bool
		val bool
		i   int
		out int
	}{
		{[]bool{_0, _1, _1, _0, _1, _1, _0, _0, _1, _0, _0}, _0, 6, 2},
		{[]bool{_0, _1, _1, _0, _1, _1, _0, _0, _1, _0, _0}, _1, 6, 4},
		{[]bool{_0, _1, _1, _0, _1, _1, _0, _0, _1, _0, _0}, _1, 10, 5},
	}

	for _, test := range tests {
		actual, err := NewBitArrayWithInit(test.in).Rank(test.val, test.i)
		if err != nil {
			t.Errorf(err.Error())
		}
		expected := test.out
		if actual != expected {
			t.Errorf("Rank(%v, %v) => '%v', want '%v'",
				test.val, test.i, actual, expected)
		}
	}
}

func TestSelect(t *testing.T) {

	var tests = []struct {
		in  []bool
		val bool
		i   int
		out int
	}{
		{[]bool{_0, _1, _1, _0, _1, _1, _0, _0, _1, _0, _0}, _0, 1, 3},
		{[]bool{_0, _1, _1, _0, _1, _1, _0, _0, _1, _0, _0}, _0, 3, 7},
		{[]bool{_0, _1, _1, _0, _1, _1, _0, _0, _1, _0, _0}, _1, 3, 5},
	}

	for _, test := range tests {
		actual, err := NewBitArrayWithInit(test.in).Select(test.val, test.i)
		if err != nil {
			t.Errorf(err.Error())
		}
		expected := test.out
		if actual != expected {
			t.Errorf("Select(%v, %v) => '%v', want '%v'",
				test.val, test.i, actual, expected)
		}
	}
}
