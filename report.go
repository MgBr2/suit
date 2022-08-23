package main

import (
	"encoding/json"
	"io"
	"net/http"
)

/*
	1. 目前没有直接证据表明风控与日志上报有关，也没有直接证据表面也日志上报无关
	2. 两种日志上报已存在许久，但之前不上报时也能成功购买
*/

func spi() {
	if config.Device.B3 == "" || config.Device.B4 == "" {
		s := &Spi{}

		req, err := http.NewRequest("GET", "https://api.bilibili.com/x/frontend/finger/spi", nil)
		checkErr(err)

		req.Header.Set("x-requested-with", "tv.danmaku.bili")
		req.Header.Set("user-agent", appUserAgent)

		resp, err := client.Do(req)
		checkErr(err)

		body, err := io.ReadAll(resp.Body)
		checkErr(err)

		defer func(Body io.ReadCloser) {
			err = Body.Close()
			checkErr(err)
		}(resp.Body)

		err = json.Unmarshal(body, s)
		checkErr(err)

		config.Device.B3 = s.Data.B3
		config.Device.B4 = s.Data.B4

		writeConfig()
	}
}

func gatewayReport() {

}

func logReport() {
	// https://www.bilibili.com/log
}

func apmReport() {
	// https://api.bilibili.com/open/monitor/apm/report
}
