// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: node.proto

package node

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WalletMock struct{}

func (m *WalletMock) CreateWallet(ctx context.Context, req *CreateRequest) (*LoadResponse, error) {
	res :=
		&LoadResponse{
			Key: &PubKey{},
		}
	return res, nil
}
func (m *WalletMock) LoadWallet(ctx context.Context, req *LoadRequest) (*LoadResponse, error) {
	res :=
		&LoadResponse{
			Key: &PubKey{},
		}
	return res, nil
}
func (m *WalletMock) CreateFromSeed(ctx context.Context, req *CreateRequest) (*LoadResponse, error) {
	res :=
		&LoadResponse{
			Key: &PubKey{},
		}
	return res, nil
}
func (m *WalletMock) ClearWalletDatabase(ctx context.Context, req *EmptyRequest) (*GenericResponse, error) {
	res :=
		&GenericResponse{
			Response: "ullam",
		}
	return res, nil
}
func (m *WalletMock) GetWalletStatus(ctx context.Context, req *EmptyRequest) (*WalletStatusResponse, error) {
	res :=
		&WalletStatusResponse{
			Loaded: false,
		}
	return res, nil
}
func (m *WalletMock) GetAddress(ctx context.Context, req *EmptyRequest) (*LoadResponse, error) {
	res :=
		&LoadResponse{
			Key: &PubKey{},
		}
	return res, nil
}
func (m *WalletMock) GetBalance(ctx context.Context, req *EmptyRequest) (*BalanceResponse, error) {
	res :=
		&BalanceResponse{}
	return res, nil
}
func (m *WalletMock) GetTxHistory(ctx context.Context, req *EmptyRequest) (*TxHistoryResponse, error) {
	res :=
		&TxHistoryResponse{
			Records: []*TxRecord{
				&TxRecord{
					Direction: 1,
					Timestamp: 71,
					Type:      5,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 708,
					Type:      0,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 336,
					Type:      2,
				},
				&TxRecord{
					Direction: 1,
					Timestamp: 569,
					Type:      4,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 526,
					Type:      1,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 758,
					Type:      2,
				},
				&TxRecord{
					Direction: 1,
					Timestamp: 205,
					Type:      2,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 614,
					Type:      3,
				},
			},
		}
	return res, nil
}

type MempoolMock struct{}

func (m *MempoolMock) GetUnconfirmedBalance(ctx context.Context, req *EmptyRequest) (*BalanceResponse, error) {
	res :=
		&BalanceResponse{}
	return res, nil
}
func (m *MempoolMock) SelectTx(ctx context.Context, req *SelectRequest) (*SelectResponse, error) {
	res :=
		&SelectResponse{
			Result: []*Tx{
				&Tx{
					Type: 4,
					Id:   "9be01538-3680-4fe0-8ca1-55926b96da94",
				},
				&Tx{
					Type: 1,
					Id:   "3f2715e7-c976-4d02-8d80-0141bcee1532",
				},
				&Tx{
					Type: 3,
					Id:   "7167fc92-3350-45c8-ac96-b7c0f5198480",
				},
				&Tx{
					Type: 1,
					Id:   "ec23fd89-7f22-40b7-9521-d0a945041b56",
				},
				&Tx{
					Type: 0,
					Id:   "2732cfb0-8d15-488c-89ba-10fc070aed85",
				},
				&Tx{
					Type: 0,
					Id:   "43861f15-f05c-4d20-b166-edc22e383dc6",
				},
				&Tx{
					Type: 2,
					Id:   "989c3ee0-0704-4fd0-9b7e-4cd3c63da8a8",
				},
				&Tx{
					Type: 3,
					Id:   "68c1bfca-c150-46ea-b8fc-75231b639439",
				},
			},
		}
	return res, nil
}

type ChainMock struct{}

func (m *ChainMock) RebuildChain(ctx context.Context, req *EmptyRequest) (*GenericResponse, error) {
	res :=
		&GenericResponse{
			Response: "quidem",
		}
	return res, nil
}
func (m *ChainMock) GetSyncProgress(ctx context.Context, req *EmptyRequest) (*SyncProgressResponse, error) {
	res :=
		&SyncProgressResponse{
			Progress: 531.9317,
		}
	return res, nil
}

type TransactorMock struct{}

func (m *TransactorMock) CallContract(ctx context.Context, req *CallContractRequest) (*TransactionResponse, error) {
	res :=
		&TransactionResponse{}
	return res, nil
}
func (m *TransactorMock) Transfer(ctx context.Context, req *TransferRequest) (*TransactionResponse, error) {
	res :=
		&TransactionResponse{}
	return res, nil
}
func (m *TransactorMock) Bid(ctx context.Context, req *BidRequest) (*TransactionResponse, error) {
	res :=
		&TransactionResponse{}
	return res, nil
}
func (m *TransactorMock) Stake(ctx context.Context, req *StakeRequest) (*TransactionResponse, error) {
	res :=
		&TransactionResponse{}
	return res, nil
}

type MaintainerMock struct{}

func (m *MaintainerMock) AutomateConsensusTxs(ctx context.Context, req *EmptyRequest) (*GenericResponse, error) {
	res :=
		&GenericResponse{
			Response: "rem",
		}
	return res, nil
}
