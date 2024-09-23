package main

import (
	"context"
	"sui-transfer/internal/client"
	"sui-transfer/internal/config"
	"sui-transfer/internal/signer"
	"sui-transfer/internal/transfers"
	"sui-transfer/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {
	log := logger.GetLogger()

	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Error("Error loading .env file")
		return
	}

	cfg, err := config.Load()
	if err != nil {
		log.WithError(err).Error("Failed to load configuration")
		return
	}
	log.Info("Config loaded successfully")

	ctx := context.Background()

	suiClient, err := client.NewSuiClient(cfg.SuiNetwork)
	if err != nil {
		log.WithError(err).Error("Failed to create SUI client")
		return
	}
	log.Infof("SUI client created for network: %s", cfg.SuiNetwork)

	signerAccount, err := signer.NewSignerWithMnemonic()
	if err != nil {
		log.WithError(err).Error("Failed to create signer")
		return
	}
	log.Infof("Signer created: %s", signerAccount.Address)

	err = transfers.TransferSUI(ctx, suiClient, signerAccount, cfg.RecipientAddress, cfg.AmountToTransfer)
	if err != nil {
		log.WithError(err).Error("Failed to transfer SUI")
	} else {
		log.Info("SUI transfer completed successfully")
	}
}
