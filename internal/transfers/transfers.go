package transfers

import (
	"context"
	"fmt"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/signer"
	"github.com/block-vision/sui-go-sdk/sui"
)

func TransferSUI(ctx context.Context, suiClient sui.ISuiAPI, signerAccount *signer.Signer, recipient string, amount string) error {
	gasObject, err := GetGasObjectID(ctx, suiClient, signerAccount.Address)
	if err != nil {
		return fmt.Errorf("error getting gas object: %w", err)
	}

	rsp, err := suiClient.TransferSui(ctx, models.TransferSuiRequest{
		Signer:      signerAccount.Address,
		SuiObjectId: gasObject,
		GasBudget:   "10000000",
		Recipient:   recipient,
		Amount:      amount,
	})
	if err != nil {
		return fmt.Errorf("error preparing SUI transfer: %w", err)
	}

	txnRsp, err := suiClient.SignAndExecuteTransactionBlock(ctx, models.SignAndExecuteTransactionBlockRequest{
		TxnMetaData: rsp,
		PriKey:      signerAccount.PriKey,
		Options: models.SuiTransactionBlockOptions{
			ShowEffects: true,
		},
		RequestType: "WaitForLocalExecution",
	})
	if err != nil {
		return fmt.Errorf("error executing SUI transfer: %w", err)
	}

	fmt.Printf("Transaction executed successfully. Digest: %s\n", txnRsp.Digest)
	return nil
}

func GetGasObjectID(ctx context.Context, client sui.ISuiAPI, address string) (string, error) {
	coins, err := client.SuiXGetCoins(ctx, models.SuiXGetCoinsRequest{
		Owner:    address,
		CoinType: "0x2::sui::SUI",
	})
	if err != nil {
		return "", fmt.Errorf("error getting coins: %w", err)
	}
	if len(coins.Data) == 0 {
		return "", fmt.Errorf("no SUI coins found for address")
	}

	return coins.Data[0].CoinObjectId, nil
}
