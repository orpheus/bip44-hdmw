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
	t.Log(c)
}
