package vscan

import (
	"encoding/json"
	"github.com/yhy0/FuckFingerprint/fingerPrints"
	"sync"
)

func mapToJson(param map[string][]string) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

func FingerScan(headers map[string][]string, bodyString string, title string, url string) []string {
	var wg sync.WaitGroup
	wg.Add(2)
	var cms []string

	fingerprint := make(chan string)
	headersjson := mapToJson(headers)
	favhash := getfavicon(bodyString, url)

	go func() {
		for _, finp := range fingerPrints.EholeFinpx.Fingerprint {
			if finp.Location == "body" {
				if finp.Method == "keyword" {
					if iskeyword(bodyString, finp.Keyword) {
						fingerprint <- finp.Cms
					}
				}
				if finp.Method == "faviconhash" {
					if favhash == finp.Keyword[0] {
						fingerprint <- finp.Cms
					}
				}
				if finp.Method == "regular" {
					if isregular(bodyString, finp.Keyword) {
						fingerprint <- finp.Cms
					}
				}
			}
			if finp.Location == "header" {
				if finp.Method == "keyword" {
					if iskeyword(headersjson, finp.Keyword) {
						fingerprint <- finp.Cms
					}
				}
				if finp.Method == "regular" {
					if isregular(headersjson, finp.Keyword) {
						fingerprint <- finp.Cms
					}
				}
			}
			if finp.Location == "title" {
				if finp.Method == "keyword" {
					if iskeyword(title, finp.Keyword) {
						fingerprint <- finp.Cms
					}
				}
				if finp.Method == "regular" {
					if isregular(title, finp.Keyword) {
						fingerprint <- finp.Cms
					}
				}
			}
		}
		wg.Done()
	}()

	go func() {
		for _, finp := range fingerPrints.LocalFinpx.Fingerprint {
			if finp.Location == "body" {
				if finp.Method == "keyword" {
					if iskeyword(bodyString, finp.Keyword) {
						fingerprint <- finp.Cms
					}
				}
				if finp.Method == "faviconhash" {
					if favhash == finp.Keyword[0] {
						fingerprint <- finp.Cms
					}
				}
				if finp.Method == "regular" {
					if isregular(bodyString, finp.Keyword) {
						fingerprint <- finp.Cms
					}
				}
			}
			if finp.Location == "header" {
				if finp.Method == "keyword" {
					if iskeyword(headersjson, finp.Keyword) {
						fingerprint <- finp.Cms
					}
				}
				if finp.Method == "regular" {
					if isregular(headersjson, finp.Keyword) {
						fingerprint <- finp.Cms
					}
				}
			}
			if finp.Location == "title" {
				if finp.Method == "keyword" {
					if iskeyword(title, finp.Keyword) {
						fingerprint <- finp.Cms
					}
				}
				if finp.Method == "regular" {
					if isregular(title, finp.Keyword) {
						fingerprint <- finp.Cms
					}
				}
			}
		}
		wg.Done()
	}()

	// 等待关闭channel
	go func() {
		wg.Wait()
		close(fingerprint)
	}()

	for i := range fingerprint {
		cms = append(cms, i)
	}
	return cms
}
