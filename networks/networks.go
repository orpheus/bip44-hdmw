package networks

//import insight-explorer
//import varIntBuffer

var floFeePerKb = 100000
var bitcoinFeePerKb = 100000
var litecoinFeePerKb = 100000

type NetworkInterface interface {
}

type CoinNetwork struct {
	CoinName      string
	DisplayName   string
	Ticker        string
	SatPerCoin    float32
	FeePerKb      int
	FeePerByte    int
	MaxFeePerByte int
	MinFee        int
	Dust          int
	TxVersion     int
	Explorer      *Insight
	GetExtraBytes func()
	Network       *Network
}
//insight -> interface
var FloNetwork = CoinNetwork{
	CoinName:      "flo",
	DisplayName:   "Flo",
	Ticker:        "FLO",
	SatPerCoin:    1e8,
	FeePerKb:      floFeePerKb,
	FeePerByte:    floFeePerKb / 1024,
	MaxFeePerByte: 100,
	MinFee:        floFeePerKb,
	Dust:          100000,
	TxVersion:     2,
	Explorer:      nil,
	GetExtraBytes: nil,
	Network:       nil,
}
