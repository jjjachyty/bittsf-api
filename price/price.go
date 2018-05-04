package price

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/tidwall/gjson"
)

var CNY_USD = 6.361

type Price struct {
	Name string
	URL  string
	Last float64
	Sell float64
	Buy  float64
	Vol  float64
}

var MarketPrice = make([]Price, 0)
var vpnClient = http.Client{}

func InitPrice() {
	fmt.Println("开始时间", time.Now().String())
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:1087")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:1087")
	waitgroup := sync.WaitGroup{}
	waitgroup.Add(15)
	go OKEx(&waitgroup)
	go Exx(&waitgroup)
	go Gdax(&waitgroup)
	go Huobi(&waitgroup)
	go Zb(&waitgroup)
	go Binance(&waitgroup)
	go Bitfinex(&waitgroup)
	go BitFlyer(&waitgroup)
	go Bitstamp(&waitgroup)
	go Weex(&waitgroup)
	go BitAsiaEx(&waitgroup)
	go Coinnice(&waitgroup)
	go BigOne(&waitgroup)
	go Liqui(&waitgroup)
	go Okcoin(&waitgroup)
	waitgroup.Wait()
	fmt.Println("结束时间", time.Now().String())

}

//获取OKEx行情数据
func OKEx(wg *sync.WaitGroup) error {
	var err error
	resp, err := vpnClient.Get("https://www.okex.com/api/v1/ticker.do?symbol=btc_usdt")
	if err == nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)

		if err == nil {
			ticker := gjson.Get(string(data), "ticker")
			MarketPrice = append(MarketPrice, Price{Name: "OKEx", URL: "https://www.okex.com", Last: ticker.Get("last").Float() * CNY_USD, Sell: ticker.Get("sell").Float() * CNY_USD, Buy: ticker.Get("buy").Float() * CNY_USD, Vol: ticker.Get("vol").Float()})
		}
	}
	defer wg.Done()
	return err
}

//获取OKEx行情数据
func Exx(wg *sync.WaitGroup) error {
	var err error
	resp, err := http.Get("https://api.exx.com/data/v1/ticker?currency=btc_usdt")
	if err == nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			ticker := gjson.Get(string(data), "ticker")
			MarketPrice = append(MarketPrice, Price{Name: "Exx", URL: "https://www.exx.com", Last: ticker.Get("last").Float() * CNY_USD, Sell: ticker.Get("sell").Float() * CNY_USD, Buy: ticker.Get("buy").Float() * CNY_USD, Vol: ticker.Get("vol").Float()})
		}
	}
	defer wg.Done()
	return err
}

//获取Gdax行情数据
func Gdax(wg *sync.WaitGroup) error {
	var err error
	resp, err := http.Get("https://api.gdax.com/products/BTC-USD/ticker")
	if err == nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err == nil {

			MarketPrice = append(MarketPrice, Price{Name: "Gdax", URL: "https://www.gdax.com/", Last: gjson.Get(string(data), "price").Float() * CNY_USD, Sell: gjson.Get(string(data), "ask").Float() * CNY_USD, Buy: gjson.Get(string(data), "bid").Float() * CNY_USD, Vol: gjson.Get(string(data), "volume").Float()})
		}
	}
	defer wg.Done()
	return err
}

//获取OKEx行情数据
func Huobi(wg *sync.WaitGroup) error {
	var err error
	resp, err := vpnClient.Get("https://api.huobipro.com/market/detail/merged?symbol=btcusdt")
	if err == nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			ticker := gjson.Get(string(data), "tick")

			MarketPrice = append(MarketPrice, Price{Name: "火币网", URL: "https://www.huobipro.com/", Last: ticker.Get("close").Float() * CNY_USD, Sell: ticker.Get("ask").Array()[0].Float() * CNY_USD, Buy: ticker.Get("bid").Array()[0].Float() * CNY_USD, Vol: gjson.Get(string(data), "volume").Float()})
		}
	}
	defer wg.Done()
	return err
}

