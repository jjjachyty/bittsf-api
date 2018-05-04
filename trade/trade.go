package trade

import (
	"bittsf/store"
	"bittsf/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Weex struct{}

func (Weex) Sell() {
	data := make(url.Values)
	data["access_id"] = []string{store.MarKetAPIs.Weex.AccessID}
	data["market"] = []string{"BTCUSD"}
	data["type"] = []string{"buy"}
	data["source_id"] = []string{"weex"}

	var authorization = utils.EncodeMD5(data.Encode())
	resp, err := http.PostForm("https://api.weex.com/v1/order/market", data)
	resp.Header.Set("authorization", authorization)
	if err == nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Println("data", string(data))
		}
	}
}
