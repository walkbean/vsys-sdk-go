package vsys

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataEncoder(t *testing.T) {
	data := DataEncoder{}
	var a int64 = 10
	data.EncodeArgAmount(2)
	data.Encode("AUB4EkWUBkZ22YzSxmTRRqvPUMrJ1AHg6P1", DeTypeAddress)
	data.Encode(a, DeTypeAmount)
	result := Base58Encode(data.Result())
	de := DataEncoder{}
	bytes := Base58Decode(result)
	list := de.Decode(bytes)
	assert.Equal(t, list[0].Type, int8(DeTypeAddress))
	assert.Equal(t, list[1].Type, int8(DeTypeAmount))
	assert.Equal(t, list[0].Value.(string), "AUB4EkWUBkZ22YzSxmTRRqvPUMrJ1AHg6P1")
	assert.Equal(t, list[1].Value.(int64), int64(10))
}

func TestContractBuild(t *testing.T) {
	// register
	c := Contract{}
	c.Max = 123
	c.Unity = 1000000
	c.TokenDescription = "hello"

	data := c.BuildRegisterData()
	fmt.Println("register", string(Base58Encode(data)))

	c = Contract{}
	c.Amount = 255

	data = c.BuildIssueData()
	fmt.Println("issue", string(Base58Encode(data)))

	c = Contract{}
	c.Recipient = "AUB4EkWUBkZ22YzSxmTRRqvPUMrJ1AHg6P1"
	c.Amount = 100

	data = c.BuildSendData()
	fmt.Println("send", string(Base58Encode(data)))

	c = Contract{}
	c.Amount = 15

	data = c.BuildDestroyData()
	fmt.Println("destroy", string(Base58Encode(data)))
}

func TestContractId2TokenId(t *testing.T) {
	assert.Equal(t, ContractId2TokenId(testContract, 0), testToken)
}

func TestTokenId2ContractId(t *testing.T) {
	assert.Equal(t, TokenId2ContractId(testToken), testContract)
}

func TestBase58Decode(t *testing.T) {
	assert.Equal(t, len(Base58Decode("DXiEzkh1H9LR19SGn3by7emNDVMAuTjmMyQXX6f1NdPR")), 32)
}

func TestContract_DecodeTexture(t *testing.T) {
	c := Contract{
		Textual: Textual{
			Triggers:    "124VnyFU9tQUn4Z19KBbV8aAQF4aCgWrQWrLL1yK5RpWY2sU74P8GU6wJ6dwyuFHP3Xt5Kmpm",
			Descriptors: "1RypGiL5eNbDESxn2SVM8HrLF6udvXV6YmwvFsp4fLJfRcr7nQuVFMvXn6KmWJeq8c53tdrutZcsQA2zyHb8Wj1tQUjGmitP6kLzcnpQXEq7AUZpMT6j7LCrhJvs3oLCCr7SSpz3h4iJJJg9WuL7Acbsw1x2AK4tRSZWXyrnLgqWhgqbTdfmxFGHjD58XrScBibJ9AUwEWCAeAna3NFofSZaSDxFJAK2adrrHhJdktQCQobMJMmC164HtJKF569naoMREkncYedQwXWk4uyPzGTUKsfXFwLaR77wv8gtNEjqwvGtpdFJELyJ3RC2F7exhqiiVxTaoGrAanuv1bianVbKqPAygPaGrhA1H3JmQWksNhg6q7dtPvBuqWDqDs4DkhV35JhNFeiER18o49pxX8zR1n1jvis6QrU2cD1Cn3yXwSZaW8TXWMKZ7ULRo1UcJykQvQCLq3EBVfzf6iULhuRagTnJ3Sq4tFSxgnNPhATLDreQpEe1BA3SfRWKRskLFjXV5aMeYxgFLfqYEFJ37BaRVyFZDSUgrKLMnNzrZZG2P81t7MhT6GpDApLZkNtjdGRMQGFsRN2azGruQReFnXeB3mScaxgfhGxcu9B",
		},
	}
	c.DecodeTexture()
	assert.Equal(t, c.Functions[0].Name, "supersede")
	assert.Equal(t, c.Functions[1].Name, "issue")
	assert.Equal(t, c.Functions[2].Name, "destroy")
	assert.Equal(t, c.Functions[3].Name, "split")
	assert.Equal(t, c.Functions[4].Name, "send")
	assert.Equal(t, c.Functions[5].Name, "transfer")
	assert.Equal(t, c.Functions[6].Name, "deposit")
	assert.Equal(t, c.Functions[7].Name, "withdraw")
	assert.Equal(t, c.Functions[8].Name, "totalSupply")
	assert.Equal(t, c.Functions[9].Name, "maxSupply")
	assert.Equal(t, c.Functions[10].Name, "balanceOf")
	assert.Equal(t, c.Functions[11].Name, "getIssuer")
}
