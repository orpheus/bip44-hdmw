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

func TestCreateAndInitWallet(t *testing.T) {
	wallet := hdmw.CreateWalletWithPassword("")
	wallet.Initialize([]uint32{hdmw.TypeBitcoin})

	if len(wallet.Coins) != 1 {
		t.Error("wallet.coins does not have an len of 1")
	}

	w := spew.Sdump(wallet)
	t.Log(w)
}

func TestDeriveBTCAccount0(t *testing.T) {
	wallet := hdmw.CreateWalletWithPassword("")
	wallet.Initialize([]uint32{hdmw.TypeBitcoin})
	acc, err := wallet.Coins[0].GenerateAccountNode(0)
	if err != nil {
		t.Error("Failed to generate account node")
	}

	a := spew.Sdump(acc)
	t.Log(a)
}

func TestDeriveBTCChain0(t *testing.T) {
	wallet := hdmw.CreateWalletWithPassword("")
	wallet.Initialize([]uint32{hdmw.TypeBitcoin})
	acc, err := wallet.Coins[0].GenerateAccountNode(0)
	if err != nil {
		t.Error("Failed to generate account node")
	}

	ch, err := acc.GenerateChainNode(0)
	if err != nil {
		t.Error("Failed to generate chain node")
	}

	c := spew.Sdump(ch)
	t.Log(c)
}

func TestDeriveBTCAddress0(t *testing.T) {
	wallet := hdmw.CreateWalletWithPassword("")
	wallet.Initialize([]uint32{hdmw.TypeBitcoin})
	acc, err := wallet.Coins[0].GenerateAccountNode(0)
	if err != nil {
		t.Error("Failed to generate account node")
	}

	ch, err := acc.GenerateChainNode(0)
	if err != nil {
		t.Error("Failed to generate chain node")
	}

	addr, err := ch.GenerateAddressNode(0)
	if err != nil {
		t.Error("Failed to generate address node")
	}

	a := spew.Sdump(addr)
	t.Log(a)
}
