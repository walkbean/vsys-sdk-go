package vsys

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAesDecrypt(t *testing.T) {
	for _, origin := range [][]byte{
		[]byte(""),
		[]byte("1"),
		[]byte("12"),
		[]byte("123"),
		[]byte("1234"),
		[]byte("12345"),
		[]byte("123456"),
		[]byte("1234567"),
		[]byte("12345678"),
		[]byte("123456789"),
		[]byte("1234567890"),
		[]byte("123456789012345"),
		[]byte("1234567890123456"),
		[]byte("12345678901234567"),
	} {
		ob, err1 := AesEncrypt([]byte("1"), origin)
		assert.Equal(t, err1, nil)
		ret, err2 := AesDecrypt([]byte("1"), ob)
		assert.Equal(t, err2, nil)
		assert.Equal(t, ret, origin)
	}
}

func TestAesEncrypt(t *testing.T) {
	data, err := AesEncrypt([]byte("test"), []byte("hello"))
	assert.Equal(t, err, nil)
	fmt.Println(string(data))
}
