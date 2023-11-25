package core

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/useloopso/transaction-relayer/config"
	"github.com/useloopso/transaction-relayer/contracts"
	"github.com/useloopso/transaction-relayer/types"
)

func ExecuteRelayCall(execPayload types.ExecutePayload) error {
	client, err := ethclient.Dial(config.Get().NodeUrl)
	if err != nil {
		return err
	}

	universalProfile, err := contracts.NewUniversalProfile(common.HexToAddress(execPayload.Address), client)
	if err != nil {
		return err
	}

	keyManagerAddr, err := universalProfile.Owner(nil)
	if err != nil {
		return err
	}

	keyManager, err := contracts.NewLSP6(keyManagerAddr, client)
	if err != nil {
		return err
	}

	auth, err := getAuth()
	if err != nil {
		return err
	}

	sig, err := hexutil.Decode(execPayload.Signature)
	if err != nil {
		return err
	}

	abiPayload, err := hexutil.Decode(execPayload.Abi)
	if err != nil {
		return err
	}

	nonce := big.NewInt(int64(execPayload.Nonce))

	// TODO: Replace w input from user - this timestamp means valid until 1 Jan 2025
	validityTstamp := big.NewInt(0)

	tx, err := keyManager.ExecuteRelayCall(auth, sig, nonce, validityTstamp, abiPayload)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return err
	}

	fmt.Println("Relay transaction executed: ", tx.Hash())

	return nil
}

func getAuth() (*bind.TransactOpts, error) {
	privKey, err := crypto.HexToECDSA(config.Get().Key)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(int64(4201)))
	if err != nil {
		return nil, err
	}

	return auth, nil
}
