package tokenprog

import (
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

type Instruction uint8

const (
	InstructionInitializeMint Instruction = iota
	InstructionInitializeAccount
	InstructionInitializeMultisig
	InstructionTransfer
	InstructionApprove
	InstructionRevoke
	InstructionSetAuthority
	InstructionMintTo
	InstructionBurn
	InstructionCloseAccount
	InstructionFreezeAccount
	InstructionThawAccount
	InstructionTransferChecked
	InstructionApproveChecked
	InstructionMintToChecked
	InstructionBurnChecked
	InstructionInitializeAccount2
)

type InitializeMintInstruction struct {
	Instruction     Instruction
	Decimals        uint8
	MintAuthority   common.PublicKey
	Option          bool
	FreezeAuthority common.PublicKey
}
type InitializeAccountInstruction struct {
	Instruction Instruction
}
type InitializeMultisigInstruction struct {
	Instruction     Instruction
	MinimumRequired uint8
}
type TransferInstruction struct {
	Instruction Instruction
	Amount      uint64
}
type ApproveInstruction struct {
	Instruction Instruction
	Amount      uint64
}
type RevokeInstruction struct {
	Instruction Instruction
}
type MintToInstruction struct {
	Instruction Instruction
	Amount      uint64
}
type BurnInstruction struct {
	Instruction Instruction
	Amount      uint64
}
type CloseAccountInstruction struct {
	Instruction Instruction
}
type FreezeAccountInstruction struct {
	Instruction Instruction
}
type ThawAccountInstruction struct {
	Instruction Instruction
}
type TransferCheckedInstruction struct {
	Instruction Instruction
	Amount      uint64
	Decimals    uint8
}
type ApproveCheckedInstruction struct {
	Instruction Instruction
	Amount      uint64
	Decimals    uint8
}
type MintToCheckedInstruction struct {
	Instruction Instruction
	Amount      uint64
	Decimals    uint8
}
type BurnCheckedInstruction struct {
	Instruction Instruction
	Amount      uint64
	Decimals    uint8
}
type InitializeAccount2Instruction struct {
	Instruction Instruction
	Owner       common.PublicKey
}

// InitializeMint init a mint, if you don't need to freeze, pass the empty pubKey common.PublicKey{}
func InitializeMint(decimals uint8, mint, mintAuthority common.PublicKey, freezeAuthority common.PublicKey) types.Instruction {
	data, err := common.SerializeData(InitializeMintInstruction{
		Instruction:     InstructionInitializeMint,
		Decimals:        decimals,
		MintAuthority:   mintAuthority,
		Option:          freezeAuthority != common.PublicKey{},
		FreezeAuthority: freezeAuthority,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: mint, IsSigner: false, IsWritable: true},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		Data: data,
	}
}

// InitializeAccount init a token account which can receive token
func InitializeAccount(accountPublicKey, mintPublicKey, ownerPublickey common.PublicKey) types.Instruction {
	data, err := common.SerializeData(InitializeAccountInstruction{
		Instruction: InstructionInitializeAccount,
	})
	if err != nil {
		panic(err)
	}

	accounts := []types.AccountMeta{
		{PubKey: accountPublicKey, IsSigner: false, IsWritable: true},
		{PubKey: mintPublicKey, IsSigner: false, IsWritable: false},
		{PubKey: ownerPublickey, IsSigner: false, IsWritable: false},
		{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
	}
	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

func InitializeMultisig(authPubkey common.PublicKey, signerPubkeys []common.PublicKey, miniRequired uint8) types.Instruction {
	if len(signerPubkeys) < 1 {
		panic("minimum of signer is 1")
	}
	if len(signerPubkeys) > 11 {
		panic("maximum of signer is 11")
	}
	if miniRequired > uint8(len(signerPubkeys)) {
		panic("required number too big")
	}

	data, err := common.SerializeData(InitializeMultisigInstruction{
		Instruction:     InstructionInitializeMultisig,
		MinimumRequired: miniRequired,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 2+len(signerPubkeys))
	accounts = append(accounts,
		types.AccountMeta{PubKey: authPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
	)
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

func Transfer(srcPubkey, destPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64) types.Instruction {
	data, err := common.SerializeData(TransferInstruction{
		Instruction: InstructionTransfer,
		Amount:      amount,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts, types.AccountMeta{PubKey: srcPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: destPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false})
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}
	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

func Approve(sourcePubkey, delegatePubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64) types.Instruction {
	data, err := common.SerializeData(ApproveInstruction{
		Instruction: InstructionApprove,
		Amount:      amount,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts, types.AccountMeta{PubKey: sourcePubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: delegatePubkey, IsSigner: false, IsWritable: false})
	accounts = append(accounts, types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false})
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

func Revoke(srcPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey) types.Instruction {
	data, err := common.SerializeData(RevokeInstruction{
		Instruction: InstructionRevoke,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 2+len(signerPubkeys))
	accounts = append(accounts,
		types.AccountMeta{PubKey: srcPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false},
	)
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

func SetAuthority() types.Instruction {
	panic("not implement yet")
}

func MintTo(mintPubkey, destPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64) types.Instruction {
	data, err := common.SerializeData(MintToInstruction{
		Instruction: InstructionMintTo,
		Amount:      amount,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts,
		types.AccountMeta{PubKey: mintPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: destPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false},
	)
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

func Burn(accountPubkey, mintPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64) types.Instruction {
	data, err := common.SerializeData(BurnInstruction{
		Instruction: InstructionBurn,
		Amount:      amount,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts,
		types.AccountMeta{PubKey: accountPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: mintPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false},
	)
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

// Close an account and transfer its all SOL to dest, only account's token balance is zero can be closed.
func CloseAccount(accountPubkey, destPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey) types.Instruction {
	data, err := common.SerializeData(CloseAccountInstruction{
		Instruction: InstructionCloseAccount,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts, types.AccountMeta{PubKey: accountPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: destPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false})
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

func FreezeAccount(accountPubkey, mintPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey) types.Instruction {
	data, err := common.SerializeData(FreezeAccountInstruction{
		Instruction: InstructionFreezeAccount,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts, types.AccountMeta{PubKey: accountPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: mintPubkey, IsSigner: false, IsWritable: false})
	accounts = append(accounts, types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false})
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

func ThawAccount(accountPubkey, mintPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey) types.Instruction {
	data, err := common.SerializeData(ThawAccountInstruction{
		Instruction: InstructionThawAccount,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts, types.AccountMeta{PubKey: accountPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: mintPubkey, IsSigner: false, IsWritable: false})
	accounts = append(accounts, types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false})
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

func TransferChecked(srcPubkey, destPubkey, mintPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64, decimals uint8) types.Instruction {
	data, err := common.SerializeData(TransferCheckedInstruction{
		Instruction: InstructionTransferChecked,
		Amount:      amount,
		Decimals:    decimals,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 4+len(signerPubkeys))
	accounts = append(accounts, types.AccountMeta{PubKey: srcPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: mintPubkey, IsSigner: false, IsWritable: false})
	accounts = append(accounts, types.AccountMeta{PubKey: destPubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false})
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

func ApproveChecked(sourcePubkey, mintPubkey, delegatePubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64, decimals uint8) types.Instruction {
	data, err := common.SerializeData(ApproveCheckedInstruction{
		Instruction: InstructionApproveChecked,
		Amount:      amount,
		Decimals:    decimals,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 4+len(signerPubkeys))
	accounts = append(accounts, types.AccountMeta{PubKey: sourcePubkey, IsSigner: false, IsWritable: true})
	accounts = append(accounts, types.AccountMeta{PubKey: mintPubkey, IsSigner: false, IsWritable: false})
	accounts = append(accounts, types.AccountMeta{PubKey: delegatePubkey, IsSigner: false, IsWritable: false})
	accounts = append(accounts, types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false})
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

func MintToChecked(mintPubkey, destPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64, decimals uint8) types.Instruction {
	data, err := common.SerializeData(MintToCheckedInstruction{
		Instruction: InstructionMintToChecked,
		Amount:      amount,
		Decimals:    decimals,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts,
		types.AccountMeta{PubKey: mintPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: destPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false},
	)
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

func BurnChecked(accountPubkey, mintPubkey, authPubkey common.PublicKey, signerPubkeys []common.PublicKey, amount uint64, decimals uint8) types.Instruction {
	data, err := common.SerializeData(BurnCheckedInstruction{
		Instruction: InstructionBurnChecked,
		Amount:      amount,
		Decimals:    decimals,
	})
	if err != nil {
		panic(err)
	}

	accounts := make([]types.AccountMeta, 0, 3+len(signerPubkeys))
	accounts = append(accounts,
		types.AccountMeta{PubKey: accountPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: mintPubkey, IsSigner: false, IsWritable: true},
		types.AccountMeta{PubKey: authPubkey, IsSigner: len(signerPubkeys) == 0, IsWritable: false},
	)
	for _, signerPubkey := range signerPubkeys {
		accounts = append(accounts, types.AccountMeta{PubKey: signerPubkey, IsSigner: true, IsWritable: false})
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts:  accounts,
		Data:      data,
	}
}

func InitializeAccount2(accountPubkey, mintPubkey, ownerPubkey common.PublicKey) types.Instruction {
	data, err := common.SerializeData(InitializeAccount2Instruction{
		Instruction: InstructionInitializeAccount2,
		Owner:       ownerPubkey,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.TokenProgramID,
		Accounts: []types.AccountMeta{
			{PubKey: accountPubkey, IsSigner: false, IsWritable: true},
			{PubKey: mintPubkey, IsSigner: false, IsWritable: false},
			{PubKey: common.SysVarRentPubkey, IsSigner: false, IsWritable: false},
		},
		Data: data,
	}
}
