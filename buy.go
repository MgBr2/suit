package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

var (
	navs        = &Navs{}
	details     = &Details{}
	checkCoupon bool
)

func nav() {
	headers := map[string]string{
		"Content-Type":       "application/json, text/plain, */*",
		"native_api_from":    "h5",
		"refer":              "https://www.bilibili.com/h5/mall/home?navhide=1&from=myservice&native.theme=1",
		"env":                "prod",
		"app-key":            "android64",
		"user-agent":         appUserAgent,
		"x-bili-trace-id":    genTraceID(),
		"x-bili-aurora-eid":  config.Bili.XBiliAuroraEid,
		"x-bili-mid":         config.Cookies.DedeUserID,
		"x-bili-aurora-zone": "",
		"bili-bridge-engine": "cronet",
	}

	params := map[string]string{
		"access_key":   config.AccessKey,
		"appkey":       "1d8b6e7d45233436",
		"csrf":         config.Cookies.BiliJct,
		"disable_rcmd": "0",
		"statistics":   statistics,
		"ts":           strconv.FormatInt(time.Now().Unix(), 10),
	}

	sign := appSign(params)
	params["sign"] = sign

	_, err := client.R().
		SetHeaders(headers).
		SetResult(navs).
		SetQueryParams(params).
		Get("/web-interface/nav")

	checkErr(err)
}

// popup, 未知用途
func popup() {
	headers := map[string]string{
		"Content-Type":       "application/json, text/plain, */*",
		"native_api_from":    "h5",
		"refer":              "https://www.bilibili.com/h5/mall/home?navhide=1&from=myservice&native.theme=1",
		"env":                "prod",
		"app-key":            "android64",
		"user-agent":         appUserAgent,
		"x-bili-trace-id":    genTraceID(),
		"x-bili-aurora-eid":  config.Bili.XBiliAuroraEid,
		"x-bili-mid":         config.Cookies.DedeUserID,
		"x-bili-aurora-zone": "",
		"bili-bridge-engine": "cronet",
	}

	params := map[string]string{
		"access_key":   config.AccessKey,
		"appkey":       "1d8b6e7d45233436",
		"csrf":         config.Cookies.BiliJct,
		"disable_rcmd": "0",
		"statistics":   statistics,
		"ts":           strconv.FormatInt(time.Now().Unix(), 10),
	}

	sign := appSign(params)
	params["sign"] = sign

	_, err := client.R().
		SetHeaders(headers).
		SetQueryParams(params).
		Get("/garb/popup")

	checkErr(err)
}

// 装扮信息
func detail() {
	headers := map[string]string{
		"Content-Type":       "application/json, text/plain, */*",
		"native_api_from":    "h5",
		"refer":              fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1", config.Buy.ItemId),
		"env":                "prod",
		"app-key":            "android64",
		"user-agent":         appUserAgent,
		"x-bili-trace-id":    genTraceID(),
		"x-bili-aurora-eid":  config.Bili.XBiliAuroraEid,
		"x-bili-mid":         config.Cookies.DedeUserID,
		"x-bili-aurora-zone": "",
		"bili-bridge-engine": "cronet",
	}

	params := map[string]string{
		"access_key":   config.AccessKey,
		"appkey":       "1d8b6e7d45233436",
		"csrf":         config.Cookies.BiliJct,
		"disable_rcmd": "0",
		"item_id":      config.Buy.ItemId,
		"part":         "suit",
		"statistics":   statistics,
		"ts":           strconv.FormatInt(time.Now().Unix(), 10),
	}

	sign := appSign(params)
	params["sign"] = sign

	_, err := client.R().
		SetHeaders(headers).
		SetQueryParams(params).
		SetResult(details).
		Get("/garb/v2/mall/suit/detail")

	checkErr(err)
}