//获取OKEx行情数据
func Zb(wg *sync.WaitGroup) error {
	var err error
	resp, err := http.Get("http://api.zb.com/data/v1/ticker?market=btc_usdt")
	if err == nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			ticker := gjson.Get(string(data), "ticker")
			MarketPrice = append(MarketPrice, Price{Name: "Zb", URL: "https://www.zb.com/", Last: ticker.Get("last").Float() * CNY_USD, Sell: ticker.Get("sell").Float() * CNY_USD, Buy: ticker.Get("buy").Float() * CNY_USD, Vol: ticker.Get("low").Float()})
		}
	}
	defer wg.Done()
	return err
}

//获取币安行情数据
func Binance(wg *sync.WaitGroup) error {
	var err error
	resp, err := http.Get("https://api.binance.com/api/v1/ticker/24hr?symbol=BTCUSDT")
	if err == nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			MarketPrice = append(MarketPrice, Price{Name: "币安", URL: "https://www.binance.com/", Last: gjson.Get(string(data), "lastPrice").Float() * CNY_USD, Sell: gjson.Get(string(data), "askPrice").Float() * CNY_USD, Buy: gjson.Get(string(data), "bidPrice").Float() * CNY_USD, Vol: gjson.Get(string(data), "volume").Float()})
		}
	}
	defer wg.Done()
	return err
}

//获取币安行情数据
func Bitfinex(wg *sync.WaitGroup) error {
	var err error
	resp, err := vpnClient.Get("https://api.bitfinex.com/v1/pubticker/btcusd")
	if err == nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)

		if err == nil {
			MarketPrice = append(MarketPrice, Price{Name: "Bitfinex", URL: "https://www.bitfinex.com/?locale=zh-CN", Last: gjson.Get(string(data), "last_price").Float() * CNY_USD, Sell: gjson.Get(string(data), "ask").Float() * CNY_USD, Buy: gjson.Get(string(data), "bid").Float() * CNY_USD, Vol: gjson.Get(string(data), "volume").Float()})
		}
	}
	defer wg.Done()
	return err
}

//获取币安行情数据
func BitFlyer(wg *sync.WaitGroup) error {
	var err error
	resp, err := vpnClient.Get("https://api.bitflyer.jp/v1/ticker?product_code=BTC_USD")
	if err == nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)

		if err == nil {
			MarketPrice = append(MarketPrice, Price{Name: "BitFlyer", URL: "https://bitflyer.jp/zh-CN/?top_link", Last: gjson.Get(string(data), "ltp").Float() * CNY_USD, Sell: gjson.Get(string(data), "best_ask").Float() * CNY_USD, Buy: gjson.Get(string(data), "best_bid").Float() * CNY_USD, Vol: gjson.Get(string(data), "volume").Float()})
		}
	}
	defer wg.Done()
	return err
}

func Bitstamp(wg *sync.WaitGroup) error {
	var err error
	resp, err := vpnClient.Get("https://www.bitstamp.net/api/ticker/")
	if err == nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)

		if err == nil {
			MarketPrice = append(MarketPrice, Price{Name: "Bitstamp", URL: "https://www.bitstamp.net/", Last: gjson.Get(string(data), "last").Float() * CNY_USD, Sell: gjson.Get(string(data), "bid").Float() * CNY_USD, Buy: gjson.Get(string(data), "ask").Float() * CNY_USD, Vol: gjson.Get(string(data), "volume").Float()})
		}
	}
	defer wg.Done()
	return err
}

