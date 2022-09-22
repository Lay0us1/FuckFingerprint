package honeypot

import (
	"github.com/yhy0/FuckFingerprint/pkg/util"
	"regexp"
	"strings"
)

/**
  @author: yhy
  @since: 2022/9/21
  @desc: //TODO
**/

func FuckHoneypot(body string) bool {
	regex := regexp.MustCompile(`<script>\w+\s*=\s*\[(.+)\];`)

	sensitiveStr := regex.FindAllString(body, -1)

	for _, i := range sensitiveStr {
		if len(strings.Split(i, "','")) > 30 && strings.Count(i, "0x") > 30 && util.Contains(body, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/=") && util.Contains(body, "fromCharCode") {
			return true
		}
	}
	return false
}
