package common

var (
	SystemProgramID                    = PublicKeyFromString("11111111111111111111111111111111")
	ConfigProgramID                    = PublicKeyFromString("Config1111111111111111111111111111111111111")
	StakeProgramID                     = PublicKeyFromString("Stake11111111111111111111111111111111111111")
	VoteProgramID                      = PublicKeyFromString("Vote111111111111111111111111111111111111111")
	BPFLoaderProgramID                 = PublicKeyFromString("BPFLoader1111111111111111111111111111111111")
	Secp256k1ProgramID                 = PublicKeyFromString("KeccakSecp256k11111111111111111111111111111")
	TokenProgramID                     = PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	SPLAssociatedTokenAccountProgramID = PublicKeyFromString("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL")
)

func GetProgramName(programId PublicKey) string {
	name := "Unknown"
	switch programId {
	case SystemProgramID:
		name = "system"
		break
	case ConfigProgramID:
		name = "spl-token"
		break
	case StakeProgramID:
		name = "stake"
		break
	case VoteProgramID:
		name = "vote"
		break
	case BPFLoaderProgramID:
		name = "bpf-loader"
		break
	case Secp256k1ProgramID:
		name = "secp256k1"
		break
	case TokenProgramID:
		name = "spl-token"
		break
	case SPLAssociatedTokenAccountProgramID:
		name = "spl-associated-token-account"
		break
	}
	return name
}
