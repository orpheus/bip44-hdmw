package hdmw

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

type Chain struct {
	Chain *hdkeychain.ExtendedKey
}

type Account struct {
	Account *hdkeychain.ExtendedKey
}

type Coin struct {
	Coin *hdkeychain.ExtendedKey
}

type Wallet struct {
	Mnemonic   string
	Seed       []byte
	Entropy    []byte
	MasterNode *hdkeychain.ExtendedKey
}

func CreateWalletWithPassword(password string) *Wallet {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	seed, _ := bip39.NewSeedWithErrorChecking(mnemonic, password)
	//@ToDo: create network params for FLO and LTC, etc
	masterKey, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)

	wallet := Wallet{
		Mnemonic:   mnemonic,
		Seed:       seed,
		Entropy:    entropy,
		MasterNode: masterKey,
	}

	return &wallet
}

func CreateWalletWithMnemonic(mnemonic, password string) *Wallet {
	seed, _ := bip39.NewSeedWithErrorChecking(mnemonic, password)
	//@ToDo: create network params for FLO and LTC, etc
	masterKey, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)

	wallet := Wallet{
		Mnemonic: mnemonic,
		Seed:     seed,
		//ToDo: generate entropy from mnemonic or seed
		MasterNode: masterKey,
	}

	return &wallet
}

func (w *Wallet) GeneratePurposeNode() (*hdkeychain.ExtendedKey, error) {
	p, err := w.MasterNode.Child(Purpose)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return p, nil
}

//pkg/errors w errors.wrap
//when I capitalize the c in coin in the return params, I get an error
//semantics? better to caps or not the return args?
func (w *Wallet) GenerateCoinNode(bip44CoinConstant uint32) (coin *Coin, err error) {
	p, err := w.GeneratePurposeNode()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	c, err := p.Child(hdkeychain.HardenedKeyStart + bip44CoinConstant)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_coin := &Coin{
		Coin: c,
	}

	return _coin, nil
}

func (c *Coin) GenerateAccountNode(index uint32) (account *Account, err error) {
	a, err := c.Coin.Child(index)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_account := &Account{
		Account: a,
	}

	return _account, nil
}

func (a *Account) GenerateChainNode(index uint32) (chain *Chain, err error) {
	c, err := a.Account.Child(index)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_chain := &Chain{
		Chain: c,
	}

	return _chain, nil
}
