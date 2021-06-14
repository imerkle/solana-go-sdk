package client

type FeeCalculator struct {
	LamportsPerSignature uint64 `json:"lamportsPerSignature"`
}

type Instruction struct {
	ProgramIDIndex uint64   `json:"programIdIndex"`
	Accounts       []uint64 `json:"accounts"`
	Data           string   `json:"data"`
}

type TransactionMeta struct {
	Fee               uint64   `json:"fee"`
	PreBalances       []int64  `json:"preBalances"`
	PostBalances      []int64  `json:"postBalances"`
	LogMessages       []string `json:"logMesssages"`
	InnerInstructions []struct {
		Index        uint64        `json:"index"`
		Instructions []Instruction `json:"instructions"`
	} `json:"innerInstructions"`
	Err    interface{}            `json:"err"`
	Status map[string]interface{} `json:"status"`
}

type MessageHeader struct {
	NumRequiredSignatures       uint8 `json:"numRequiredSignatures"`
	NumReadonlySignedAccounts   uint8 `json:"numReadonlySignedAccounts"`
	NumReadonlyUnsignedAccounts uint8 `json:"numReadonlyUnsignedAccounts"`
}

type Message struct {
	Header          MessageHeader `json:"header"`
	AccountKeys     []string      `json:"accountKeys"`
	RecentBlockhash string        `json:"recentBlockhash"`
	Instructions    []Instruction `json:"instructions"`
}

type Transaction struct {
	Signatures []string `json:"signatures"`
	Message    Message  `json:"message"`
}

type Encoding string

const (
	EncodingBase58     Encoding = "base58" // limited to Account data of less than 128 bytes
	EncodingBase64     Encoding = "base64"
	EncodingBase64Zstd Encoding = "base64+zstd"
	EncodingJson       Encoding = "json"
)

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
