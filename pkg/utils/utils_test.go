package utils_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/forta-network/go-multicall"
	ssv_network_views "github.com/stafiprotocol/eth-lsd-ssv-client/bindings/SsvNetworkViews"
	"github.com/stafiprotocol/eth-lsd-ssv-client/pkg/utils"
)

func TestAppendFile(t *testing.T) {
	path := "../../log_data/append_test2.txt"
	lastLine, err := utils.ReadLastLine(path)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(lastLine)
	err = utils.AppendToFile(path, "\ntest1")
	if err != nil {
		t.Fatal(err)
	}
	err = utils.AppendToFile(path, "\ntest1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultiCall(t *testing.T) {
	caller, err := multicall.Dial(context.Background(), "https://rpc.ankr.com/eth_goerli")
	if err != nil {
		t.Fatal(err)
	}
	contract, err := multicall.NewContract(ssv_network_views.SsvNetworkViewsABI, "0xAE2C84c48272F5a1746150ef333D5E5B51F68763")
	if err != nil {
		t.Fatal(err)
	}

	calls := make([]*multicall.Call, 1)
	calls[0] = contract.NewCall(
		new(utils.OperatorInfoOnChain),
		"getOperatorById",
		uint64(1),
	).Name("get operator")
	_, err = caller.Call(nil,
		calls...,
	)
	if err != nil {
		t.Fatal(err)
	}

	for _, call := range calls {
		fmt.Println(call.CallName, ":", call.Outputs.(*utils.OperatorInfoOnChain))
	}
}

func TestBatchGetOperators(t *testing.T) {
	caller, err := multicall.Dial(context.Background(), "https://rpc.ankr.com/eth_goerli")
	if err != nil {
		t.Fatal(err)
	}
	ops, err := utils.BatchGetOperators(caller, common.HexToAddress("0xAE2C84c48272F5a1746150ef333D5E5B51F68763"), []uint64{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}
	for _, op := range ops {

		t.Logf("%+v", op)
	}
}

func TestBatchGetClusterBalance(t *testing.T) {
	caller, err := multicall.Dial(context.Background(), "https://rpc.ankr.com/eth_goerli")
	if err != nil {
		t.Fatal(err)
	}
	ops, err := utils.BatchGetClusterBalance(
		caller,
		common.HexToAddress("0xAE2C84c48272F5a1746150ef333D5E5B51F68763"),
		common.HexToAddress("0xAE2C84c48272F5a1746150ef333D5E5B51F68763"),
		[][]uint64{{1, 2, 3, 4}},
		[]*ssv_network_views.ISSVNetworkCoreCluster{&ssv_network_views.ISSVNetworkCoreCluster{
			ValidatorCount:  1,
			NetworkFeeIndex: 0,
			Index:           0,
			Active:          false,
			Balance:         big.NewInt(5),
		}})
	if err != nil {
		t.Fatal(err)
	}
	for _, op := range ops {

		t.Logf("%+v", op)
	}
}
