// Package sdk is golang sdk for v systems blockchain
// example code for vsys-sdk-go
// func Usage() {
//
//	// Get Account
//	acc := vsys.InitAccount(vsys.Testnet)
//	acc.BuildFromSeed("<SEED>", 0)
//
//	// Get Address
//	acc.Address()
//	// Get PublicKey
//	acc.PublicKey()
//
//	acc := vsys.InitAccount(vsys.Testnet)
//	acc.BuildFromPrivateKey("<PRIVATE_KEY>")
//
//	acc.PrivateKey()
//
//
//	// Init Api
//	vsys.InitApi("https://wallet.v.systems/api", vsys.Mainnet)
//	// Create Payment Transaction (send 1 vsys, attachment empty)
//	tx := acc.BuildPayment("<RECIPIENT_ADDRESS>", 1e8, []byte{})
//	vsys.SendPaymentTx(tx)
//
//	// Create Lease Transaction (lease 1 vsys)
//	tx = acc.BuildLeasing("<RECIPIENT_ADDRESS>", 1e8)
//	vsys.SendLeasingTx(tx)
//
//	// Create Cancel Lease Transaction
//	tx = acc.BuildCancelLeasing("<TRANSACTION_ID>")
//	vsys.SendCancelLeasingTx(tx)
//}

package sdk


