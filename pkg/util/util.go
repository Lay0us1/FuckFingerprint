package util

import (
	"bytes"
	"encoding/base64"
	"github.com/spaolacci/murmur3"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
)

/**
  @author: yhy
  @since: 2022/9/19
  @desc: //TODO
**/

// RemoveDuplicateElement  数组去重
func RemoveDuplicateElement(strs []string) []string {
	var result []string
	for _, item := range strs {
		item = strings.TrimSpace(item)
		if item != "" && !In(item, result) {
			item = strings.TrimSpace(item)
			result = append(result, item)
		}
	}
	return result
}

// In 判断一个字符串是否在另一个字符数组里面，存在返回true
// target 中是否包含 strs 中的某一个
// 这个是 target 范围大， element只是一个子串
func In(target string, strs []string) bool {
	for _, element := range strs {
		element = strings.TrimSpace(element)
		if Contains(target, element) {
			return true
		}
	}
	return false
}

func Contains(stringA, stringB string) bool {
	// stringA 原始串，stringB 要查找的字串
	return strings.Contains(strings.ToLower(stringA), strings.ToLower(stringB))
}

func FaviconHash(data []byte) int32 {
	stdBase64 := base64.StdEncoding.EncodeToString(data)
	stdBase64 = InsertInto(stdBase64, 76, '\n')
	hasher := murmur3.New32WithSeed(0)
	hasher.Write([]byte(stdBase64))
	return int32(hasher.Sum32())
}

func InsertInto(s string, interval int, sep rune) string {
	var buffer bytes.Buffer
	before := interval - 1
	last := len(s) - 1
	for i, char := range s {
		buffer.WriteRune(char)
		if i%interval == before && i != last {
			buffer.WriteRune(sep)
		}
	}
	buffer.WriteRune(sep)
	return buffer.String()
}

// Decodegbk converts GBK to UTF-8
func Decodegbk(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func DecodeKorean(s []byte) ([]byte, error) {
	koreanDecoder := korean.EUCKR.NewDecoder()
	return koreanDecoder.Bytes(s)
}
