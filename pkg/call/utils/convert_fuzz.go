// +build gofuzz

package utils

import fuzz "github.com/google/gofuzz"

func FuzzBytesToHeaders(data []byte) int {
	headers := BytesToHeaders(data)
	if headers == nil {
		return 0
	}
	return 1
}

func FuzzHeadersToBytes(data []byte) int {
	var (
		key string
		val string

		keyValPair1 [2]string
		keyValPair2 [2]string
	)

	// cannot use HeadersToBytes(headers) (value of type []byte) as [][2]string value in assignment
	fuzz.NewFromGoFuzz(data).Fuzz(&key)
	fuzz.NewFromGoFuzz(data).Fuzz(&val)
	fuzz.NewFromGoFuzz(data).Fuzz(&keyValPair1)
	fuzz.NewFromGoFuzz(data).Fuzz(&keyValPair2)

	var headers = [][2]string{
		{key, val},
		keyValPair1,
		keyValPair2,
	}

	bytes := HeadersToBytes(headers)
	if bytes == nil {
		return 0
	}

	headers = BytesToHeaders(bytes)
	if headers == nil {
		panic("could not covert bytes back to headers type")
	}

	return 1
}

func FuzzPropertyPathToBytes(data []byte) int {
	var (
		path []string
	)
	fuzz.NewFromGoFuzz(data).Fuzz(&path)
	bytes := PropertyPathToBytes(path)
	if bytes == nil {
		return 0
	}
	return 1
}
