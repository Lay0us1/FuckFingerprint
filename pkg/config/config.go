package config

/**
  @author: yhy
  @since: 2022/9/19
  @desc: //TODO
**/

var Proxy = ""

var DefaultHeader = map[string]string{
	"Accept-Language": "zh,zh-TW;q=0.9,en-US;q=0.8,en;q=0.7,zh-CN;q=0.6",
	"User-agent":      "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36",
	"Cookie":          "rememberMe=3",
	"Accept":          "text/html,*/*;",
}

const EHoleFingerDataOnline = "https://raw.githubusercontent.com/yhy0/FuckFingerprint/main/fingerPrints/eHoleFinger.json"
const LocalFingerDataOnline = "https://raw.githubusercontent.com/yhy0/FuckFingerprint/main/fingerPrints/localFinger.json"
const AfrogFingerDataOnline = "https://raw.githubusercontent.com/yhy0/FuckFingerprint/main/fingerPrints/web_fingerprint_v3.json"
