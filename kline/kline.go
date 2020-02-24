package kline

type Kline struct {
	CoinType   string // 交易币种对
	High       string // 最高价
	Low        string // 最低价
	Open       string // 开盘价
	Close      string // 收盘价
	CreateTime int64  // 创建时间
	TimeScale  int64  // 时间刻度
	Volume     int64  // 24小时交易量
}

func (k *Kline) Copy() *Kline {
	return &Kline{
		CoinType:   k.CoinType,
		High:       k.High,
		Low:        k.Low,
		Open:       k.Open,
		Close:      k.Close,
		CreateTime: k.CreateTime,
		TimeScale:  k.TimeScale,
		Volume:     k.Volume,
	}
}
