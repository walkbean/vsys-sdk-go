package vsys

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetTokenInfo(t *testing.T) {
	tests := []struct {
		TokenId string
	}{
		{TokenId: "TWscu6rbRF2PEsnY1bRky4aKxxKTzn69WMFLFdLxK"},
		{TokenId: "TWtSxBEx7rmsQ34MyWzwBCYYwRJh4K9xsL9zPkMK8"},
		{TokenId: "TWuyTczrVc4KeDUBpksxY8bpcogKfKqoVGE7cwcs3"},
	}

	for _, test := range tests {
		info, err := GetTokenInfo(test.TokenId)
		assert.NoError(t, err)
		assert.Equal(t, test.TokenId, info.TokenId)
	}
}

func TestMain(m *testing.M) {
	InitApi("http://test.v.systems:9922", Testnet)
	exitVal := m.Run()
	os.Exit(exitVal)
}
