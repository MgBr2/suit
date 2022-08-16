package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	navs    = &Navs{}
	details = &Details{}
	coupons = &Coupon{}
)

// 登录状态
func nav() {
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/web-interface/nav", nil)
	checkErr(err)

	req = commonHeaders(req, "https://www.bilibili.com/h5/mall/home?navhide=1&from=myservice&native.theme=1")

	u := url.Values{}

	u.Add("access_key", config.AccessKey)
	u.Add("appKey", "1d8b6e7d45233436")
	u.Add("csrf", config.Cookies.BiliJct)
	u.Add("disable_rcmd", "0")
	u.Add("statistics", statistics)
	u.Add("ts", strconv.FormatInt(time.Now().Unix(), 10))

	// 拼接 Sign, 并格式化字符串
	params, err := url.QueryUnescape(fmt.Sprintf("%v&sign=%v", u.Encode(), appSign(u)))
	checkErr(err)

	// 注入灵魂
	req.URL.RawQuery = params

	// 执行请求
	resp, err := client.Do(req)
	checkErr(err)

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	checkErr(err)

	// 延迟关闭
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		checkErr(err)
	}(resp.Body)

	// 解析 JSON
	err = json.Unmarshal(body, navs)
	checkErr(err)
}

// popup, 未知用途
func popup() {
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/garb/popup", nil)
	checkErr(err)

	req = commonHeaders(req, "https://www.bilibili.com/h5/mall/home?navhide=1&from=myservice&native.theme=1")

	u := url.Values{}

	u.Add("access_key", config.AccessKey)
	u.Add("appKey", "1d8b6e7d45233436")
	u.Add("csrf", config.Cookies.BiliJct)
	u.Add("disable_rcmd", "0")
	u.Add("statistics", statistics)
	u.Add("ts", strconv.FormatInt(time.Now().Unix(), 10))

	// 拼接 Sign, 并格式化字符串
	params, err := url.QueryUnescape(fmt.Sprintf("%v&sign=%v", u.Encode(), appSign(u)))
	checkErr(err)

	// 注入灵魂
	req.URL.RawQuery = params

	// 执行请求
	_, err = client.Do(req)
	checkErr(err)
}

// 装扮信息
func detail() {
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/garb/v2/mall/suit/detail", nil)
	checkErr(err)

	req = commonHeaders(req, fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1&f_source=shop&from=feed.card", config.Buy.ItemId))

	u := url.Values{}

	u.Add("access_key", config.AccessKey)
	u.Add("appKey", "1d8b6e7d45233436")
	u.Add("csrf", config.Cookies.BiliJct)
	u.Add("disable_rcmd", "0")
	u.Add("from", "")
	u.Add("from_id", "")
	u.Add("item_id", config.Buy.ItemId)
	u.Add("part", "suit")
	u.Add("statistics", statistics)
	u.Add("ts", strconv.FormatInt(time.Now().Unix(), 10))

	// 拼接 Sign, 并格式化字符串
	params, err := url.QueryUnescape(fmt.Sprintf("%v&sign=%v", u.Encode(), appSign(u)))
	checkErr(err)

	// 注入灵魂
	req.URL.RawQuery = params

	// 执行请求
	resp, err := client.Do(req)
	checkErr(err)

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	checkErr(err)

	// 延迟关闭
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		checkErr(err)
	}(resp.Body)

	// 解析 JSON
	err = json.Unmarshal(body, details)
	checkErr(err)
}

// 拥有的装扮信息
func asset() {
	assests := &Asset{}

	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/garb/user/asset", nil)
	checkErr(err)

	req = commonHeaders(req, fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1&f_source=shop&from=feed.card", config.Buy.ItemId))

	u := url.Values{}

	u.Add("access_key", config.AccessKey)
	u.Add("appKey", "1d8b6e7d45233436")
	u.Add("csrf", config.Cookies.BiliJct)
	u.Add("disable_rcmd", "0")
	u.Add("item_id", config.Buy.ItemId)
	u.Add("part", "suit")
	u.Add("statistics", statistics)
	u.Add("ts", strconv.FormatInt(time.Now().Unix(), 10))

	// 拼接 Sign, 并格式化字符串
	params, err := url.QueryUnescape(fmt.Sprintf("%v&sign=%v", u.Encode(), appSign(u)))
	checkErr(err)

	// 注入灵魂
	req.URL.RawQuery = params

	// 执行请求
	resp, err := client.Do(req)
	checkErr(err)

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	checkErr(err)

	// 延迟关闭
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		checkErr(err)
	}(resp.Body)

	// 解析 JSON
	err = json.Unmarshal(body, assests)
	checkErr(err)
}

// 装扮排名
func rank() {
	ranks := &Rank{}

	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/garb/rank/fan/recent", nil)
	checkErr(err)

	req = commonHeaders(req, fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1&f_source=shop&from=feed.card", config.Buy.ItemId))
	req = commonParams(req)

	// 执行请求
	resp, err := client.Do(req)
	checkErr(err)

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	checkErr(err)

	// 延迟关闭
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		checkErr(err)
	}(resp.Body)

	// 解析 JSON
	err = json.Unmarshal(body, ranks)
	checkErr(err)

	rankInfo = ranks
}

