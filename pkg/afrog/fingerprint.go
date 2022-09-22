package afrog

import (
	_ "embed"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/yhy0/FuckFingerprint/fingerPrints"
	"github.com/yhy0/FuckFingerprint/pkg/config"
	"github.com/yhy0/FuckFingerprint/pkg/honeypot"
	"github.com/yhy0/FuckFingerprint/pkg/logging"
	"github.com/yhy0/FuckFingerprint/pkg/util"
	"github.com/yhy0/FuckFingerprint/pkg/vscan"
	"strings"
)

// reference https://github.com/0x727/FingerprintHub

type Result struct {
	Url        string // 网址
	StatusCode string // 状态码
	Title      string // 标题
	Name       string // 指纹
}

func Run(url string) (res []string, title string, err error) {
	resp, err := util.HttpRequset(url, "GET", "", false, config.DefaultHeader)

	if err != nil {
		logging.Logger.Error("HttpRequset err: ", err)
		return nil, "", err
	}

	if honeypot.FuckHoneypot(resp.Body) {
		res = append(res, "蜜罐")
	}

	title = util.ExtractTitle(resp)
	res = vscan.FingerScan(resp.Header, resp.Body, title, url)

	// todo 这里其实应该在上次都自动跳转的，但由于指纹不统一，有的是跳转后识别，有的跳转前就能识别，所以这里多次验证一下
	if len(res) == 0 || resp.StatusCode == 302 || resp.Header.Get("Location") != "" {
		respTmp, err := util.HttpRequset(url, "GET", "", true, config.DefaultHeader)

		if err == nil {
			resp = respTmp
			res = append(res, vscan.FingerScan(resp.Header, resp.Body, "", url)...)
			title = util.ExtractTitle(resp)
		}

	}

	url = strings.TrimRight(url, "/")

	for k, v := range fingerPrints.AfrogFingerPrintMap {
		target := url
		if k != "/" {
			target = url + k
			resp, err = util.HttpRequset(target, "GET", "", false, nil)
			if err != nil {
				continue
			}
		}

		for _, fingerPrint := range v {
			if strings.ToLower(fingerPrint.RequestMethod) == "get" && len(fingerPrint.RequestHeaders) == 0 {
				if matching(resp, fingerPrint) {
					res = append(res, fingerPrint.Name)
				}
			} else if strings.ToLower(fingerPrint.RequestMethod) == "get" && len(fingerPrint.RequestHeaders) != 0 {
				resp, err = util.HttpRequset(target, "GET", "", false, nil)
				if err != nil {
					continue
				}

				if matching(resp, fingerPrint) {
					res = append(res, fingerPrint.Name)
				}

			} else if strings.ToLower(fingerPrint.RequestMethod) == "post" && len(fingerPrint.RequestHeaders) == 0 {
				resp, err = util.HttpRequset(target, strings.ToUpper(fingerPrint.RequestMethod), fingerPrint.RequestData, false, nil)
				if err != nil {
					continue
				}
				if matching(resp, fingerPrint) {
					res = append(res, fingerPrint.Name)
				}

			} else {
				resp, err = util.HttpRequset(target, strings.ToUpper(fingerPrint.RequestMethod), fingerPrint.RequestData, false, fingerPrint.RequestHeaders)
				if err != nil {
					continue
				}
				if matching(resp, fingerPrint) {
					res = append(res, fingerPrint.Name)
				}
			}
		}

	}

	res = util.RemoveDuplicateElement(res)
	logging.Logger.Infoln(fmt.Sprintf("%s %s %v", aurora.Cyan(url), aurora.Green(title), aurora.Red(res)))

	return res, title, nil
}

// 注意 header 请书写规范
func matching(resp *util.Response, f fingerPrints.AfrogFingerPrint) bool {
	flag := false
	hflag := true
	if len(f.Headers) > 0 {
		hflag = false
		for k, h := range f.Headers {
			if len(resp.Header[k]) == 0 {
				hflag = false
				break
			}
			if len(resp.Header[k]) > 0 {
				if !strings.Contains(resp.Header[k][0], h) {
					hflag = false
					break
				}
				hflag = true
			}
		}
	}
	if len(f.Headers) > 0 && hflag {
		flag = true
	}

	kflag := true
	if len(f.Keyword) > 0 {
		kflag = false
		for _, k := range f.Keyword {
			if !strings.Contains(resp.Body, k) {
				kflag = false
				break
			}
			kflag = true
		}
	}
	if len(f.Keyword) > 0 && kflag {
		flag = true
	}

	if flag {
		return true
	}

	return false
}
