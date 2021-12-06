package client

import (
	"context"
	"github.com/portto/solana-go-sdk/common"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAccountInfo_Parsed(t *testing.T) {
	type args struct {
		ctx        context.Context
		base58Addr string
		cfg        GetAccountInfoConfig
	}
	tests := []struct {
		name         string
		requestBody  string
		responseBody string
		args         args
		want         GetAccountInfoParsedResponse
		err          error
	}{
		{
			requestBody:  `{"id":0,"jsonrpc":"2.0","method":"getAccountInfo","params":["93d6j7EWFZrRSzmLXQxG5WXDEcNrG48iPuPMidEwvm4H",{"encoding":"jsonParsed"}]}`,
			responseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":106714417},"value":{"data":{"parsed":{"info":{"authority":"95kKDZPo7MYUneHGfcyYPo2xxxi35HcR8k9R8j6bGqve","blockhash":"4jTYsNkM3s2nvnW2Akt2Wx7GE8CWMGy6J9Z2ajZqUbkA","feeCalculator":{"lamportsPerSignature":"5000"}},"type":"initialized"},"program":"nonce","space":80},"executable":false,"lamports":100000000,"owner":"11111111111111111111111111111111","rentEpoch":259}}}`,
			args: args{
				ctx:        context.Background(),
				base58Addr: "93d6j7EWFZrRSzmLXQxG5WXDEcNrG48iPuPMidEwvm4H",
				cfg: GetAccountInfoConfig{
					Encoding: "jsonParsed",
				},
			},
			want: GetAccountInfoParsedResponse{
				RentEpoch: 259,
				Lamports:  100000000,
				Owner:     common.SystemProgramID.ToBase58(),
				Excutable: false,
				Data: AccountData{
					AccountParsed: AccountParsed{
						AccountInfo: AccountInfo{
							Authority: "95kKDZPo7MYUneHGfcyYPo2xxxi35HcR8k9R8j6bGqve",
							BlockHash: "4jTYsNkM3s2nvnW2Akt2Wx7GE8CWMGy6J9Z2ajZqUbkA",
							AccountFeeCalculator: AccountFeeCalculator{
								"5000",
							},
						},
					},
					Program: "nonce",
					Space:   80,
				},
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					body, err := ioutil.ReadAll(req.Body)
					assert.Nil(t, err)
					assert.JSONEq(t, tt.requestBody, string(body))
					n, err := rw.Write([]byte(tt.responseBody))
					assert.Nil(t, err)
					assert.Equal(t, len([]byte(tt.responseBody)), n)
				}))
				c := NewClient(server.URL)
				got, err := c.GetAccountInfoParsed(tt.args.ctx, tt.args.base58Addr)
				assert.Equal(t, tt.err, err)
				assert.Equal(t, tt.want, got)
				server.Close()
			})
		})
	}
}
