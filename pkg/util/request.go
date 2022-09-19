package util

import (
	"crypto/tls"
	"github.com/corpix/uarand"
	"github.com/yhy0/FuckFingerprint/pkg/config"
	"io/ioutil"
	"net/http/httputil"

	"net/http"
	"net/http/cookiejar"
	"net/url"

	"strings"
	"time"
)

/**
  @author: yhy
  @since: 2022/6/1
  @desc: //TODO
**/

type Response struct {
	Status        string
	StatusCode    int
	Body          string
	RequestDump   string
	Header        http.Header
	ContentLength int
	RequestUrl    string
	Location      string
}

func HttpRequset(urlstring string, method string, postdata string, isredirect bool, headers map[string]string) (*Response, error) {
	var tr *http.Transport
	if config.Proxy != "" {
		uri, _ := url.Parse(config.Proxy)
		tr = &http.Transport{
			MaxIdleConnsPerHost: -1,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives:   true,
			Proxy:               http.ProxyURL(uri),
		}
	} else {
		tr = &http.Transport{
			MaxIdleConnsPerHost: -1,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives:   true,
		}
	}

	client := &http.Client{
		Timeout:   time.Duration(10) * time.Second,
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	if isredirect {
		jar, _ := cookiejar.New(nil)
		client = &http.Client{
			Timeout:   time.Duration(10) * time.Second,
			Transport: tr,
			Jar:       jar,
		}
	}
	req, err := http.NewRequest(strings.ToUpper(method), urlstring, strings.NewReader(postdata))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", uarand.GetRandom())
	flag := true
	for k, v := range headers {
		if k == "Content-Type" {
			flag = false
		}
		req.Header[k] = []string{v}
	}

	if flag {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	}

	requestDump, _ := httputil.DumpRequestOut(req, true)

	resp, err := client.Do(req)
	if err != nil {
		//防止空指针
		return &Response{"999", 999, "", "", nil, 0, "", ""}, err
	}
	var location string
	var respbody string
	defer resp.Body.Close()
	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		respbody = string(body)
	}
	if resplocation, err := resp.Location(); err == nil {
		location = resplocation.String()
	}
	return &Response{resp.Status, resp.StatusCode, respbody, string(requestDump), resp.Header, len(respbody), resp.Request.URL.String(), location}, nil
}
