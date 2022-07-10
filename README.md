<h1 align="center">Bilibili 装扮抢购脚本</h1>
<p align="center">基于 GO 语言编写的抢购脚本，包含扫码登录与自动匹配 UA 功能，无需抓包</p>
<p align="center">
<a href="https://github.com/KaguraMika/bili-suit-v3/releases/latest"><img src="https://img.shields.io/badge/LICENSE-GPL--3.0-blue" alt="License"></a>
<a href="https://github.com/go-resty/resty/releases/latest"><img src="https://img.shields.io/badge/VERSION-3.1.0-brightgreen" alt="Release Version"></a>
</p>
<br><br>

## 简介
由于个人原因，此版本暂停开发（要是有人接手最好了喵~）

尚未支持 IOS (懒～)

傻瓜式，扫码登录，自动填写配置文件 (Mika真是太优雅了喵～)

尽量模拟了APP端抢购过程（风控 -1）

~~已通过大量测试 ✅~~

## 更新
* 切换至 APP 端API
* 时间校准切换至 NTP
* ~~提前生成表单与 Sign 值~~

## 使用方法
1. 下载并解压 `Release` 中对应的文件，哪个平台就用哪个
2. 填写 `config.json` 中的 `item_id` （装扮ID）
3. 运行脚本: 在终端中运行 `./bili-suit-tool` (windows 运行 `./bili-suit-tool.exe`)
4. 按照提示，在 APP 中访问 `https://api.bilibili.com/client_info`,
   并将所有信息复制并填入
5. 等待开售

## 小提示：

* 使用 `-c` 可指定配置文件，例如: `./bili-suit-tool -c /etc/bili/1.json`
* 使用 `-i` 可指定装扮 ID，例如: `./bili-suit-tool -i 114514 `
* 使用 `-b` 可指定购买数量，例如: `./bili-suit-tool -b 19 `
* 使用 `-t` 可设置下单延迟, 正数提前，负数延后，例如: `./bili-suit-tool -t -100 `
* `cookies` 必要参数留空可使用扫码登录
* `bp_enough` 为 `true` 时开启 b币余额校验，b币余额不足时不下单，为 `false` 将会忽略校验直接下单

## 配置文件

**config.json**

```
{
  "accessKey": "",
  "buy": {
    "bp_enough": true,
    "buy_num": "1",
    "coupon_token": "",
    "device": "android",
    "item_id": "",
    "time_before": 0
  },
  "bili": {
    "app_version": "",
    "app_build_version": "",
    "app_inner_version": "",
    "tv_version": "1.5.0",
    "x-bili-aurora-eid": ""
  },
  "cookies": {
    "SESSDATA": "",
    "bili_jct": "",
    "DedeUserID": "",
    "DedeUserID__ckMd5": "",
    "sid": "",
    "Buvid": ""
  },
  "phone": {
    "android_version": "",
    "android_api_level": "",
    "build": "",
    "chrome_version": "",
    "device_name": ""
  }
}
```

# Author
[**超急玛丽**](https://space.bilibili.com/24924450)  
[**恋利普贝当**](https://space.bilibili.com/2932835)

