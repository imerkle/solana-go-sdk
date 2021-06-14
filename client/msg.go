package client

import (
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/sysprog"
	"github.com/portto/solana-go-sdk/types"
)

func NewMessageWithNonce(feePayer common.PublicKey, instructions []types.Instruction, nonceAccountPubkey common.PublicKey, nonceAuthorityPubkey common.PublicKey) types.Message {
	ins := sysprog.AdvanceNonceAccount(nonceAccountPubkey, nonceAuthorityPubkey)
	instructions = append([]types.Instruction{ins}, instructions...)
	message := types.NewMessage(feePayer, instructions, "")
	return message
}
