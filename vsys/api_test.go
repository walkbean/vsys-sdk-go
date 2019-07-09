package vsys

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testPrivateKey = "4FpJ4hWMMpMW2pX7UMKdrC4WnXQ8KBRqxKSd1xE9MM9t"

func TestSendPaymentTx(t *testing.T) {
	InitApi("https://wallet.v.systems/api", Mainnet)
	acc := InitAccount(Mainnet)
	acc.BuildFromPrivateKey(testPrivateKey)
	tx := acc.BuildPayment("ARMNeqmATUgKy1hj866Pva3oyaTyyYtLjUv", 1e7, []byte{})
	resp, err := SendPaymentTx(tx)
	assert.Equal(t, nil, err)
	assert.Equal(t, resp.Error, 0)
}

func TestSendLeasingTx(t *testing.T) {
	InitApi("https://wallet.v.systems/api", Mainnet)
	acc := InitAccount(Mainnet)
	acc.BuildFromPrivateKey(testPrivateKey)
	tx := acc.BuildLeasing("ARMNeqmATUgKy1hj866Pva3oyaTyyYtLjUv", 1e7)
	resp, err := SendLeasingTx(tx)
	assert.Equal(t, nil, err)
	assert.Equal(t, resp.Error, 0)
	fmt.Println("debug lease tx id: ", resp.Id)
}

func TestSendCancelLeasingTx(t *testing.T) {
	InitApi("https://wallet.v.systems/api", Mainnet)
	acc := InitAccount(Mainnet)
	acc.BuildFromPrivateKey(testPrivateKey)
	tx := acc.BuildCancelLeasing("5p42Z3dL7pbKH8dPWpFvCXzQ6WpMDRUgd1N1FNYseUtv")
	resp, err := SendCancelLeasingTx(tx)
	assert.Equal(t, nil, err)
	assert.Equal(t, resp.Error, 0)
}
