package types

import "github.com/portto/solana-go-sdk/common"

type CompiledInstruction struct {
	ProgramIDIndex int
	Accounts       []int
	Data           []byte
}

type Instruction struct {
	ProgramID common.PublicKey
	Accounts  []AccountMeta
	Data      []byte
}

type ParsedTransaction struct {
	Signatures []string      `json:"signatures"`
	Message    ParsedMessage `json:"message"`
}
type ParsedMessage struct {
	Header          MessageHeader       `json:"header"`
	AccountKeys     []string            `json:"accountKeys"`
	RecentBlockhash string              `json:"recentBlockhash"`
	Instructions    []ParsedInstruction `json:"instructions"`
}
type ParsedInstruction struct {
	Accounts  []string         `json:"accounts,omitempty"`
	Data      string           `json:"data,omitempty"`
	Parsed    *InstructionInfo `json:"parsed,omitempty"`
	Program   string           `json:"program,omitempty"`
	ProgramID string           `json:"programId"`
}

type InstructionInfo struct {
	Info            map[string]interface{} `json:"info"`
	InstructionType string                 `json:"type"`
}

func GetUniqueSigners(ins []Instruction) []string {
	var signers []string
	var signersMap map[string]bool = make(map[string]bool)
	for _, v := range ins {
		for _, v1 := range v.Accounts {
			address := v1.PubKey.ToBase58()
			if v1.IsSigner {
				if _, ok := signersMap[address]; !ok {
					signersMap[address] = true
					signers = append(signers, address)
				}
			}
		}
	}
	return signers
}
