package handlers

import (
	"github.com/utopia-eco/walletReader/consts"
	"github.com/utopia-eco/walletReader/models"
	"reflect"
	"testing"
)

func TestGetWalletTokenValues(t *testing.T) {
	type args struct {
		req *models.WalletTokensValueReq
	}
	tests := []struct {
		name    string
		args    args
		want    *models.WalletTokensValueResp
		wantErr bool
	}{
		{
			name: "success test",
			args: args{req: &models.WalletTokensValueReq{
				WalletAddress: "abc",
				Tokens: []*models.Token{
					{
						TokenAddress: consts.BnbAddr,
						Amount:       5.6,
					},
					{
						TokenAddress: "0xfb6115445bff7b52feb98650c87f44907e58f802",
						Amount:       12,
					},
					{
						TokenAddress: "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d",
						Amount:       1.1,
					},
				},
			}},
			want: &models.WalletTokensValueResp{
				WalletAddress: "abc",
				WalletValue:   0,
				TokenValues: []*models.TokenValue{
					{
						TokenAddress: consts.BnbAddr,
						TokenSymbol:  "BNB",
						UnitPrice:    0,
						TotalValue:   0,
					},
					{
						TokenAddress: 0xfb6115445bff7b52feb98650c87f44907e58f802,
						TokenSymbol:  "AAVE",
						UnitPrice:    0,
						TotalValue:   0,
					},
					{
						TokenAddress: "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d",
						TokenSymbol:  "USDC",
						UnitPrice:    0,
						TotalValue:   0,
					},
				},
			},
		},
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetWalletTokenValues(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWalletTokenValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWalletTokenValues() got = %v, want %v", got, tt.want)
			}
		})
	}
}
