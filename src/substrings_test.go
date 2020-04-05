package src

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var input1 = "ABCDEABDPKMABBAB"
var input2 = "ABCDEABDPKMABBCDEABDPKMABBCDEABDPKMAB"
var input3 = "ACBCDEBDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPBAKMABCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABB"
var input4 = "ACBCDEBDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPFFFFFFFFFFFFFFFFFFFFFFFGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPBAKMABCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABBCDEABDPKMABB"


func Test1(t *testing.T) {
	assert.Equal(t, CheckString(input1), "YES")
}

func Test2(t *testing.T) {
	assert.Equal(t, CheckString(input2), "NO")
}

func Test3(t *testing.T) {
	assert.Equal(t, CheckString(input3), "YES")
}


func Benchmark1(b *testing.B) {
	benchmarkSubstrings(input1, b)
}

func Benchmark2(b *testing.B) {
	benchmarkSubstrings(input2, b)
}

func Benchmark3(b *testing.B) {
	benchmarkSubstrings(input3, b)
}

func BenchmarkFun(b *testing.B) {
	benchmarkSubstrings(input4, b)
}

func Benchmark4(b *testing.B) {
	benchmarkSubstrings2(input1, b)
}

func Benchmark5(b *testing.B) {
	benchmarkSubstrings2(input2, b)
}

func Benchmark6(b *testing.B) {
	benchmarkSubstrings2(input3, b)
}

func BenchmarkFun2(b *testing.B) {
	benchmarkSubstrings2(input4, b)
}

func benchmarkSubstrings(input string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		CheckString(input)
	}
}

func benchmarkSubstrings2(input string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		CheckString2(input)
	}
}