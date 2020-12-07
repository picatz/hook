package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBytePtrToByteSlice(t *testing.T) {
	require := require.New(t)

	rawByte := &[]byte("a")[0]
	require.NotNil(rawByte)

	bytes := BytePtrToByteSlice(rawByte, 1)
	require.Equal(1, len(bytes))
	require.Equal("a", string(bytes[0]))
}

func TestBytePtrToString(t *testing.T) {
	require := require.New(t)

	rawByte := &[]byte("a")[0]
	require.NotNil(rawByte)

	str := BytePtrToString(rawByte, 1)
	require.Equal(1, len(str))
	require.Equal("a", str)
}

var resultBenchmarkBytePtrToString string

func BenchmarkBytePtrToString(b *testing.B) {
	rawByte := &[]byte("a")[0]

	var r string
	for n := 0; n < b.N; n++ {
		r = BytePtrToString(rawByte, 1)
	}

	resultBenchmarkBytePtrToString = r
}

var resultBenchmarkBytePtrToByteSlice []byte

func BenchmarkBytePtrToByteSlice(b *testing.B) {
	rawByte := &[]byte("a")[0]

	var r []byte
	for n := 0; n < b.N; n++ {
		r = BytePtrToByteSlice(rawByte, 1)
	}

	resultBenchmarkBytePtrToByteSlice = r
}