// 关于这个装扮的订单情况
func stat() {
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/garb/order/user/stat", nil)
	checkErr(err)

	req = commonHeaders(req, fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1&f_source=shop&from=feed.card", config.Buy.ItemId))
	req = commonParams(req)

	// 执行请求
	_, err = client.Do(req)
	checkErr(err)
}

// 预约人数及预约状态
func state() {
	reserveInfo := &Reserve{}

	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/garb/user/reserve/state", nil)
	checkErr(err)

	req = commonHeaders(req, fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1&f_source=shop&from=feed.card", config.Buy.ItemId))

	u := url.Values{}

	u.Add("access_key", config.AccessKey)
	u.Add("appKey", "1d8b6e7d45233436")
	u.Add("csrf", config.Cookies.BiliJct)
	u.Add("disable_rcmd", "0")
	u.Add("item_id", config.Buy.ItemId)
	u.Add("part", "suit")
	u.Add("statistics", statistics)
	u.Add("ts", strconv.FormatInt(time.Now().Unix(), 10))

	// 拼接 Sign, 并格式化字符串
	params, err := url.QueryUnescape(fmt.Sprintf("%v&sign=%v", u.Encode(), appSign(u)))
	checkErr(err)

	// 注入灵魂
	req.URL.RawQuery = params

	// 执行请求
	resp, err := client.Do(req)
	checkErr(err)

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	checkErr(err)

	// 延迟关闭
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		checkErr(err)
	}(resp.Body)

	// 解析 JSON
	err = json.Unmarshal(body, reserveInfo)
	checkErr(err)

	log.Printf("当前装扮预约人数: %v.", reserveInfo.Data.ReserveCount)

	if reserveInfo.Data.Reserved == false && reserveInfo.Data.ReserveState == true {
		if config.Buy.Reserve == true {
			log.Println("你还没有预约这个装扮哦，Mika 这就帮你预约喵～")
			reserve()
		} else {
			log.Println("你还没有预约这个装扮哦，Mika 建议你预约一下喵～")
		}
	} else if reserveInfo.Data.Reserved == true {
		log.Println("当前装扮已经预约了喵～")
	}
}

// 执行预约请求
func reserve() {
	reserveInfo := &Reserve{}
	u := url.Values{}

	u.Add("access_key", config.AccessKey)
	u.Add("appKey", "1d8b6e7d45233436")
	u.Add("csrf", config.Cookies.BiliJct)
	u.Add("disable_rcmd", "0")
	u.Add("item_id", config.Buy.ItemId)
	u.Add("part", "suit")
	u.Add("statistics", statistics)
	u.Add("ts", strconv.FormatInt(time.Now().Unix(), 10))

	data, err := url.QueryUnescape(fmt.Sprintf("%v&sign=%v", u.Encode(), appSign(u)))
	checkErr(err)

	req, err := http.NewRequest("POST", "https://api.bilibili.com/x/garb/user/reserve", strings.NewReader(data))
	checkErr(err)

	req = commonHeaders(req, fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1&f_source=shop&from=feed.card", config.Buy.ItemId))

	// 执行请求
	resp, err := client.Do(req)
	checkErr(err)

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	checkErr(err)

	// 延迟关闭
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		checkErr(err)
	}(resp.Body)

	// 解析 JSON
	err = json.Unmarshal(body, reserveInfo)
	checkErr(err)

	if reserveInfo.Code != 0 {
		go log.Println("预约失败了喵！")
	}
}

// 优惠券
func coupon() {
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/garb/coupon/usable", nil)
	checkErr(err)

	req = commonHeaders(req, fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1&f_source=shop&from=feed.card", config.Buy.ItemId))
	req = commonParams(req)

	// 执行请求
	resp, err := client.Do(req)
	checkErr(err)

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	checkErr(err)

	// 延迟关闭
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		checkErr(err)
	}(resp.Body)

	// 解析 JSON
	err = json.Unmarshal(body, coupons)
	checkErr(err)
}

// 创建订单
func create() {
	creates := &Create{}

Loop:
	for {
		// 1s 循环一次
		task := time.NewTimer(1 * time.Second)
		t := time.Now()
		u := url.Values{}

		u.Add("access_key", config.AccessKey)
		u.Add("add_month", "-1")
		u.Add("appkey", "1d8b6e7d45233436")
		u.Add("buy_num", config.Buy.BuyNum)
		u.Add("coupon_token", config.Buy.CouponToken)
		u.Add("csrf", config.Cookies.BiliJct)
		u.Add("currency", "bp")
		u.Add("disable_rcmd", "0")
		u.Add("f_source", "shop")
		u.Add("from", "feed.card")
		u.Add("from_id", "")
		u.Add("item_id", config.Buy.ItemId)
		u.Add("platform", config.Buy.Device)
		u.Add("statistics", statistics)
		u.Add("ts", strconv.FormatInt(t.Unix(), 10))

		data, err := url.QueryUnescape(fmt.Sprintf("%v&sign=%v", u.Encode(), appSign(u)))
		checkErr(err)

		req, err := http.NewRequest("POST", "https://api.bilibili.com/x/garb/v2/trade/create", strings.NewReader(data))
		checkErr(err)

		req.Header.Set("accept", "application/json, text/plain, */*")
		req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=utf-8")
		req.Header.Set("native_api_from", "h5")
		req.Header.Set("refer", fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1&f_source=shop&from=feed.card", config.Buy.ItemId))
		req.Header.Set("env", "prod")
		req.Header.Set("app-key", "android64")
		req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
		req.Header.Set("x-bili-trace-id", genTraceID())
		req.Header.Set("x-bili-aurora-eid", "")
		req.Header.Set("x-bili-mid", "")
		req.Header.Set("x-bili-aurora-zone", "")
		req.Header.Set("bili-bridge-engine", "cronet")

		resp, err := client.Do(req)
		checkErr(err)

		elapsed := time.Since(t)
		log.Printf("本次请求用时: %v，时间差: %v ms，时间调整:%v ms.", elapsed, diffTime, config.Buy.TimeBefore)

		body, err := io.ReadAll(resp.Body)
		checkErr(err)
		r := string(body)

		err = json.Unmarshal(body, creates)
		checkErr(err)

		switch creates.Code {
		case 0: // 这里好像有问题，还需要再看看
			if creates.Data.BpEnough == -1 {
				log.Println(resp)
				log.Fatalln("余额不足.")
			}
			// 订单号
			orderId = creates.Data.OrderId
			if creates.Data.State != "paying" {
				log.Println(resp)
			}
			break Loop
		case -3:
			log.Println(r)
			log.Fatalln("请将此问题上报给 Mika.")
		case -101:
			log.Fatalln(r)
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

		err = resp.Body.Close()
		checkErr(err)

		<-task.C
	}
}

// 跟踪订单
func tradeQuery() {
	query := &Query{}

Loop:
	for {
		task := time.NewTimer(500 * time.Millisecond)

		req, err := http.NewRequest("GET", "https://api.bilibili.com/x/garb/trade/query", nil)
		checkErr(err)

		req = commonHeaders(req, fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/detail?id=%v&navhide=1&f_source=shop&from=feed.card", config.Buy.ItemId))

		u := url.Values{}

		u.Add("access_key", config.AccessKey)
		u.Add("appKey", "1d8b6e7d45233436")
		u.Add("csrf", config.Cookies.BiliJct)
		u.Add("disable_rcmd", "0")
		u.Add("order_id", orderId)
		u.Add("statistics", statistics)
		u.Add("ts", strconv.FormatInt(time.Now().Unix(), 10))

		// 拼接 Sign, 并格式化字符串
		params, err := url.QueryUnescape(fmt.Sprintf("%v&sign=%v", u.Encode(), appSign(u)))
		checkErr(err)

		// 注入灵魂
		req.URL.RawQuery = params

		// 执行请求
		resp, err := client.Do(req)
		checkErr(err)

		// 读取响应
		body, err := io.ReadAll(resp.Body)
		checkErr(err)
		r := string(body)

		// 解析 JSON
		err = json.Unmarshal(body, query)
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

		err = resp.Body.Close()
		checkErr(err)

		<-task.C
	}
}

// 钱包余额
func wallet() {
	myWallet := &Wallet{}

	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/garb/user/wallet", nil)
	checkErr(err)

	req = commonHeaders(req, fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/manage/%v?navhide=1&native.theme=1", config.Buy.ItemId))

	u := url.Values{}

	u.Add("platform", "android")

	// 注入灵魂
	req.URL.RawQuery = u.Encode()

	// 执行请求
	resp, err := client.Do(req)
	checkErr(err)

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	checkErr(err)

	// 延迟关闭
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		checkErr(err)
	}(resp.Body)

	// 解析 JSON
	err = json.Unmarshal(body, myWallet)
	checkErr(err)

	log.Printf("购买完成! 余额: %v.", myWallet.Data.BcoinBalance)
}

// 查询编号等信息
func suitAsset() {
	mySuit := &SuitAsset{}

	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/garb/user/wallet", nil)
	checkErr(err)

	req = commonHeaders(req, fmt.Sprintf("https://www.bilibili.com/h5/mall/suit/manage/%v?navhide=1&native.theme=1", config.Buy.ItemId))

	u := url.Values{}

	u.Add("item_id", config.Buy.ItemId)
	u.Add("part", "suit")
	u.Add("trial", "0")

	// 注入灵魂
	req.URL.RawQuery = u.Encode()

	// 执行请求
	resp, err := client.Do(req)
	checkErr(err)

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	checkErr(err)

	// 延迟关闭
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		checkErr(err)
	}(resp.Body)

	// 解析 JSON
	err = json.Unmarshal(body, mySuit)
	checkErr(err)

	log.Printf("名称: %v 编号: %v.", itemName, mySuit.Data.Fan.Number)
	if mySuit.Data.Fan.Number <= 10 {
		log.Println("恭喜拿下前 10 喵～")
	}
}
