package client

import (
	"context"

	"github.com/portto/solana-go-sdk/types"
)

type GetConfirmedTransactionResponse struct {
	Slot        uint64          `json:"slot"`
	Meta        TransactionMeta `json:"meta"`
	Transaction Transaction     `json:"transaction"`
}

type GetConfirmedTransactionParsedResponse struct {
	Slot        uint64                  `json:"slot"`
	Meta        TransactionMeta         `json:"meta"`
	Transaction types.ParsedTransaction `json:"transaction"`
}

func (s *Client) GetConfirmedTransaction(ctx context.Context, txhash string) (GetConfirmedTransactionResponse, error) {
	res := struct {
		GeneralResponse
		Result GetConfirmedTransactionResponse `json:"result"`
	}{}
	err := s.request(ctx, "getConfirmedTransaction", []interface{}{txhash, "json"}, &res)
	if err != nil {
		return GetConfirmedTransactionResponse{}, err
	}
	return res.Result, nil
}

func (s *Client) GetConfirmedTransactionParsed(ctx context.Context, txhash string) (GetConfirmedTransactionParsedResponse, error) {
	res := struct {
		GeneralResponse
		Result GetConfirmedTransactionParsedResponse `json:"result"`
	}{}
	err := s.request(ctx, "getConfirmedTransaction", []interface{}{txhash, "jsonParsed"}, &res)
	if err != nil {
		return GetConfirmedTransactionParsedResponse{}, err
	}
	return res.Result, nil
}