// 拥有的装扮信息
func asset() {
	headers := map[string]string{
		"native_api_from":    "h5",
		"refer":              fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1", config.Buy.ItemId),
		"env":                "prod",
		"app-key":            "android64",
		"user-agent":         appUserAgent,
		"x-bili-trace-id":    genTraceID(),
		"x-bili-aurora-eid":  config.Bili.XBiliAuroraEid,
		"x-bili-mid":         config.Cookies.DedeUserID,
		"x-bili-aurora-zone": "",
		"bili-bridge-engine": "cronet",
	}

	params := map[string]string{
		"access_key":   config.AccessKey,
		"appkey":       "1d8b6e7d45233436",
		"csrf":         config.Cookies.BiliJct,
		"disable_rcmd": "0",
		"item_id":      config.Buy.ItemId,
		"part":         "suit",
		"statistics":   statistics,
		"ts":           strconv.FormatInt(time.Now().Unix(), 10),
	}

	sign := appSign(params)
	params["sign"] = sign
	response := &Asset{}

	_, err := client.R().
		SetHeaders(headers).
		SetQueryParams(params).
		SetResult(response).
		Get("/garb/user/asset")

	checkErr(err)
}

// 装扮排名
func rank() {
	headers := map[string]string{
		"Content-Type":       "application/json, text/plain, */*",
		"native_api_from":    "h5",
		"refer":              fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1", config.Buy.ItemId),
		"env":                "prod",
		"app-key":            "android64",
		"user-agent":         appUserAgent,
		"x-bili-trace-id":    genTraceID(),
		"x-bili-aurora-eid":  config.Bili.XBiliAuroraEid,
		"x-bili-mid":         config.Cookies.DedeUserID,
		"x-bili-aurora-zone": "",
		"bili-bridge-engine": "cronet",
	}

	params := map[string]string{
		"access_key":   config.AccessKey,
		"appkey":       "1d8b6e7d45233436",
		"csrf":         config.Cookies.BiliJct,
		"disable_rcmd": "0",
		"item_id":      config.Buy.ItemId,
		"part":         "suit",
		"statistics":   statistics,
		"ts":           strconv.FormatInt(time.Now().Unix(), 10),
	}

	sign := appSign(params)
	params["sign"] = sign
	ranks := &Rank{}

	_, err := client.R().
		SetHeaders(headers).
		SetQueryParams(params).
		SetResult(ranks).
		Get("/garb/rank/fan/recent")

	checkErr(err)

	rankInfo = ranks
}

// 关于这个装扮的订单情况
func stat() {
	headers := map[string]string{
		"Content-Type":       "application/json, text/plain, */*",
		"native_api_from":    "h5",
		"refer":              fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1", config.Buy.ItemId),
		"env":                "prod",
		"app-key":            "android64",
		"user-agent":         appUserAgent,
		"x-bili-trace-id":    genTraceID(),
		"x-bili-aurora-eid":  config.Bili.XBiliAuroraEid,
		"x-bili-mid":         config.Cookies.DedeUserID,
		"x-bili-aurora-zone": "",
		"bili-bridge-engine": "cronet",
	}

	params := map[string]string{
		"access_key":   config.AccessKey,
		"appkey":       "1d8b6e7d45233436",
		"csrf":         config.Cookies.BiliJct,
		"disable_rcmd": "0",
		"item_id":      config.Buy.ItemId,
		"part":         "suit",
		"statistics":   statistics,
		"ts":           strconv.FormatInt(time.Now().Unix(), 10),
	}

	sign := appSign(params)
	params["sign"] = sign

	_, err := client.R().
		SetHeaders(headers).
		SetQueryParams(params).
		Get("/garb/order/user/stat")

	checkErr(err)
}

// 预约人数及预约状态
func state() {
	headers := map[string]string{
		"Content-Type":       "application/json, text/plain, */*",
		"native_api_from":    "h5",
		"refer":              fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1", config.Buy.ItemId),
		"env":                "prod",
		"app-key":            "android64",
		"user-agent":         appUserAgent,
		"x-bili-trace-id":    genTraceID(),
		"x-bili-aurora-eid":  config.Bili.XBiliAuroraEid,
		"x-bili-mid":         config.Cookies.DedeUserID,
		"x-bili-aurora-zone": "",
		"bili-bridge-engine": "cronet",
	}

	params := map[string]string{
		"access_key":   config.AccessKey,
		"appkey":       "1d8b6e7d45233436",
		"csrf":         config.Cookies.BiliJct,
		"disable_rcmd": "0",
		"item_id":      config.Buy.ItemId,
		"part":         "suit",
		"statistics":   statistics,
		"ts":           strconv.FormatInt(time.Now().Unix(), 10),
	}

	reserveInfo := &Reserve{}
	sign := appSign(params)
	params["sign"] = sign

	_, err := client.R().
		SetHeaders(headers).
		SetResult(reserveInfo).
		SetQueryParams(params).
		Get("/garb/user/reserve/state")

	checkErr(err)

	if reserveInfo.Data.Reserved == false && reserveInfo.Data.ReserveState == true {
		log.Println("你还没有预约哦，Mika 这就帮你预约喵～")
		reserve()
	} else if reserveInfo.Data.Reserved == true && reserveInfo.Data.ReserveState == true {
		log.Println("当前装扮已经预约了喵～")
	}
}

