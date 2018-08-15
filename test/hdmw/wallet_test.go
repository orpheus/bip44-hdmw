package test

import (
	"testing"
	"github.com/orpheus/bip44-hdmw/hdmw"
	"github.com/davecgh/go-spew/spew"
	)

func TestCreateWallet(t *testing.T) {
	wallet := hdmw.CreateWallet("")
	str := spew.Sdump(wallet)
	t.Log(str)
}