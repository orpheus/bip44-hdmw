package test

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/orpheus/bip44-hdmw/hdmw"
	"testing"
)

func TestCreateWallet(t *testing.T) {
	wallet := hdmw.CreateWalletWithPassword("")
	str := spew.Sdump(wallet)
	t.Log(str)
}

func TestCreatePurposeNode(t *testing.T) {
	wallet := hdmw.CreateWalletWithPassword("")
	str := spew.Sdump(wallet)
	t.Log(str)
	p, err := wallet.GeneratePurposeNode()
	if err != nil {
		t.Log(err)
	}
	t.Log(p)
}

func TestCreateBitcoinNode(t *testing.T) {
	wallet := hdmw.CreateWalletWithPassword("")
	str := spew.Sdump(wallet)
	t.Log(str)
	c, err := wallet.GenerateCoinNode(hdmw.TypeBitcoin)
	if err != nil {
		t.Log(err)
	}
	coin := spew.Sdump(c)
	t.Log(coin)
}

func TestCreateAndInitWallet(t *testing.T) {
	wallet := hdmw.CreateWalletWithPassword("")
	wallet.Initialize([]uint32{hdmw.TypeBitcoin})

	if len(wallet.Coins) != 1 {
		t.Error("wallet.coins does not have an len of 1")
	}

	w := spew.Sdump(wallet)
	t.Log(w)
}