func reserve() {
	headers := map[string]string{
		"Content-Type":       "application/json, text/plain, */*",
		"native_api_from":    "h5",
		"refer":              fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1", config.Buy.ItemId),
		"env":                "prod",
		"app-key":            "android64",
		"user-agent":         appUserAgent,
		"x-bili-trace-id":    genTraceID(),
		"x-bili-aurora-eid":  config.Bili.XBiliAuroraEid,
		"x-bili-mid":         config.Cookies.DedeUserID,
		"x-bili-aurora-zone": "",
		"bili-bridge-engine": "cronet",
	}

	data := map[string]string{
		"access_key":   config.AccessKey,
		"appkey":       "1d8b6e7d45233436",
		"csrf":         config.Cookies.BiliJct,
		"disable_rcmd": "0",
		"item_id":      config.Buy.ItemId,
		"part":         "suit",
		"statistics":   statistics,
		"ts":           strconv.FormatInt(time.Now().Unix(), 10),
	}

	reserveInfo := Reserve{}
	sign := appSign(data)
	data["sign"] = sign

	_, err := client.R().
		SetHeaders(headers).
		SetResult(reserveInfo).
		SetFormData(data).
		Post("/garb/user/reserve")

	checkErr(err)

	if reserveInfo.Code != 0 {
		go log.Println("预约失败了喵！")
	}
}

// 优惠券（有多张优惠券时的情况未收集!）
func coupon() {
	headers := map[string]string{
		"Content-Type":       "application/json, text/plain, */*",
		"native_api_from":    "h5",
		"refer":              fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1", config.Buy.ItemId),
		"env":                "prod",
		"app-key":            "android64",
		"user-agent":         appUserAgent,
		"x-bili-trace-id":    genTraceID(),
		"x-bili-aurora-eid":  config.Bili.XBiliAuroraEid,
		"x-bili-mid":         config.Cookies.DedeUserID,
		"x-bili-aurora-zone": "",
		"bili-bridge-engine": "cronet",
	}

	params := map[string]string{
		"access_key":   config.AccessKey,
		"appkey":       "1d8b6e7d45233436",
		"csrf":         config.Cookies.BiliJct,
		"disable_rcmd": "0",
		"item_id":      config.Buy.ItemId,
		"part":         "suit",
		"statistics":   statistics,
		"ts":           strconv.FormatInt(time.Now().Unix(), 10),
	}

	c := &Coupon{}
	sign := appSign(params)
	params["sign"] = sign

	_, err := client.R().
		SetHeaders(headers).
		SetResult(c).
		SetQueryParams(params).
		Get("/garb/coupon/usable")

	checkErr(err)

	// 还要再看看
	if len(c.Data) != 0 && checkCoupon == false {
		checkCoupon = true
		log.Println("您有优惠券可以使用喵～")
		config.Buy.CouponToken = c.Data[0].CouponToken
		writeConfig()
	} else if len(c.Data) == 0 && checkCoupon == false {
		checkCoupon = true
		if config.Buy.CouponToken != "" {
			config.Buy.CouponToken = ""
			writeConfig()
		}
	}
}

