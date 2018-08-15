package hdmw

type CoinInterface interface {
	fromSeed()
	setAccount()
	getAccount()
}

type Coin struct {
	Coin        string
	Accounts    []*Account
	discover    bool
	Bip44Number int
	Root        int
}

func (c *Coin) fromSeed()   {}
func (c *Coin) setAccount() {}
func (c *Coin) getAccount() {}
