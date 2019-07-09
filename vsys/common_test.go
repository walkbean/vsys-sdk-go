package vsys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublicKeyToAddress(t *testing.T) {
	acc1 := InitAccount(Mainnet)
	acc1.BuildFromSeed("nephew hurry tent airport upon tape lonely enough noise sorry address almost drama apple best", 0)
	assert.Equal(t, "AR8z62ZuyQGkmBXbtpSyPbZuQkYpmcQeXk7", acc1.Address())
	assert.Equal(t, "AR8z62ZuyQGkmBXbtpSyPbZuQkYpmcQeXk7", PublicKeyToAddress(acc1.PublicKey(), Mainnet))
}

func TestIsValidateAddress(t *testing.T) {
	ok := IsValidateAddress("AR8z62ZuyQGkmBXbtpSyPbZuQkYpmcQeXk7", Mainnet)
	assert.Equal(t, ok, true)
	ok = IsValidateAddress("AR8z62ZuyQGkmBXbtpSyPbZuQkYpmcQedf3", Mainnet)
	assert.Equal(t, ok, false)
	ok = IsValidateAddress("AU5XoSXq8KcPnxEp1zLSD9xoLRkDj8bzJY4", Testnet)
	assert.Equal(t, ok, true)
}

func TestIsValidatePhrase(t *testing.T) {
	ok := IsValidatePhrase("nephew hurry tent airport upon tape lonely enough noise sorry address almost drama apple best")
	assert.Equal(t, ok, true)
	ok = IsValidatePhrase("great hurry tent airport upon tape lonely enough noise sorry address almost drama apple best ")
	assert.Equal(t, ok, false)
}
