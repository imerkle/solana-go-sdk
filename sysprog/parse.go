package sysprog

import (
	"github.com/ghostiam/binstruct"
	"github.com/portto/solana-go-sdk/types"
)

func ParseSystem(ins types.Instruction) (types.ParsedInstruction, error) {
	var parsedInstruction types.ParsedInstruction
	var err error
	var s struct {
		Instruction Instruction
	}
	err = binstruct.UnmarshalLE(ins.Data, &s)
	var instructionType string
	var parsedInfo map[string]interface{}
	switch s.Instruction {
	case InstructionCreateAccount:
		var a CreateAccountInstruction
		err = binstruct.UnmarshalLE(ins.Data, &a)
		instructionType = "createAccount"
		parsedInfo = map[string]interface{}{
			"source":     ins.Accounts[0].PubKey.ToBase58(),
			"newAccount": ins.Accounts[1].PubKey.ToBase58(),
			"lamports":   a.Lamports,
			"space":      a.Space,
			"owner":      a.Owner.ToBase58(),
		}
		break
	case InstructionAssign:
		var a AssignInstruction
		err = binstruct.UnmarshalLE(ins.Data, &a)
		instructionType = "assign"
		parsedInfo = map[string]interface{}{
			"account": ins.Accounts[0].PubKey.ToBase58(),
			"owner":   a.AssignToProgramID.ToBase58(),
		}
		break
	case InstructionTransfer:
		var a TransferInstruction
		err = binstruct.UnmarshalLE(ins.Data, &a)
		instructionType = "transfer"
		parsedInfo = map[string]interface{}{
			"source":      ins.Accounts[0].PubKey.ToBase58(),
			"destination": ins.Accounts[1].PubKey.ToBase58(),
			"lamports":    a.Lamports,
		}
		parsedInstruction.Parsed = &types.InstructionInfo{}
		break
	case InstructionCreateAccountWithSeed:
		var a CreateAccountWithSeedInstruction
		err = binstruct.UnmarshalLE(ins.Data, &a)
		instructionType = "createAccountWithSeed"
		parsedInfo = map[string]interface{}{
			"source":     ins.Accounts[0].PubKey.ToBase58(),
			"newAccount": ins.Accounts[1].PubKey.ToBase58(),
			"base":       a.Base,
			"seed":       a.Seed,
			"space":      a.Space,
			"lamports":   a.Lamports,
			"owner":      a.ProgramID,
		}
		break
	case InstructionAdvanceNonceAccount:
		var a AdvanceNonceAccountInstruction
		err = binstruct.UnmarshalLE(ins.Data, &a)
		instructionType = "advanceNonce"
		parsedInfo = map[string]interface{}{
			"nonceAccount":            ins.Accounts[0].PubKey.ToBase58(),
			"recentBlockhashesSysvar": ins.Accounts[1].PubKey.ToBase58(),
			"nonceAuthority":          ins.Accounts[2].PubKey.ToBase58(),
		}
		break
	case InstructionWithdrawNonceAccount:
		var a WithdrawNonceAccountInstruction
		err = binstruct.UnmarshalLE(ins.Data, &a)
		instructionType = "withdrawFromNonce"
		parsedInfo = map[string]interface{}{
			"nonceAccount":            ins.Accounts[0].PubKey.ToBase58(),
			"destination":             ins.Accounts[1].PubKey.ToBase58(),
			"recentBlockhashesSysvar": ins.Accounts[2].PubKey.ToBase58(),
			"rentSysvar":              ins.Accounts[3].PubKey.ToBase58(),
			"nonceAuthority":          ins.Accounts[4].PubKey.ToBase58(),
			"lamports":                a.Lamports,
		}
		break
	case InstructionInitializeNonceAccount:
		var a InitializeNonceAccountInstruction
		err = binstruct.UnmarshalLE(ins.Data, &a)
		instructionType = "initializeNonce"
		parsedInfo = map[string]interface{}{
			"nonceAccount":            ins.Accounts[0].PubKey.ToBase58(),
			"recentBlockhashesSysvar": ins.Accounts[1].PubKey.ToBase58(),
			"rentSysvar":              ins.Accounts[2].PubKey.ToBase58(),
			"nonceAuthority":          a.Auth.ToBase58(),
		}
		break
	case InstructionAuthorizeNonceAccount:
		var a AuthorizeNonceAccountInstruction
		err = binstruct.UnmarshalLE(ins.Data, &a)
		instructionType = "authorizeNonce"
		parsedInfo = map[string]interface{}{
			"nonceAccount":   ins.Accounts[0].PubKey.ToBase58(),
			"nonceAuthority": ins.Accounts[1].PubKey.ToBase58(),
			"rentSysvar":     ins.Accounts[2].PubKey.ToBase58(),
			"newAuthorized":  a.Auth.ToBase58(),
		}
		break
	case InstructionAllocate:
		var a AllocateInstruction
		err = binstruct.UnmarshalLE(ins.Data, &a)
		instructionType = "allocate"
		parsedInfo = map[string]interface{}{
			"account": ins.Accounts[0].PubKey.ToBase58(),
			"space":   a.Space,
		}
		break
	case InstructionAllocateWithSeed:
		var a AllocateWithSeedInstruction
		err = binstruct.UnmarshalLE(ins.Data, &a)
		instructionType = "allocateWithSeed"
		parsedInfo = map[string]interface{}{
			"account": ins.Accounts[0].PubKey.ToBase58(),
			"base":    a.Base,
			"seed":    a.Seed,
			"space":   a.Space,
			"owner":   a.ProgramID.ToBase58(),
		}
		break
	case InstructionAssignWithSeed:
		var a AssignWithSeedInstruction
		err = binstruct.UnmarshalLE(ins.Data, &a)
		instructionType = "assignWithSeed"
		parsedInfo = map[string]interface{}{
			"account": ins.Accounts[0].PubKey.ToBase58(),
			"base":    a.Base,
			"seed":    a.Seed,
			"owner":   a.AssignToProgramID.ToBase58(),
		}
		break
	case InstructionTransferWithSeed:
		var a TransferWithSeedInstruction
		err = binstruct.UnmarshalLE(ins.Data, &a)
		instructionType = "transferWithSeed"
		parsedInfo = map[string]interface{}{
			"source":      ins.Accounts[0].PubKey.ToBase58(),
			"sourceBase":  ins.Accounts[1].PubKey.ToBase58(),
			"destination": ins.Accounts[2].PubKey.ToBase58(),
			"lamports":    a.Lamports,
			"sourceSeed":  a.Seed,
			"sourceOwner": a.ProgramID.ToBase58(),
		}
		break
	}
	parsedInstruction.Parsed = &types.InstructionInfo{
		Info:            parsedInfo,
		InstructionType: instructionType,
	}
	return parsedInstruction, err
}
