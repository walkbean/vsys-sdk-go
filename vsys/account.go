package vsys

import (
	"golang.org/x/crypto/curve25519"
)

type Account struct {
	publicKey  []byte
	privateKey []byte
	network    byte
	accSeed    string
}

// Address get address with base58 encoded
func (acc *Account) Address() string {
	unCheckSumAddress := make([]byte, 0)
	unCheckSumAddress = append(unCheckSumAddress, addrVersion, acc.network)
	unCheckSumAddress = append(unCheckSumAddress, HashChain(acc.publicKey)[:20]...)
	address := Base58Encode(append(unCheckSumAddress, HashChain(unCheckSumAddress)[:4]...))
	return address
}

func (acc *Account) PrivateKey() string {
	return Base58Encode(acc.privateKey)
}

func (acc *Account) PublicKey() string {
	return Base58Encode(acc.publicKey)
}

func (acc *Account) AccountSeed() string {
	return acc.accSeed
}

// SignData sign data bytes and
// the output is base58 encoded data
func (acc *Account) SignData(data []byte) string {
	return Base58Encode(Sign(acc.privateKey, data, genRandomBytes(64)))
}

// VerifySignature check if signature is correct
func (acc *Account) VerifySignature(data, signature []byte) bool {
	return Verify(acc.publicKey, data, signature) == 1
}

// InitAccount return account with network initiated
func InitAccount(network byte) *Account {
	return &Account{network: network}
}

// BuildFromPrivateKey build account using privateKey
func (acc *Account) BuildFromPrivateKey(privateKey string) {
	var bPrivateKey [32]byte
	var originPublicKey = new([32]byte)
	copy(bPrivateKey[:], Base58Decode(privateKey)[:])
	curve25519.ScalarBaseMult(originPublicKey, &bPrivateKey)
	acc.publicKey = originPublicKey[:]
	acc.privateKey = bPrivateKey[:]
}

// BuildFromPrivateKey build account using seed and nonce
func (acc *Account) BuildFromSeed(seed string, nonce int) {
	seedHash := BuildSeedHash(seed, nonce)
	keyPair := GenerateKeyPair(seedHash)
	acc.publicKey = keyPair.publicKey
	acc.privateKey = keyPair.privateKey
	acc.accSeed = seed
}

// BuildPayment build payment transaction
// recipient should be address
// amount is in minimum unit
// attachment can be empty
func (acc *Account) BuildPayment(recipient string, amount int64, attachment []byte) *Transaction {
	transaction := NewPaymentTransaction(recipient, amount, attachment)
	transaction.SenderPublicKey = acc.PublicKey()
	transaction.Signature = acc.SignData(transaction.BuildTxData())
	return transaction
}

// BuildLeasing build leasing transaction
// recipient should be address
// amount is in minimum unit
func (acc *Account) BuildLeasing(recipient string, amount int64) *Transaction {
	transaction := NewLeaseTransaction(recipient, amount)
	transaction.SenderPublicKey = acc.PublicKey()
	transaction.Signature = acc.SignData(transaction.BuildTxData())
	return transaction
}

// BuildCancelLeasing build Cancel transaction
func (acc *Account) BuildCancelLeasing(txId string) *Transaction {
	transaction := NewCancelLeaseTransaction(txId)
	transaction.SenderPublicKey = acc.PublicKey()
	transaction.Signature = acc.SignData(transaction.BuildTxData())
	return transaction
}
