package vsys

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

const (
	ApiNodeVersion      = "/node/version"
	ApiBroadcastPayment = "/vsys/broadcast/payment"
	ApiBlockHeight      = "/blocks/height"
	ApiGetBlockAt       = "/blocks/at/%d"
	ApiGetBlockBySig    = "/blocks/signature/%s"
	ApiBlocks           = "/blocks/seq/%d/%d"

	// leasing
	ApiBroadcastLease       = "/leasing/broadcast/lease"
	ApiBroadcastCancelLease = "/leasing/broadcast/cancel"

	// transactions
	ApiUnConfirmedTransaction  = "/transactions/unconfirmed"
	ApiGetTransactionInfo      = "/transactions/info/%s"
	ApiGetTransactionByAddress = "/transactions/address/%s/limit/%d"
	ApiGetTransactionList      = "/transactions/list"

	// peers
	ApiGetPeersConnected = "/peers/connected"

	// consensus
	ApiGetAllSlotsInfo  = "/consensus/allSlotsInfo"
	ApiGetConsensusAlgo = "/consensus/algo"

	// address
	ApiGetAddressDetail  = "/addresses/balance/details/%s"
	ApiGetAddressValid   = "/addresses/validate/%s"
	ApiGetAddressBalance = "/addresses/balance/%s"

	//contract
	ApiContractBroadcastRegister = "/contract/broadcast/register"
	ApiContractBroadcastExecute  = "/contract/broadcast/execute"
	ApiContractInfo              = "/contract/info/%s"
	ApiTokenInfo                 = "/contract/tokenInfo/%s"
	ApiContractTokenBalance      = "/contract/balance/%s/%s" // /contract/balance/{address}/{tokenId}
)

type VsysApi struct {
	NodeAddress string
	Network     NetType
}

type TransactionResponse struct {
	Type       uint8   `json:"type"`
	Id         string  `json:"id"`
	Fee        int64   `json:"fee"`
	Timestamp  int64   `json:"timestamp"`
	Proofs     []Proof `json:"proofs"`
	Recipient  string  `json:"recipient"`
	FeeScale   int16   `json:"feeScale"`
	Amount     int64   `json:"amount"`
	Attachment string  `json:"attachment"`
	CommonResp
}

type Proof struct {
	ProofType string `json:"proofType"`
	PublicKey string `json:"publicKey"`
	Signature string `json:"signature"`
}

type CommonResp struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

func (a *VsysApi) Post(path string, data interface{}) ([]byte, error) {
	return Post(fmt.Sprintf("%s%s", a.NodeAddress, path), data)
}

func (a *VsysApi) Get(path string) ([]byte, error) {
	return UrlGetContent(fmt.Sprintf("%s%s", a.NodeAddress, path))
}

var api *VsysApi

func InitApi(nodeAddress string, network NetType) {
	api = &VsysApi{
		NodeAddress: nodeAddress,
		Network:     network,
	}
}

func GetVsysApi() *VsysApi {
	return api
}

func postSendTx(path string, tx *Transaction) (resp TransactionResponse, err error) {
	data, err := GetVsysApi().Post(path, tx)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

func SendPaymentTx(tx *Transaction) (resp TransactionResponse, err error) {
	return postSendTx(ApiBroadcastPayment, tx)
}

func SendLeasingTx(tx *Transaction) (resp TransactionResponse, err error) {
	return postSendTx(ApiBroadcastLease, tx)
}

func SendCancelLeasingTx(tx *Transaction) (resp TransactionResponse, err error) {
	return postSendTx(ApiBroadcastCancelLease, tx)
}

func SendRegisterContractTx(tx *Transaction) (resp TransactionResponse, err error) {
	return postSendTx(ApiContractBroadcastRegister, tx)
}

func SendExecuteContractTx(tx *Transaction) (resp TransactionResponse, err error) {
	return postSendTx(ApiContractBroadcastExecute, tx)
}

// txType eg: TxTypePayment | TxTypeLeasing
// txType <= 0 will return all kind of transactions
func GetTransactionList(address string,limit int64, offset int64, txType int64) (TxHistoryList, error){
	params := url.Values{}
	params.Set("address", address)
	if txType > 0 {
		params.Set("txType", strconv.FormatInt(txType, 10))
	}
	params.Set("limit", strconv.FormatInt(limit, 10))
	params.Set("offset", strconv.FormatInt(offset, 10))
	path := fmt.Sprintf("%s?%s", ApiGetTransactionList, params.Encode())
	resp, err := GetVsysApi().Get(path)
	if err != nil {
		return TxHistoryList{}, err
	}
	var data TxHistoryList
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return TxHistoryList{}, err
	}
	return data, nil
}