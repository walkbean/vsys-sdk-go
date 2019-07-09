package vsys

import(
	rand2 "crypto/rand"
	"encoding/binary"
)

func genRandomBytes(n int) []byte {
	var retBytes []byte
	for i := n / 8; i > 0; i-- {
		var rb = make([]byte, 8)
		rand2.Read(rb)
		retBytes = append(retBytes, rb...)
	}
	return retBytes
}

func bytesToByteArrayWithSize(bytes []byte) (result []byte) {
	result = append(result, uint16ToByte(int16(len(bytes)))...)
	result = append(result, bytes...)
	return
}

func uint64ToByte(data int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(data))
	return b
}

func uint32ToByte(data int32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(data))
	return b
}

func uint16ToByte(data int16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(data))
	return b
}
