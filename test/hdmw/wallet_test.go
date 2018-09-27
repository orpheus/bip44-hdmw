package test

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/davecgh/go-spew/spew"
	"github.com/orpheus/bip44-hdmw/hdmw"
	"testing"
)

const (
	mnemonic = "alien indicate barely erosion please recall start together anger panic law latin"
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
		t.Error("Error generating MasterNodes. Failed equality check.")
	}
}

func TestDeriveMainAddressFromMnemonic(t *testing.T) {
	wallet := hdmw.CreateWalletWithMnemonic(mnemonic, "")
	wallet.Initialize([]uint32{hdmw.TypeBitcoin})

	btc := wallet.Coins[0]
	acc, err := btc.DeriveAccountNode(hdkeychain.HardenedKeyStart + 0)
	if err != nil {
		t.Error("Failed to derive account node")
	}

	ch, err := acc.DeriveChainNode(0)
	if err != nil {
		t.Error("Failed to derive chain node")
	}

	addr, err := ch.DeriveAddressNode(0)
	if err != nil {
		t.Error("Failed to derive address node")
	}

	p2p, err := addr.Address.Address(&chaincfg.MainNetParams)

	t.Log(p2p)
	t.Error("Failing to derive address correctly")

}

func TestDeriveBTCAccount0(t *testing.T) {
	wallet := hdmw.CreateWalletWithPassword("")
	wallet.Initialize([]uint32{hdmw.TypeBitcoin})
	acc, err := wallet.Coins[0].DeriveAccountNode(0)
	if err != nil {
		t.Error("Failed to derive account node")
	}

	a := spew.Sdump(acc)
	t.Log(a)
}

func TestDeriveBTCChain0(t *testing.T) {
	wallet := hdmw.CreateWalletWithPassword("")
	wallet.Initialize([]uint32{hdmw.TypeBitcoin})
	acc, err := wallet.Coins[0].DeriveAccountNode(0)
	if err != nil {
		t.Error("Failed to derive account node")
	}

	ch, err := acc.DeriveChainNode(0)
	if err != nil {
		t.Error("Failed to derive chain node")
	}

	c := spew.Sdump(ch)
	t.Log(c)
}

func TestDeriveBTCAddress0(t *testing.T) {
	wallet := hdmw.CreateWalletWithPassword("")
	wallet.Initialize([]uint32{hdmw.TypeBitcoin})
	acc, err := wallet.Coins[0].DeriveAccountNode(0)
	if err != nil {
		t.Error("Failed to derive account node")
	}

	ch, err := acc.DeriveChainNode(0)
	if err != nil {
		t.Error("Failed to derive chain node")
	}

	addr, err := ch.DeriveAddressNode(0)
	if err != nil {
		t.Error("Failed to derive address node")
	}

	a := spew.Sdump(addr)
	t.Log(a)
}
