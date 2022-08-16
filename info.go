package main

type AuthCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Url      string `json:"url"`
		AuthCode string `json:"auth_code"`
	} `json:"data"`
}

type Config struct {
	AccessKey string `json:"accessKey"`
	Buy       struct {
		BpEnough    bool   `json:"bp_enough"`
		BuyNum      string `json:"buy_num"`
		CouponToken string `json:"coupon_token"`
		Device      string `json:"device"`
		ItemId      string `json:"item_id"`
		TimeBefore  int    `json:"time_before"`
		Reserve     bool   `json:"reserve"`
	} `json:"buy"`
	Bili struct {
		AppVersion      string `json:"app_version"`
		AppBuildVersion string `json:"app_build_version"`
		AppInnerVersion string `json:"app_inner_version"`
		TvVersion       string `json:"tv_version"`
		XBiliAuroraEid  string `json:"x-bili-aurora-eid"`
	} `json:"bili"`
	Cookies struct {
		SESSDATA        string `json:"SESSDATA"`
		BiliJct         string `json:"bili_jct"`
		DedeUserID      string `json:"DedeUserID"`
		DedeUserIDCkMd5 string `json:"DedeUserID__ckMd5"`
		Sid             string `json:"sid"`
		Buvid           string `json:"Buvid"`
	} `json:"cookies"`
	Phone struct {
		AndroidVersion  string `json:"android_version"`
		AndroidApiLevel string `json:"android_api_level"`
		Build           string `json:"build"`
		ChromeVersion   string `json:"chrome_version"`
		DeviceName      string `json:"device_name"`
	} `json:"phone"`
}

type QrcodePoll struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		IsNew        bool   `json:"is_new"`
		Mid          int    `json:"mid"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
		TokenInfo    struct {
			Mid          int    `json:"mid"`
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
			ExpiresIn    int    `json:"expires_in"`
		} `json:"token_info"`
		CookieInfo struct {
			Cookies []struct {
				Name     string `json:"name"`
				Value    string `json:"value"`
				HttpOnly int    `json:"http_only"`
				Expires  int    `json:"expires"`
			} `json:"cookies"`
			Domains []string `json:"domains"`
		} `json:"cookie_info"`
		Sso []string `json:"sso"`
	} `json:"data"`
}

type Static struct {
	AppId    int    `json:"appId"`
	Platform int    `json:"platform"`
	Version  string `json:"version"`
	Abtest   string `json:"abtest"`
}

type Now struct {
	Data struct {
		Now int64 `json:"now"`
	} `json:"data"`
}

type Navs struct {
	Code int `json:"code"`
	Data struct {
		Wallet struct {
			BcoinBalance float64 `json:"bcoin_balance"`
		} `json:"wallet"`
		Uname string `json:"uname"`
	} `json:"data"`
}

type Details struct {
	Data struct {
		Name       string `json:"name"`
		Properties struct {
			SaleTimeBegin    string `json:"sale_time_begin"`
			SaleBpForeverRaw string `json:"sale_bp_forever_raw"`
			SaleBuyNumLimit  string `json:"sale_buy_num_limit"`
		}
		CurrentActivity struct {
			PriceBpForever float64 `json:"price_bp_forever"`
		} `json:"current_activity"`
	} `json:"data"`
}

type Asset struct {
	Data struct {
		Id   int `json:"id"`
		Item struct {
			ItemId int `json:"item_id"`
		} `json:"item"`
	} `json:"data"`
}

type Rank struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Rank []struct {
			Mid      int    `json:"mid"`
			Nickname string `json:"nickname"`
			Avatar   string `json:"avatar"`
			Number   int    `json:"number"`
		} `json:"rank"`
	} `json:"data"`
}

type Create struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		OrderId  string `json:"order_id"`
		State    string `json:"state"`
		BpEnough int    `json:"bp_enough"`
	} `json:"data"`
}

type Query struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		OrderId  string `json:"order_id"`
		Mid      int    `json:"mid"`
		Platform string `json:"platform"`
		ItemId   int    `json:"item_id"`
		PayId    string `json:"pay_id"`
		State    string `json:"state"`
	} `json:"data"`
}

type Wallet struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		BcoinBalance  float64 `json:"bcoin_balance"`
		CouponBalance int     `json:"coupon_balance"`
	} `json:"data"`
}

type SuitAsset struct {
	Data struct {
		Fan struct {
			IsFan      bool   `json:"is_fan"`
			Token      string `json:"token"`
			Number     int    `json:"number"`
			Color      string `json:"color"`
			Name       string `json:"name"`
			LuckItemId int    `json:"luck_item_id"`
			Date       string `json:"date"`
		} `json:"fan"`
	} `json:"data"`
}

type Reserve struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Reserved     bool `json:"reserved"`
		ReserveCount int  `json:"reserve_count"`
		ReserveState bool `json:"reserve_state"`
	} `json:"data"`
}

type Coupon struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    []struct {
		Title           string        `json:"title"`
		CouponToken     string        `json:"coupon_token"`
		CouponType      int           `json:"coupon_type"`
		FullAmount      int           `json:"full_amount"`
		Allowance       int           `json:"allowance"`
		Discount        int           `json:"discount"`
		LeftUsableTimes int           `json:"left_usable_times"`
		OrderId         string        `json:"order_id"`
		TimeLimit       int           `json:"time_limit"`
		StartTime       int           `json:"start_time"`
		ExpireTime      int           `json:"expire_time"`
		Explain         string        `json:"explain"`
		LimitInfo       string        `json:"limit_info"`
		LimitRule       interface{}   `json:"limit_rule"`
		ItemInfo        interface{}   `json:"item_info"`
		UsedItems       []interface{} `json:"used_items"`
		State           int           `json:"state"`
		PriorityItem    int           `json:"priority_item"`
	} `json:"data"`
}

type CreateBody struct {
	AccessKey   string `json:"access_key"`
	AddMonth    string `json:"add_month"`
	Appkey      string `json:"appkey"`
	BuyNum      string `json:"buy_num"`
	CouponToken string `json:"coupon_token"`
	Csrf        string `json:"csrf"`
	Currency    string `json:"currency"`
	DisableRcmd string `json:"disable_rcmd"`
	From        string `json:"from"`
	FromId      string `json:"from_id"`
	ItemId      string `json:"item_id"`
	Platform    string `json:"platform"`
	Statistics  string `json:"statistics"`
	Ts          string `json:"ts"`
	Sign        string `json:"sign"`
}

//type CreatePayload struct {
//	AccessKey   string `json:"access_key"`
//	AddMonth    string `json:"add_month"`
//	Appkey      string `json:"appkey"`
//	BuyNum      string `json:"buy_num"`
//	CouponToken string `json:"coupon_token"`
//	Csrf        string `json:"csrf"`
//	Currency    string `json:"currency"`
//	DisableRcmd string `json:"disable_rcmd"`
//	ItemId      string `json:"item_id"`
//	Platform    string `json:"platform"`
//	Statistics  string `json:"statistics"`
//	Ts          string `json:"ts"`
//	Sign        string `json:"sign"`
//}
