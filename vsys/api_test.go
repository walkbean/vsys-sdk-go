package vsys

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

const testPrivateKey = "TEST_PRIVATE_KEY"

func TestSendPaymentTx(t *testing.T) {
	InitApi("https://wallet.v.systems/api", MainnetByte)
	acc := InitAccount(MainnetByte)
	acc.BuildFromPrivateKey(testPrivateKey)
	tx := acc.BuildPayment("ARMNeqmATUgKy1hj866Pva3oyaTyyYtLjUv", 1e7, []byte{})
	resp, err := SendPaymentTx(tx)
	assert.Equal(t, nil, err)
	assert.Equal(t, resp.Error, 0)
}

func TestSendLeasingTx(t *testing.T) {
	InitApi("https://wallet.v.systems/api", MainnetByte)
	acc := InitAccount(MainnetByte)
	acc.BuildFromPrivateKey(testPrivateKey)
	tx := acc.BuildLeasing("ARMNeqmATUgKy1hj866Pva3oyaTyyYtLjUv", 1e7)
	resp, err := SendLeasingTx(tx)
	assert.Equal(t, nil, err)
	assert.Equal(t, resp.Error, 0)
	fmt.Println("debug lease tx id: ", resp.Id)
}

func TestSendCancelLeasingTx(t *testing.T) {
	InitApi("https://wallet.v.systems/api", MainnetByte)
	acc := InitAccount(MainnetByte)
	acc.BuildFromPrivateKey(testPrivateKey)
	tx := acc.BuildCancelLeasing("fDVDQbBii2tUsefApxBXEsQ2KeBdY2zG2MxuuZvx7NY")
	resp, err := SendCancelLeasingTx(tx)
	assert.Equal(t, nil, err)
	assert.Equal(t, resp.Error, 0)
}