// 创建订单
func create() {
Loop:
	for {
		// 1s 循环一次
		task := time.NewTimer(1 * time.Second)

		headers := map[string]string{
			"native_api_from":    "h5",
			"refer":              fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1", config.Buy.ItemId),
			"env":                "prod",
			"app-key":            "android64",
			"user-agent":         appUserAgent,
			"x-bili-trace-id":    genTraceID(),
			"x-bili-aurora-eid":  config.Bili.XBiliAuroraEid,
			"x-bili-mid":         config.Cookies.DedeUserID,
			"x-bili-aurora-zone": "",
			"bili-bridge-engine": "cronet",
		}

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
			"ts":           strconv.FormatInt(time.Now().Unix(), 10),
		}

		//payload := CreatePayload{
		//	AccessKey:   config.AccessKey,
		//	AddMonth:    "-1",
		//	Appkey:      "1d8b6e7d45233436",
		//	BuyNum:      config.Buy.BuyNum,
		//	CouponToken: config.Buy.CouponToken,
		//	Csrf:        config.Cookies.BiliJct,
		//	Currency:    "bp",
		//	DisableRcmd: "0",
		//	ItemId:      config.Buy.ItemId,
		//	Platform:    config.Buy.Device,
		//	Statistics:  statistics,
		//	Ts:          strconv.FormatInt(time.Now().Unix(), 10),
		//}

		//t := reflect.TypeOf(payload)
		//v := reflect.ValueOf(payload)
		//
		//va := url.Values{}
		//
		//for k := 0; k < t.NumField(); k++ {
		//	if t.Field(k).Tag.Get("json") == "sign" {
		//		continue
		//	}
		//	va.Add(t.Field(k).Tag.Get("json"), v.Field(k).String())
		//}
		//
		//body, _ := url.QueryUnescape(va.Encode())
		//sign := strMd5(fmt.Sprintf("%v%v", body, "560c52ccd288fed045859ed18bffd973"))
		//payload.Sign = sign
		//
		//j, err := json.Marshal(payload)
		//fmt.Println(j)

		sign := appSign(data)
		data["sign"] = sign
		creates := &Create{}

		r, err := client.R().
			SetHeaders(headers).
			SetFormData(data).
			//SetBody(j).
			SetResult(creates).
			EnableTrace().
			Post("/garb/v2/trade/create")

		checkErr(err)

		//fmt.Println(r.Request.Header)
		//fmt.Println()
		//fmt.Println(r.Request.URL)
		//fmt.Println()
		//fmt.Println(r.Request.Body)
		//fmt.Println()
		//fmt.Println(r.Request.FormData)

		log.Printf("本次请求用时: %v，时间差: %v ms，时间调整:%v ms.", r.Request.TraceInfo().TotalTime, diffTime, config.Buy.TimeBefore)

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
			}
			break Loop
		case -400:
			log.Fatalln(r)
		case -403: //号被封了
			log.Fatalln("您已被封禁.")
		case 26102: //商品不存在，可能是未到抢购时间，立即重新执行
			errorTime += 1
			if errorTime >= 5 {
				log.Fatalln("失败次数已达到五次，退出执行...")
			}
			log.Println(r)
			go coupon()
			goto Loop
		case 26106: //购买数量达到上限
			log.Fatalln(r)
		case 26120: //请求频率过快，等一下执行，需要测试是否延迟执行
			fastTime++
			if fastTime >= 5 {
				log.Println(r)
				log.Fatalln("请求频率过快!失败次数已达到五次，退出执行...")
			}
			log.Println(r)
			go coupon()
		case 26113: //号被封了
			//log.Fatalln("当前设备/账号/环境存在风险，暂时无法下单.")
			log.Fatalln(r)
		case 26134: //当前抢购人数过多, 风控等级似乎比 26135 高 (好久没遇到了, 草)
			errorTime += 1
			if errorTime >= 5 {
				log.Println(r)
				log.Fatalln("失败次数已达到五次，退出执行...")
			}
			log.Println(r)
			go coupon()
		case 26135: //当前抢购人数过多，失败四次或者锁四秒后能够购买
			errorTime += 1
			if errorTime >= 5 {
				log.Println(r)
				log.Fatalln("失败次数已达到五次，退出执行...")
			}
			log.Println(r)
			go coupon()
		case 69949: //老风控代码，疑似封锁设备
			errorTime += 1
			log.Println(r)
			log.Println("已触发69949.")
			go coupon()
			if errorTime >= 5 {
				log.Fatalln("失败次数已达到五次，退出执行...")
			}
		default:
			errorTime += 1
			log.Println(r)
			go coupon()
			if errorTime >= 5 {
				log.Fatalln("失败次数已达到五次，退出执行...")
			}
		}
		<-task.C
	}
}

