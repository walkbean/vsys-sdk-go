package vsys

import (
	"encoding/json"
	"fmt"
)

type TokenInfo struct {
	TokenId 		string `json:"tokenId"`
	ContractId 		string `json:"contractId"`
	Max 			int64 `json:"max"`
	Total  			int64 `json:"total"`
	Unity 			int64 `json:"unity"`
	Description  	string `json:"description"`
}

// Get token info by token_id
func GetTokenInfo(tokenId string) (TokenInfo, error) {
	path := fmt.Sprintf(ApiTokenInfo, tokenId)
	resp, err := GetVsysApi().Get(path)
	if err != nil {
		return TokenInfo{}, err
	}

	var info TokenInfo
	err = json.Unmarshal(resp, &info)
	if err != nil {
		return TokenInfo{}, err
	}

	return info, nil
}