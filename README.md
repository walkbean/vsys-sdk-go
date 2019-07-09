# vsys-sdk-go
The golang library for V Systems Blockchain.

## Installing

Use `go get` to retrieve the SDK sto add it to your `GOPATH` workspace, or
project's Go module dependencies.

	go get github.com/walkbean/vsys-sdk-go
	
### Dependencies
The SDK includes a vendor folder containing the runtime dependencies of the SDK. The metadata of the SDK's dependencies can be found in the GoVendor file vendor/vendor.json.

## Usage

### Account 

#### Create Account From Seed

```go
acc := vsys.InitAccount(TestnetByte)
acc.BuildFromSeed("<SEED>", 0)
```

#### Create Account From PrivateKey
```go
acc := vsys.InitAccount(TestnetByte)
acc.BuildFromPrivateKey("<PRIVATE_KEY>")
```

### Transaction

1. Init Api

```go
// For Mainnet
vsys.InitApi("https://wallet.v.systems/api", vsys.Mainnet)
// For TestNet
vsys.InitApi("http://test.v.systems:9922", vsys.Testnet)

```

2. Make Transaction
```go
// Create Payment Transaction (send 1 vsys, attachment empty)
tx := acc.BuildPayment("<RECIPIENT_ADDRESS>", 1e8, []byte{})
vsys.SendPaymentTx(tx)
	
// Create Lease Transaction (lease 1 vsys)
tx = acc.BuildLeasing("<RECIPIENT_ADDRESS>", 1e8)
vsys.SendLeasingTx(tx)
    
// Create Cancel Lease Transaction
tx = acc.BuildCancelLeasing("<TRANSACTION_ID>")
vsys.SendCancelLeasingTx(tx)
```



