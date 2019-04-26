package test

import (
	"encoding/hex"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/orpheus/bip44-hdmw/hdmw"
	"testing"
)

const testMnemonic1 = "fragile chalk speed absorb enter weasel hurdle eternal tooth acoustic cost boss"

func TestCreateWallet(t *testing.T) {
	wallet, _ := hdmw.CreateWalletWithPassword("")
	if wallet.PurposeNode == nil {
		t.Fail()
	}
}

func TestCreateWalletFromhMnemonic(t *testing.T) {
	wallet, _ := hdmw.CreateWalletWithPassword("")
	wallet2, _ := hdmw.CreateWalletFromMnemonic(wallet.Mnemonic, "")

	if wallet.Seed != wallet2.Seed {
		t.Error("Mnemonics did not produce the same seed")
	}

	m1 := wallet.MasterNode.String()
	m2 := wallet2.MasterNode.String()

	if m1 != m2 {
		t.Error("Error generating MasterNodes. Failed equality check.")
	}
}

func TestCreateWalletFromSeed(t *testing.T) {
	seed := "000102030405060708090a0b0c0d0e0f"
	decodedSeed, _ := hex.DecodeString(seed)
	wallet, _ := hdmw.CreateWalletFromSeed(decodedSeed)

	coin, _ := wallet.PurposeNode.Child(hdkeychain.HardenedKeyStart)
	account, _ := coin.Child(hdkeychain.HardenedKeyStart)
	chain, _ := account.Child(0)
	address, _ := chain.Child(0)

	if address.String() != "xprvA4A9CuBXhdBtCaLxwrw64Jaran4n1rgzeS5mjH47Ds8V67uZS8tTkG8jV3BZi83QqYXPcN4v8EjK2Aof4YcEeqLt688mV57gF4j6QZWdP9U" {
		t.Fail()
	}
}

func TestCreateAndInitWallet(t *testing.T) {
	wallet, _ := hdmw.CreateWalletWithPassword("")
	wallet.Initialize([]uint32{hdmw.TypeBitcoin})

	if len(wallet.Coins) != 1 {
		t.Error("wallet.coins does not have an len of 1")
	}
}

func TestDeriveMainAddressFromMnemonic(t *testing.T) {
	wallet, _ := hdmw.CreateWalletFromMnemonic(testMnemonic1, "")
	wallet.Initialize([]uint32{hdmw.TypeBitcoin, hdmw.TypeFlo})

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

	address, _ := addr.Address.Address(addr.Network)
	if address.String() != "17MvGBBDkFcMixe43TD6PnRgVRNS7tXm6" {
		t.Fail()
	}
}