func Weex(wg *sync.WaitGroup) error {
	var err error
	resp, err := http.Get("https://api.weex.com/v1/market/ticker?market=BTCUSD")
	if err == nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		ticker := gjson.Get(string(data), "data").Get("ticker")
		if err == nil {
			MarketPrice = append(MarketPrice, Price{Name: "Weex", URL: "https://www.weex.com/", Last: ticker.Get("last").Float() * CNY_USD, Sell: ticker.Get("sell").Float() * CNY_USD, Buy: ticker.Get("buy").Float() * CNY_USD, Vol: ticker.Get("vol").Float()})
		}
	}
	defer wg.Done()
	return err
}
func BitAsiaEx(wg *sync.WaitGroup) error {
	var err error
	resp, err := http.Get("https://www.bitasiabit.com/app/v1/getMarket?pairname=BTCCNY")
	if err == nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		ticker := gjson.Get(string(data), "data")
		if err == nil {
			MarketPrice = append(MarketPrice, Price{Name: "BitAsiaEx-比特亚洲(支持人民币)", URL: "https://www.bitasiaex.com/", Last: ticker.Get("price").Float(), Sell: ticker.Get("ask").Float(), Buy: ticker.Get("bid").Float(), Vol: ticker.Get("total").Float()})
		}
	}
	defer wg.Done()
	return err
}

func Coinnice(wg *sync.WaitGroup) error {
	var err error
	resp, err := http.Get("https://www.coinnice.com/api/v1/spot/btc/ticker")
	if err == nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err == nil && gjson.Get(string(data), "result").Bool() {
			MarketPrice = append(MarketPrice, Price{Name: "Coinnice)", URL: "https://www.coinnice.com/", Last: gjson.Get(string(data), "last").Float() * CNY_USD, Sell: gjson.Get(string(data), "sell").Float() * CNY_USD, Buy: gjson.Get(string(data), "buy").Float() * CNY_USD, Vol: gjson.Get(string(data), "vol").Float()})
		}
	}
	defer wg.Done()
	return err
}

func BigOne(wg *sync.WaitGroup) error {
	var err error
	resp, err := http.Get("https://api.b1.run/markets/BTC-USDT/book")
	if err == nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		ask := gjson.Get(string(data), "data").Get("asks").Array()[0].Get("price").Float()
		bids := gjson.Get(string(data), "data").Get("bids").Array()[0].Get("price").Float()
		if err == nil {
			resp, err := http.Get("https://api.b1.run/markets/BTC-USDT")
			if err == nil {
				defer resp.Body.Close()
				data, err := ioutil.ReadAll(resp.Body)
				if err == nil {
					last := gjson.Get(string(data), "data").Get("ticker").Get("price").Float()
					volume := gjson.Get(string(data), "data").Get("ticker").Get("volume").Float()
					MarketPrice = append(MarketPrice, Price{Name: "BigOne)", URL: "https://b1.run/", Last: last * CNY_USD, Sell: ask * CNY_USD, Buy: bids * CNY_USD, Vol: volume})

				}

			}
		}
	}
	defer wg.Done()
	return err
}

func Liqui(wg *sync.WaitGroup) error {
	var err error
	resp, err := http.Get("https://api.liqui.io/api/3/ticker/btc_usdt")
	if err == nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)

		ticker := gjson.Get(string(data), "btc_usdt")
		if err == nil {
			MarketPrice = append(MarketPrice, Price{Name: "Liqui)", URL: "https://liqui.io/", Last: ticker.Get("last").Float() * CNY_USD, Sell: ticker.Get("sell").Float() * CNY_USD, Buy: ticker.Get("buy").Float() * CNY_USD, Vol: ticker.Get("vol").Float()})
		}
	}
	defer wg.Done()
	return err
}

func Okcoin(wg *sync.WaitGroup) error {
	var err error
	resp, err := http.Get("https://www.okcoin.com/api/v1/ticker.do?symbol=btc_usd")
	if err == nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)

		ticker := gjson.Get(string(data), "ticker")
		if err == nil {
			MarketPrice = append(MarketPrice, Price{Name: "Okcoin)", URL: "https://www.okcoin.com", Last: ticker.Get("last").Float() * CNY_USD, Sell: ticker.Get("sell").Float() * CNY_USD, Buy: ticker.Get("buy").Float() * CNY_USD, Vol: ticker.Get("vol").Float()})
		}
	}
	defer wg.Done()
	return err
}
