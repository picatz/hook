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

func TestBytesToHeaders_new(t *testing.T) {
	headerBytes := []byte{
		7, 0, 0, 0, 10, 0, 0, 0, 14, 0, 0, 0, 5, 0, 0, 0, 1,
		0, 0, 0, 7, 0, 0, 0, 3, 0, 0, 0, 10, 0, 0, 0, 11, 0,
		0, 0, 6, 0, 0, 0, 3, 0, 0, 0, 17, 0, 0, 0, 5, 0, 0,
		0, 12, 0, 0, 0, 36, 0, 0, 0, 58, 97, 117, 116, 104,
		111, 114, 105, 116, 121, 0, 108, 111, 99, 97, 108, 104,
		111, 115, 116, 58, 57, 48, 57, 48, 0, 58, 112, 97, 116, 104,
		0, 47, 0, 58, 109, 101, 116, 104, 111, 100, 0, 71, 69, 84,
		0, 117, 115, 101, 114, 45, 97, 103, 101, 110, 116, 0, 99, 117,
		114, 108, 47, 55, 46, 54, 52, 46, 49, 0, 97, 99, 99, 101, 112,
		116, 0, 42, 47, 42, 0, 120, 45, 102, 111, 114, 119, 97, 114, 100,
		101, 100, 45, 112, 114, 111, 116, 111, 0, 104, 116, 116, 112, 115,
		0, 120, 45, 114, 101, 113, 117, 101, 115, 116, 45, 105, 100, 0,
		99, 97, 49, 52, 51, 53, 101, 52, 45, 53, 52, 52, 53, 45, 52, 48,
		49, 53, 45, 57, 56, 52, 55, 45, 100, 52, 52, 55, 52, 99, 53,
		48, 100, 99, 51, 49, 0,
	}

	headers := BytesToHeaders(headerBytes)

	require.Equal(t, 7, len(headers))

	expectedHeaders := Headers{
		Header{":authority", "localhost:9090"},
		Header{":path", "/"},
		Header{":method", "GET"},
		Header{"user-agent", "curl/7.64.1"},
		Header{"accept", "*/*"},
		Header{"x-forwarded-proto", "https"},
		Header{"x-request-id", "ca1435e4-5445-4015-9847-d4474c50dc31"},
	}

	for i, expectedHeader := range expectedHeaders {
		require.Equal(t, expectedHeader[0], headers[i][0])
		require.Equal(t, expectedHeader[1], headers[i][1])
	}
}
