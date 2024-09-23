package client

import (
	"fmt"

	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/sui"
)

func NewSuiClient(network string) (sui.ISuiAPI, error) {
	var endpoint string

	switch network {
	case "mainnet":
		endpoint = constant.BvMainnetEndpoint
	case "testnet":
		endpoint = constant.BvTestnetEndpoint
	default:
		return nil, fmt.Errorf("unsupported network: %s", network)
	}

	return sui.NewSuiClient(endpoint), nil
}
