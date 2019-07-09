# vsys-sdk-go
The golang library for V Systems Blockchain.

## Installing

Use `go get` to retrieve the SDK sto add it to your `GOPATH` workspace, or
project's Go module dependencies.

	go get github.com/walkbean/vsys-sdk-go
	
### Dependencies
The SDK includes a vendor folder containing the runtime dependencies of the SDK. The metadata of the SDK's dependencies can be found in the Go module file go.mod or Dep file Gopkg.toml.

## Usage

### Create Account 

#### Create Account From Seed

```go
	acc := InitAccount(TestnetByte)
	acc.BuildFromSeed("<SEED>", 0)
```

#### Create Account From PrivateKey
```go
	acc := InitAccount(TestnetByte)
	acc.BuildFromPrivateKey("<PRIVATE_KEY>")
```

### Transaction

#### Send Payment Transaction

