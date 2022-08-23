package main

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/casper-ecosystem/casper-golang-sdk/keypair/secp256k1"
	"github.com/casper-ecosystem/casper-golang-sdk/sdk"
	"github.com/casper-ecosystem/casper-golang-sdk/serialization"
	"github.com/casper-ecosystem/casper-golang-sdk/types"
)

func Trial() {
	rpcClient := sdk.NewRpcClient("http://65.21.237.153:7777/rpc")
	privKeyPath := "/Users/srivallabh-prof/Downloads/test_account_2/secret_key.pem"
	pubKeyPath := "/Users/srivallabh-prof/Downloads/test_account_2/public_key.pem"

	pair := secp256k1.ParseKeyFiles(pubKeyPath, privKeyPath)

	var hash32 [32]byte

	// contract hash
	// hash-c1f0f08d9a3cfc022e5baa5d6cbc645cd4a725503ab1df9eb5cc5b356788cbf8
	decodedHash, err2 := hex.DecodeString("c1f0f08d9a3cfc022e5baa5d6cbc645cd4a725503ab1df9eb5cc5b356788cbf8")
	if err2 != nil {
		return
	}

	for i := 0; i < 32; i++ {
		hash32[i] = decodedHash[i]
	}

	argsOrder := append(make([]string, 0), "recipient", "token_id")

	// ==== args =====
	// public key
	recipient := "011542c5f1909889ac1f4937d9043c0f135fe229993f15780c45246a8d170617c7"
	// token id
	token_id, _ := serialization.Marshal("helloworld")
	// ===============

	args := sdk.NewRunTimeArgs(map[string]sdk.Value{
		"recipient": {
			Tag:         types.CLTypePublicKey,
			StringBytes: recipient,
		},
		"token_id": {
			Tag:         types.CLTypeString,
			StringBytes: hex.EncodeToString(token_id),
		},
	}, argsOrder)

	session := sdk.NewStoredContractByHash(hash32, "transfer_token", *args)

	payment := sdk.StandardPayment(big.NewInt(10000000000))

	deploy := sdk.MakeDeploy(sdk.NewDeployParams(pair.PublicKey(), "casper-net-1", nil, 0), payment, session)

	deploy.SignDeploy(pair)
	result, err := rpcClient.PutDeploy(*deploy)
	if err != nil {
		fmt.Printf("err is:  %+v\n", err)
	}

	fmt.Printf("hash is => \n%+v\n", result.Hash)

}

func main() {
	Trial()
}
