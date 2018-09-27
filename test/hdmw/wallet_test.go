package test

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/davecgh/go-spew/spew"
	"github.com/orpheus/bip44-hdmw/hdmw"
	"testing"
)

const mnemonic = "fragile chalk speed absorb enter weasel hurdle eternal tooth acoustic cost boss"

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

func TestCreateWalletFromSeed(t *testing.T) {
	//seed := "000102030405060708090a0b0c0d0e0f"
	//decoded_seed, _ := hex.DecodeString(seed)
	//wallet := hdmw.CreateWalletFromSeed(decoded_seed, "")
	wallet := hdmw.CreateWalletWithMnemonic(mnemonic, "")

	test, _ := wallet.PurposeNode.Child(hdkeychain.HardenedKeyStart)
	test, _ = test.Child(hdkeychain.HardenedKeyStart)
	test, _ = test.Child(0)
	test, _ = test.Child(0)

	spew.Dump(test.Address(&chaincfg.MainNetParams))
	spew.Dump(test)

	//if test.String() != "xprvA4A9CuBXhdBtCaLxwrw64Jaran4n1rgzeS5mjH47Ds8V67uZS8tTkG8jV3BZi83QqYXPcN4v8EjK2Aof4YcEeqLt688mV57gF4j6QZWdP9U" {
	//	t.Fail()
	//}
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

	spew.Dump(addr)

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
