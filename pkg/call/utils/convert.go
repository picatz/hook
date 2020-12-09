package utils

import (
	"encoding/binary"
	"unsafe"
)

/*

   ~ BEWARE TRAVELER! ~
    ._________.
   (/\___   /|\\           ._____________.
         \_/ | \\        //|\   ______/ \)
           \_|  \\      // | \_/\/
             \|\/|\_   //  /\/
              (,,)\ \_//  /\ .
              /./\_\/ /  |  /.
                 |= _ \ /   \\
                 \_= _ \ |  //
                   \==  \|\//
                __(\===\(  )
			   (((~) __(_//
                    (((~)

    ~ HERE BE DRAGONS! ~

*/

func BytePtrToString(raw *byte, size int) string {
	// protect against nil byte ptr
	if raw == nil {
		return ""
	}
	// covert to proper type
	str := *(*string)(unsafe.Pointer(&raw))
	// limit the size of the result
	str = str[:size]
	// now we gucci
	return str
}

func BytePtrToByteSlice(raw *byte, size int) []byte {
	// protect against nil byte ptr
	if raw == nil {
		return []byte{}
	}
	// covert to proper type
	bytes := *(*[]byte)(unsafe.Pointer(&raw))
	// limit the size of the result
	bytes = bytes[:size]
	// now we gucci
	return bytes
}

func StringToBytePtr(str string) *byte {
	// protect against empty string
	if len(str) == 0 {
		return nil
	}
	// othersie cast it to a slice of bytes
	bytes := *(*[]byte)(unsafe.Pointer(&str))
	// and return a pointer to the first one
	return &bytes[0]
}

type Header = [2]string
type Headers = []Header

func BytesToHeaders(data []byte) Headers {
	var (
		numberOfHeaders = binary.LittleEndian.Uint32(data[0:4])
		sizeIndexes     = make([]int, numberOfHeaders*2)
		sizeIndex       = 0
		dataIndex       = 4 * (1 + 2*int(numberOfHeaders))
		headers         = make(Headers, numberOfHeaders)
	)

	for i := 0; i < len(sizeIndexes); i++ {
		sizeIndex := 4 + i*4
		sizeIndexes[i] = int(binary.LittleEndian.Uint32(data[sizeIndex : sizeIndex+4]))
	}

	for i := range headers {
		// get header key string
		keySize := sizeIndexes[sizeIndex]
		sizeIndex++
		keyBytes := data[dataIndex : dataIndex+keySize]
		key := *(*string)(unsafe.Pointer(&keyBytes))
		dataIndex += keySize + 1

		// get header value string
		valueSize := sizeIndexes[sizeIndex]
		sizeIndex++
		valueBytes := data[dataIndex : dataIndex+valueSize]
		value := *(*string)(unsafe.Pointer(&valueBytes))
		dataIndex += valueSize + 1

		// populate returned headers with found header
		headers[i] = Header{key, value}
	}

	return headers
}

const headerBaseSize = 4

func HeadersToBytes(headers Headers) []byte {
	// determine size of returned byte slice
	var bytesSize = headerBaseSize
	for _, header := range headers {
		keySize := len(header[0])
		valueSize := len(header[1])
		bytesSize += keySize + valueSize + ((headerBaseSize * 2) + 2)
	}

	// allocate bytes slice
	var bytes = make([]byte, bytesSize)
	binary.LittleEndian.PutUint32(bytes[0:headerBaseSize], uint32(len(headers)))

	var (
		baseIndex    = headerBaseSize
		incBaseIndex = func() {
			baseIndex++
		}
	)

	for _, m := range headers {
		binary.LittleEndian.PutUint32(bytes[baseIndex:baseIndex+headerBaseSize], uint32(len(m[0])))
		baseIndex += headerBaseSize
		binary.LittleEndian.PutUint32(bytes[baseIndex:baseIndex+headerBaseSize], uint32(len(m[1])))
		baseIndex += headerBaseSize
	}

	// encode headers
	for _, header := range headers {
		// encode key
		for i := 0; i < len(header[0]); i++ {
			bytes[baseIndex] = header[0][i]
			incBaseIndex()
		}
		incBaseIndex()

		// encode value
		for i := 0; i < len(header[1]); i++ {
			bytes[baseIndex] = header[1][i]
			incBaseIndex()
		}
		incBaseIndex()
	}

	return bytes
}

func PropertyPathToBytes(path []string) []byte {
	if len(path) == 0 {
		return []byte{}
	}

	var size int
	for _, p := range path {
		size += len(p) + 1
	}

	var bytes = make([]byte, 0, size)
	for _, p := range path {
		bytes = append(bytes, p...)
		bytes = append(bytes, 0)
	}

	return bytes[:len(bytes)-1]
}
