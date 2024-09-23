package signer

import (
	"fmt"
	"os"

	"github.com/block-vision/sui-go-sdk/signer"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func NewSignerWithMnemonic() (*signer.Signer, error) {
	mnemonic := os.Getenv("MNEMONIC")
	if mnemonic == "" {
		return nil, fmt.Errorf("MNEMONIC not set in environment")
	}

	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, fmt.Errorf("invalid mnemonic")
	}

	seed := bip39.NewSeed(mnemonic, "")

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, fmt.Errorf("failed to generate master key: %v", err)
	}

	childKey, err := masterKey.NewChildKey(bip32.FirstHardenedChild + 44) // m/44'
	if err != nil {
		return nil, fmt.Errorf("failed to derive child key: %v", err)
	}
	childKey, err = childKey.NewChildKey(bip32.FirstHardenedChild + 784) // m/44'/784'
	if err != nil {
		return nil, fmt.Errorf("failed to derive child key: %v", err)
	}
	childKey, err = childKey.NewChildKey(bip32.FirstHardenedChild + 0) // m/44'/784'/0'
	if err != nil {
		return nil, fmt.Errorf("failed to derive child key: %v", err)
	}
	childKey, err = childKey.NewChildKey(0) // m/44'/784'/0'/0
	if err != nil {
		return nil, fmt.Errorf("failed to derive child key: %v", err)
	}

	return signer.NewSigner(childKey.Key), nil
}
