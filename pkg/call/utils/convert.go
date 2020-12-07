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
	var numberOfHeaders = binary.LittleEndian.Uint32(data[0:4])

	var sizeIndexes = make([]int, numberOfHeaders*2)
	for i := 0; i < len(sizeIndexes); i++ {
		x := 4 + i*4
		sizeIndexes[i] = int(binary.LittleEndian.Uint32(data[x : x+4]))
	}

	var (
		sizeIndex int
		dataIndex = 4 * (1 + 2*int(numberOfHeaders))
		headers   = make(Headers, numberOfHeaders)

		getSize = func() int {
			return sizeIndexes[sizeIndex]
		}

		incSize = func() {
			sizeIndex++
			return
		}

		bumpDataIndex = func(size int) {
			dataIndex += size + 1
			return
		}

		str = func(dataIndex, size int) string {
			ptr := data[dataIndex : dataIndex+size]
			str := *(*string)(unsafe.Pointer(&ptr))
			bumpDataIndex(size)
			return str
		}

		setHeader = func(index int, key, value string) {
			headers[index] = Header{key, value}
		}
	)

	for i := range headers {
		var keySize = getSize()
		incSize()
		var key = str(dataIndex, keySize)

		var valueSize = getSize()
		incSize()
		var value = str(dataIndex, valueSize)

		setHeader(i, key, value)
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
