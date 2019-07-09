package vsys

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenRandomBytes(t *testing.T) {
	fmt.Print(genRandomBytes(64))
}

func TestAccount_BuildFromSeed(t *testing.T) {
	acc1 := InitAccount(Mainnet)
	acc1.BuildFromSeed("nephew hurry tent airport upon tape lonely enough noise sorry address almost drama apple best", 0)
	assert.Equal(t, acc1.PrivateKey(), "5MCdijm6ayZk4CzxDBmpy8Xs6RSci9tsBF1vQNgauCoP")
	assert.Equal(t, acc1.Address(), "AR8z62ZuyQGkmBXbtpSyPbZuQkYpmcQeXk7")
}

func TestAccount_BuildFromPrivateKey(t *testing.T) {
	acc := InitAccount(Testnet)
	acc.BuildFromPrivateKey("DeYZPEQ1xWLDnKvHnx5wWWbjTTHdU93AWNK6WAv54MmS")
	assert.Equal(t, "AU5XoSXq8KcPnxEp1zLSD9xoLRkDj8bzJY4", acc.Address())
	assert.Equal(t, "57JxWCaasg3ToZHo6iNDLhVijji4aBSHFkFyJmomnwpF", acc.PublicKey())
	assert.Equal(t, "DeYZPEQ1xWLDnKvHnx5wWWbjTTHdU93AWNK6WAv54MmS", acc.PrivateKey())
}
