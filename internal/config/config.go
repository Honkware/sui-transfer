package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PrivateKey       string
	RecipientAddress string
	AmountToTransfer string
	SuiNetwork       string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	mnemonic := os.Getenv("MNEMONIC")
	recipientAddress := os.Getenv("RECIPIENT_ADDRESS")
	amountToTransfer := os.Getenv("AMOUNT_TO_TRANSFER")
	suiNetwork := os.Getenv("SUI_NETWORK")

	if mnemonic == "" || recipientAddress == "" || amountToTransfer == "" || suiNetwork == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return &Config{
		PrivateKey:       mnemonic,
		RecipientAddress: recipientAddress,
		AmountToTransfer: amountToTransfer,
		SuiNetwork:       suiNetwork,
	}, nil
}
