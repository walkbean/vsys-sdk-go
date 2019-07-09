package vsys

import (
	"golang.org/x/crypto/sha3"
	"golang.org/x/crypto/blake2b"
	"strconv"
	"crypto/sha256"
	"golang.org/x/crypto/curve25519"
)

// Keccak256 calculates and returns the Keccak256 hash of the input data.
func Keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

func HashChain(nonceSecret []byte) []byte {
	blake2bHash, err := blake2b.New256(nil)
	if err != nil {
		panic(err.Error())
	}
	blake2bHash.Write(nonceSecret)
	return Keccak256(blake2bHash.Sum(nil))
}

func BuildSeedHash(seed string, nonce int) []byte {
	nonceSeed := strconv.Itoa(nonce) + seed
	return HashChain([]byte(nonceSeed))
}

func IsValidateAddress(address string, network byte) bool {
	data := Base58Decode(address)
	if len(data) != 26 {
		return false
	}
	if data[0] != addrVersion && data[1] != network {
		return false
	}
	key := data[0:22]
	check := data[22:26]
	keyHash := HashChain(key)[0:4]
	for i := 0; i < 4; i ++ {
		if check[i] != keyHash[i] {
			return false
		}
	}
	return true
}

// GenerateKeyPair generate Account using seed byte array
func GenerateKeyPair(seed []byte) *Account {
	var originPublicKey = new([32]byte)
	originPrivateKey := sha256.Sum256([]byte(seed))
	curve25519.ScalarBaseMult(originPublicKey, &originPrivateKey)
	originPrivateKey[0] &= 248
	originPrivateKey[31] &= 127
	originPrivateKey[31] |= 64
	return &Account{publicKey: originPublicKey[:], privateKey: originPrivateKey[:]}
}

