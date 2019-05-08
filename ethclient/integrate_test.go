package ethclient

import (
	"context"
	"fmt"
	"github.com/mrFranklin/web3go/common"
	"github.com/mrFranklin/web3go/core/types"
	"github.com/mrFranklin/web3go/crypto"
	"testing"
)

// TestSendTransaction tests ethclient.SendTransaction
// If run the test, you must set the flags `--miner.gasprice 0 --txpool.pricelimit 0 ` when starting the geth client
func TestSendTransaction(t *testing.T) {
	prikey, _ := crypto.GenerateKey()
	prikey.Public()

	source := "http://127.0.0.1:8545"
	client, err := Dial(source)
	if err != nil {
		t.Fatalf("error when dail ip: %v", err)
	}
	to := common.HexToAddress("0x093477FC2F72F5410132B95D71D12C40098C94E7")
	tx := types.NewTransaction(0, to, common.Big0, 1000000, common.Big0, []byte("0x00"))
	signer := types.HomesteadSigner{}
	signedTx, err := types.SignTx(tx, signer, prikey)
	if err != nil {
		t.Fatalf("error when sign tx: %v", err)
	}
	hash, err := client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Fatalf("error when send tx: %v", err)
	}
	fmt.Println(hash.String())
}