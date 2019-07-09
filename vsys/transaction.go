package vsys

import (
	"time"
)

type Transaction struct {
	TxId            string `json:"txId,omitempty"`
	Timestamp       int64  `json:"timestamp"`
	Fee             int64  `json:"fee"`
	FeeScale        int16  `json:"feeScale"`
	Amount          int64  `json:"amount"`
	SenderPublicKey string `json:"senderPublicKey"`
	Attachment      []byte `json:"attachment,omitempty"`
	Recipient       string `json:"recipient,omitempty"`
	Signature       string `json:"signature"`
	txType          uint8
}

func NewPaymentTransaction(recipient string, amount int64, attachment []byte) *Transaction {
	return &Transaction{
		Timestamp:  time.Now().Unix() * 1000000000,
		Fee:        DefaultTxFee,
		FeeScale:   DefaultFeeScale,
		Recipient:  recipient,
		Amount:     amount,
		Attachment: attachment,
		txType:     TxTypePayment}
}

func NewLeaseTransaction(recipient string, amount int64) *Transaction {
	return &Transaction{
		Timestamp: time.Now().Unix() * 1000000000,
		Fee:       DefaultTxFee,
		FeeScale:  DefaultFeeScale,
		Recipient: recipient,
		Amount:    amount,
		txType:    TxTypeLease}
}

func NewCancelLeaseTransaction(txId string) *Transaction {
	return &Transaction{
		Timestamp: time.Now().Unix() * 1000000000,
		Fee:       DefaultTxFee,
		FeeScale:  DefaultFeeScale,
		TxId:      txId,
		txType:    TxTypeCancelLease}
}

func (tx *Transaction) TxType() int {
	return int(tx.txType)
}

// BuildTxData generate data which is used to be signed
func (tx *Transaction) BuildTxData() []byte {
	data := make([]byte, 0)
	data = append(data, tx.txType)
	if tx.Timestamp <= 0 {
		tx.Timestamp = time.Now().Unix() * 1e9
	}
	if tx.Fee <= 0 {
		tx.Fee = DefaultTxFee
	}
	if tx.FeeScale <= 0 {
		tx.FeeScale = DefaultFeeScale
	}
	switch tx.txType {
	case TxTypePayment:
		return tx.buildPaymentData(data)
	case TxTypeLease:
		return tx.buildLeaseData(data)
	case TxTypeCancelLease:
		return tx.buildLeaseCancelData(data)
	}
	return data
}

func (tx *Transaction) buildPaymentData(data []byte) []byte {
	data = append(data, uint64ToByte(tx.Timestamp)...)
	data = append(data, uint64ToByte(tx.Amount)...)
	data = append(data, uint64ToByte(tx.Fee)...)
	data = append(data, uint16ToByte(tx.FeeScale)...)
	data = append(data, Base58Decode(tx.Recipient)...)
	data = append(data, uint16ToByte(int16(len(tx.Attachment)))...)
	data = append(data, tx.Attachment...)
	return data
}

func (tx *Transaction) buildLeaseData(data []byte) []byte {
	data = append(data, Base58Decode(tx.Recipient)...)
	data = append(data, uint64ToByte(tx.Amount)...)
	data = append(data, uint64ToByte(tx.Fee)...)
	data = append(data, uint16ToByte(tx.FeeScale)...)
	data = append(data, uint64ToByte(tx.Timestamp)...)
	return data
}

func (tx *Transaction) buildLeaseCancelData(data []byte) []byte {
	data = append(data, uint64ToByte(tx.Fee)...)
	data = append(data, uint16ToByte(tx.FeeScale)...)
	data = append(data, uint64ToByte(tx.Timestamp)...)
	data = append(data, Base58Decode(tx.TxId)...)
	return data
}
