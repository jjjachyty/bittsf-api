package main

import (
	"bittsf/price"
	"bittsf/trade"
)

//export http_proxy=http://127.0.0.1:1087;export https_proxy=http://127.0.0.1:1087;
var BuyTop3 = make([]price.Price, 3)
var SellTop3 = make([]price.Price, 3)

type PriceWrapper struct {
	Price []price.Price
	by    func(p, q *price.Price) bool
}
type SortBy func(p, q *price.Price) bool

func (c PriceWrapper) Len() int {
	return len(c.Price)
}
func (c PriceWrapper) Swap(i, j int) {
	c.Price[i], c.Price[j] = c.Price[j], c.Price[i]
}
func (c PriceWrapper) Less(i, j int) bool {
	return c.by(&c.Price[i], &c.Price[j])
}

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	// price.InitPrice()
	// // fmt.Println("之前", price.MarketPrice)

	// sort.Sort(PriceWrapper{price.MarketPrice, func(p, q *price.Price) bool {
	// 	return q.Buy < p.Buy
	// }})

	// copy(BuyTop3, price.MarketPrice[:3])

	// // for _, k := range price.MarketPrice[:3] {
	// // 	fmt.Println("buy-----", k.Name, k.Buy)
	// // }

	// sort.Sort(PriceWrapper{price.MarketPrice, func(p, q *price.Price) bool {
	// 	return q.Sell > p.Sell
	// }})
	// copy(SellTop3, price.MarketPrice[:3])
	// // for _, k := range SellTop3 {
	// // 	fmt.Println("sell-----", k.Name, k.Sell)
	// // }
	// fmt.Println("SellTop3", SellTop3[0])
	// fmt.Println("BuyTop3", BuyTop3[0])
	// store.Write()
	// store.Read()
	trade.Weex{}.Sell()
}
