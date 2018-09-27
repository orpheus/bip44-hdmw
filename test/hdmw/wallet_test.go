package test

import (
	"encoding/hex"
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

func TestCreateWalletWithMnemonic(t *testing.T) {
	wallet := hdmw.CreateWalletWithPassword("")
	t.Log(wallet.Seed)

	wallet2 := hdmw.CreateWalletWithMnemonic(wallet.Mnemonic, "")
	t.Log(wallet2.Seed)

	s1 := hex.EncodeToString(wallet.Seed)
	s2 := hex.EncodeToString(wallet2.Seed)
	if s1 != s2 {
		t.Error("Mnemonics did not produce the same seed")
	}

	m1 := wallet.MasterNode.String()
	m2 := wallet2.MasterNode.String()
	t.Log(m1)
	t.Log(m2)

	if m1 != m2 {
		t.Error("Error generation MasterNodes. Failed equality check.")
	}
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
