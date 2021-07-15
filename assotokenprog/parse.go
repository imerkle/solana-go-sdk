package assotokenprog

import "github.com/portto/solana-go-sdk/types"

func ParseAssocToken(ins types.Instruction) (types.ParsedInstruction, error) {
	var parsedInstruction types.ParsedInstruction
	var err error
	instructionType := "create"
	parsedInfo := map[string]interface{}{
		"source":        ins.Accounts[0].PubKey.ToBase58(),
		"account":       ins.Accounts[1].PubKey.ToBase58(),
		"wallet":        ins.Accounts[2].PubKey.ToBase58(),
		"mint":          ins.Accounts[3].PubKey.ToBase58(),
		"systemProgram": ins.Accounts[4].PubKey.ToBase58(),
		"tokenProgram":  ins.Accounts[5].PubKey.ToBase58(),
		"rentSysvar":    ins.Accounts[6].PubKey.ToBase58(),
	}
	parsedInstruction.Parsed = &types.InstructionInfo{
		Info:            parsedInfo,
		InstructionType: instructionType,
	}
	return parsedInstruction, err
}
