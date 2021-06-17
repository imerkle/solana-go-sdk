package client

import (
	"context"

	"github.com/portto/solana-go-sdk/types"
)

type Reward struct {
	Pubkey      string `json:"pubkey"`
	Lamports    int64  `json:"lamports"`
	PostBalance uint64 `json:"postBalance"`
	RewardType  string `json:"rewardType"` // type of reward: "fee", "rent", "voting", "staking"
}
type GetConfirmBlockResponse struct {
	Blockhash         string                `json:"blockhash"`
	PreviousBlockhash string                `json:"previousBlockhash"`
	ParentSlot        uint64                `json:"parentSlot"`
	BlockTime         int64                 `json:"blockTime"`
	Transactions      []TransactionWithMeta `json:"transactions"`
	Rewards           []Reward              `json:"rewards"`
}
type TransactionWithMeta struct {
	Meta        TransactionMeta `json:"meta"`
	Transaction Transaction     `json:"transaction"`
}
type ParsedTransactionWithMeta struct {
	Meta        TransactionMeta         `json:"meta"`
	Transaction types.ParsedTransaction `json:"transaction"`
}
type GetConfirmBlockParsedResponse struct {
	Blockhash         string                      `json:"blockhash"`
	PreviousBlockhash string                      `json:"previousBlockhash"`
	ParentSlot        uint64                      `json:"parentSlot"`
	BlockTime         int64                       `json:"blockTime"`
	Transactions      []ParsedTransactionWithMeta `json:"transactions"`
	Rewards           []Reward                    `json:"rewards"`
}

func (s *Client) GetConfirmedBlock(ctx context.Context, slot uint64) (GetConfirmBlockResponse, error) {
	res := struct {
		GeneralResponse
		Result GetConfirmBlockResponse `json:"result"`
	}{}
	err := s.request(ctx, "getConfirmedBlock", []interface{}{slot, "json"}, &res)
	if err != nil {
		return GetConfirmBlockResponse{}, err
	}
	return res.Result, nil
}
func (s *Client) GetConfirmedBlockParsed(ctx context.Context, slot uint64) (GetConfirmBlockParsedResponse, error) {
	res := struct {
		GeneralResponse
		Result GetConfirmBlockParsedResponse `json:"result"`
	}{}
	err := s.request(ctx, "getConfirmedBlock", []interface{}{slot, "jsonParsed"}, &res)
	if err != nil {
		return GetConfirmBlockParsedResponse{}, err
	}
	return res.Result, nil
}
