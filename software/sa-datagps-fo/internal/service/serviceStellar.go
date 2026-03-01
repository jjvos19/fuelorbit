package service

import (
	"log"
	"os"

	"github.com/stellar/go-stellar-sdk/clients/horizonclient"
	"github.com/stellar/go-stellar-sdk/keypair"
	"github.com/stellar/go-stellar-sdk/network"
	"github.com/stellar/go-stellar-sdk/txnbuild"
)

type AppServiceStellar struct {
	key *keypair.Full
}

func NewAppServiceStellar() *AppServiceStellar {
	stSecret := os.Getenv("STELLAR_SECRET")

	kp, _ := keypair.ParseFull(stSecret)
	return &AppServiceStellar{key: kp}
}

func (ss *AppServiceStellar) SaveData(name string, data string) string {
	dataByte := []byte(data)
	op := txnbuild.ManageData{
		SourceAccount: ss.key.Address(),
		Name:          name,
		Value:         dataByte,
	}

	client := horizonclient.DefaultTestNetClient
	ar := horizonclient.AccountRequest{AccountID: ss.key.Address()}
	sourceAccount, err := client.AccountDetail(ar)

	tx, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &sourceAccount,
			IncrementSequenceNum: true,
			Operations:           []txnbuild.Operation{&op},
			BaseFee:              txnbuild.MinBaseFee,
			Preconditions: txnbuild.Preconditions{
				TimeBounds: txnbuild.NewTimeout(300),
			},
		},
	)

	if err != nil {
		log.Fatalln(err)
	}

	// Sign the transaction
	tx, err = tx.Sign(network.TestNetworkPassphrase, ss.key)
	if err != nil {
		log.Fatalln(err)
	}
	// Get the base 64 encoded transaction envelope
	txe, err := tx.Base64()
	if err != nil {
		log.Fatalln(err)
	}

	// Send the transaction to the network
	resp, err := client.SubmitTransactionXDR(txe)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Hash: %s", resp.Hash)

	return resp.Hash
}