// 跟踪订单
func tradeQuery() {
Loop:
	for {
		task := time.NewTimer(500 * time.Millisecond)

		headers := map[string]string{
			"Content-Type":       "application/json, text/plain, */*",
			"native_api_from":    "h5",
			"refer":              fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1", config.Buy.ItemId),
			"env":                "prod",
			"app-key":            "android64",
			"user-agent":         appUserAgent,
			"x-bili-trace-id":    genTraceID(),
			"x-bili-aurora-eid":  config.Bili.XBiliAuroraEid,
			"x-bili-mid":         config.Cookies.DedeUserID,
			"x-bili-aurora-zone": "",
			"bili-bridge-engine": "cronet",
		}

		params := map[string]string{
			"access_key":   config.AccessKey,
			"appkey":       "1d8b6e7d45233436",
			"csrf":         config.Cookies.BiliJct,
			"disable_rcmd": "0",
			"order_id":     orderId,
			"statistics":   statistics,
			"ts":           strconv.FormatInt(time.Now().Unix(), 10), // 需要测试是否需要续一秒
		}

		sign := appSign(params)
		params["sign"] = sign
		query := &Query{}

		r, err := client.R().
			SetHeaders(headers).
			SetQueryParams(params).
			SetResult(query).
			Get("/garb/trade/query")

		checkErr(err)

		if query.Code == 0 {
			switch query.Data.State {
			case "paid":
				log.Println("已成功支付.")
				break Loop
			case "paying":
				log.Println("支付中，请稍候...")
			case "cancel_failed":
				log.Println("订单状态: 取消失败, 可能会购买失败喵~")
			default:
				errorTime += 1
				log.Println(r)
				if errorTime >= 5 {
					log.Fatalln("失败次数已达到五次，退出执行...")
				}
			}
		} else {
			errorTime += 1
			log.Println(r)
			if errorTime >= 5 {
				log.Fatalln("失败次数已达到五次，退出执行...")
			}
		}
		<-task.C
	}
}

// 钱包余额
func wallet() {
	headers := map[string]string{
		"Content-Type":       "application/json, text/plain, */*",
		"native_api_from":    "h5",
		"refer":              fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1", config.Buy.ItemId),
		"env":                "prod",
		"app-key":            "android64",
		"user-agent":         appUserAgent,
		"x-bili-trace-id":    genTraceID(),
		"x-bili-aurora-eid":  config.Bili.XBiliAuroraEid,
		"x-bili-mid":         config.Cookies.DedeUserID,
		"x-bili-aurora-zone": "",
		"bili-bridge-engine": "cronet",
	}

	params := map[string]string{
		"platform": "android",
	}

	response := &Wallet{}

	_, err := client.R().
		SetHeaders(headers).
		SetQueryParams(params).
		SetResult(response).
		Get("/garb/user/wallet?platform")

	checkErr(err)

	log.Printf("购买完成! 余额: %v.", response.Data.BcoinBalance)
}

// 查询编号等信息
func suitAsset() {
	headers := map[string]string{
		"Content-Type":       "application/json, text/plain, */*",
		"native_api_from":    "h5",
		"refer":              fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1", config.Buy.ItemId),
		"env":                "prod",
		"app-key":            "android64",
		"user-agent":         appUserAgent,
		"x-bili-trace-id":    genTraceID(),
		"x-bili-aurora-eid":  config.Bili.XBiliAuroraEid,
		"x-bili-mid":         config.Cookies.DedeUserID,
		"x-bili-aurora-zone": "",
		"bili-bridge-engine": "cronet",
	}

	params := map[string]string{
		"item_id": config.Buy.ItemId,
		"part":    "suit",
		"trial":   "0",
	}

	response := &SuitAsset{}

	_, err := client.R().
		SetHeaders(headers).
		SetQueryParams(params).
		SetResult(response).
		Get("garb/user/suit/asset")

	checkErr(err)

	log.Printf("名称: %v 编号: %v.", itemName, response.Data.Fan.Number)
	if response.Data.Fan.Number <= 10 {
		log.Println("恭喜拿下前 10 喵～")
	}
}
