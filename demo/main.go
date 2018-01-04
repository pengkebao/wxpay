package main

import "github.com/pengkebao/wxpay"
import "fmt"

func main() {
	payConfig := wxpay.WxPayConfig{
		AppId:     "wx8a92954451ec0dfa",
		AppKey:    "y7c2743910e94b220b0048af87437881",
		MchId:     "1317382401",
		NotifyUrl: "http://test.ckb.mobi",
		TradeType: "APP", // 支持 JSAPI，NATIVE，APP
	}
	// 统一下单
	params := map[string]string{
		"out_trade_no": "1000000000000", // 商户订单号
		"body":         "test",          // 商品描述
		"total_fee":    "100",           // 100分 = 1元
	}
	resp, err := wxpay.UnifiedOrder(payConfig, params)
	fmt.Println(err, resp.PrepayId)
}
