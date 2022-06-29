package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

func fake() (map[string]string, map[string]string) {
	return genFakeHeader(), genFakeData()
}

func fakeCreate(headers, data map[string]string) bool {
	creates := &Create{}

	r, err := client.R().
		SetHeaders(headers).
		SetFormData(data).
		SetResult(creates).
		EnableTrace().
		Post("/garb/v2/trade/create")

	checkErr(err)

	log.Printf("本次请求用时: %v 时间差: %v ms.", r.Request.TraceInfo().TotalTime, diffTime)

	switch creates.Code {
	case 0: // 这里好像有问题，还需要再看看
		if creates.Data.BpEnough == -1 {
			log.Println(r)
			log.Fatalln("余额不足.")
		}
		// 订单号
		orderId = creates.Data.OrderId
		if creates.Data.State != "paying" {
			log.Println(r)
			return true
		}
	case -400:
		log.Fatalln(r)
	case -403: //号被封了
		log.Fatalln("您已被封禁.")
	case 26102: //商品不存在，可能是未到抢购时间，立即重新执行
		log.Println(r)
		go coupon()
		return false
	case 26106: //购买数量达到上限
		log.Fatalln(r)
	case 26120: //请求频率过快，等一下执行，需要测试是否延迟执行
		errorTime += 1
		log.Println(r)
		time.Sleep(500 * time.Millisecond)
		go coupon()
		return false
	case 26113: //号被封了
		//log.Fatalln("当前设备/账号/环境存在风险，暂时无法下单.")
		log.Fatalln(r)
	case 26134: //当前抢购人数过多, 风控等级似乎比 26135 高 (好久没遇到了, 草)
		errorTime += 1
		log.Println(r)
		time.Sleep(500 * time.Millisecond)
		go coupon()
		return false
	case 26135: //当前抢购人数过多，失败四次或者锁四秒后能够购买
		errorTime += 1
		log.Println(r)
		time.Sleep(500 * time.Millisecond)
		go coupon()
		return false
	case 69949: //老风控代码，疑似封锁设备
		errorTime += 1
		log.Println(r)
		time.Sleep(500 * time.Millisecond)
		go coupon()
		return false
	default:
		errorTime += 1
		log.Println(r)
		time.Sleep(500 * time.Millisecond)
		go coupon()
		return false
	}
	return false
}

func genFakeHeader() map[string]string {
	headers := map[string]string{
		"native_api_from":    "h5",
		"refer":              fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1", config.Buy.ItemId),
		"env":                "prod",
		"app-key":            "android64",
		"user-agent":         appUserAgent,
		"x-bili-trace-id":    genFakeTraceID(),
		"x-bili-aurora-eid":  config.Bili.XBiliAuroraEid,
		"x-bili-mid":         config.Cookies.DedeUserID,
		"x-bili-aurora-zone": "",
		"bili-bridge-engine": "cronet",
	}

	return headers
}

func genFakeData() map[string]string {
	data := map[string]string{
		"access_key":   config.AccessKey,
		"add_month":    "-1",
		"appkey":       "1d8b6e7d45233436",
		"buy_num":      config.Buy.BuyNum,
		"coupon_token": config.Buy.CouponToken,
		"csrf":         config.Cookies.BiliJct,
		"currency":     "bp",
		"disable_rcmd": "0",
		"item_id":      config.Buy.ItemId,
		"platform":     config.Buy.Device,
		"statistics":   statistics,
		"ts":           strStartTime,
		"sign":         FakeAppSign(),
	}
	return data
}

// 返回一个虚假的 bili_trace_id
func genFakeTraceID() string {
	x := strStartTime[0:6]
	randBytes := make([]byte, 26/2)
	_, err := rand.Read(randBytes)
	if err != nil {
		return ""
	}
	r := fmt.Sprintf("%x", randBytes)
	xBiliTraceID := fmt.Sprintf("%v%v:%v%v:0:0", r, x, r[16:26], x)
	return xBiliTraceID
}

// FakeAppSign 返回一个虚假的 sign
func FakeAppSign() string {
	var query string
	var buffer bytes.Buffer

	data := map[string]string{
		"access_key":   config.AccessKey,
		"add_month":    "-1",
		"appkey":       "1d8b6e7d45233436",
		"buy_num":      config.Buy.BuyNum,
		"coupon_token": config.Buy.CouponToken,
		"csrf":         config.Cookies.BiliJct,
		"currency":     "bp",
		"disable_rcmd": "0",
		"item_id":      config.Buy.ItemId,
		"platform":     config.Buy.Device,
		"statistics":   statistics,
		"ts":           strStartTime,
	}

	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		//query += fmt.Sprintf("%v=%v&", k, params[k])
		buffer.WriteString(k)
		buffer.WriteString("=")
		buffer.WriteString(data[k])
		buffer.WriteString("&")
	}
	query = strings.TrimRight(buffer.String(), "&")
	//fmt.Println(query)
	sign := strMd5(fmt.Sprintf("%v%v", query, "560c52ccd288fed045859ed18bffd973"))
	//fmt.Println(sign)
	return sign
}
