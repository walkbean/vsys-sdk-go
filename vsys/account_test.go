package vsys

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenRandomBytes(t *testing.T) {
	fmt.Print(genRandomBytes(64))
}

func TestAccount_BuildFromSeed(t *testing.T) {
	account := InitAccount(TestnetByte)
	account.BuildFromSeed("play village reason improve round cloud pyramid absurd process meat gas nest flower hundred garden", 0)
	fmt.Println(account.Address())
	fmt.Println(account.PublicKey())
	acc1 := InitAccount(MainnetByte)
	acc1.BuildFromSeed("nephew hurry tent airport upon tape lonely enough noise sorry address almost drama apple best", 0)
	assert.Equal(t, acc1.PrivateKey(), "5MCdijm6ayZk4CzxDBmpy8Xs6RSci9tsBF1vQNgauCoP")
	assert.Equal(t, acc1.Address(), "AR8z62ZuyQGkmBXbtpSyPbZuQkYpmcQeXk7")
}

func TestAccount_BuildFromPrivateKey(t *testing.T) {
	acc := InitAccount(TestnetByte)
	acc.BuildFromPrivateKey("DeYZPEQ1xWLDnKvHnx5wWWbjTTHdU93AWNK6WAv54MmS")
	fmt.Println(acc.Address())
	fmt.Println(acc.PublicKey())
	fmt.Println(acc.PrivateKey())
}

func TestAccount_VerifySignature(t *testing.T) {
	//data,_ := com.Base64Decode("AhVr3BWva14AAAAAAACYloAAAAAAAJiWgABkBVTZyQaz8crvxWrmih+r2y8f0fHRRU/UGqkABHRlc3Q=")
	account := InitAccount(TestnetByte)
	account.BuildFromSeed("play village reason improve round cloud pyramid absurd process meat gas nest flower hundred garden", 0)
	//tx := NewCancelLeaseTransaction("5S2GHuM5u9vJkZpUCamkm6DoV346mTneHb6SY2kYPye5")
	//tx := NewPaymentTransaction("AU4yV9iAy1ziCW3iyeMjPhZ72uEEGhzXMGv", 10000000)
	//tx.Attachment = []byte("ll")
	//yeego.Print(base58.Encode(tx.Attachment))
	//data := tx.BuildTxData()
	sigData1 := account.SignData(Base58Decode("k1SpcPknx6db36cxd1PLPqYda1rKwLuEs5Lo3XJ8wjVqDompds71ESqXUS6MnV7iQPvE43KcLw"))
	//sigData2 := acc.SignData(data)
	fmt.Print(sigData1)
	fmt.Print(Verify(account.publicKey, Base58Decode("k1SpcPknx6db36cxd1PLPqYda1rKwLuEs5Lo3XJ8wjVqDompds71ESqXUS6MnV7iQPvE43KcLw"), Base58Decode("2kMPMX1WwDmTYvje1KNmw32PQ52HsaPVMfrwDJPG5c3iXxZ597aZg5BFCGBt6H53TWTfFYBpLzeS9xvmLWUznvZP")))
}
