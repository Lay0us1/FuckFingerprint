package util

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"regexp"
	"strings"
)

var (
	cutset        = "\n\t\v\f\r"
	reTitle       = regexp.MustCompile(`(?im)<\s*title.*>(.*?)<\s*/\s*title>`)
	reContentType = regexp.MustCompile(`(?im)\s*charset="(.*?)"|charset=(.*?)"\s*`)
)

// ExtractTitle from a response
func ExtractTitle(r *Response) (title string) {
	// Try to parse the DOM
	titleDom, err := getTitleWithDom(r)
	// In case of error fallback to regex
	if err != nil {
		for _, match := range reTitle.FindAllString(r.Body, -1) {
			title = match
			break
		}
	} else {
		title = renderNode(titleDom)
	}

	title = html.UnescapeString(trimTitleTags(title))

	// remove unwanted chars
	title = strings.TrimSpace(strings.Trim(title, cutset))
	title = strings.ReplaceAll(title, "\n", "")
	title = strings.ReplaceAll(title, "\r", "")

	// Non UTF-8
	if contentTypes, ok := r.Header["Content-Type"]; ok {
		contentType := strings.ToLower(strings.Join(contentTypes, ";"))

		switch {
		case In(contentType, []string{"charset=gb2312", "charset=gbk"}):
			titleUtf8, err := Decodegbk([]byte(title))
			if err != nil {
				return
			}

			return string(titleUtf8)
		case Contains(contentType, "euc-kr"):
			titleUtf8, err := DecodeKorean([]byte(title))
			if err != nil {
				return
			}
			return string(titleUtf8)
		}

		// Content-Type from head tag
		var match = reContentType.FindSubmatch([]byte(r.Body))
		var mcontentType = ""
		if len(match) != 0 {
			for i, v := range match {
				if string(v) != "" && i != 0 {
					mcontentType = string(v)
				}
			}
			mcontentType = strings.ToLower(mcontentType)
		}
		switch {
		case In(mcontentType, []string{"gb2312", "gbk"}):
			titleUtf8, err := Decodegbk([]byte(title))
			if err != nil {
				return
			}
			return string(titleUtf8)
		}
	}

	return //nolint
}

func getTitleWithDom(r *Response) (*html.Node, error) {
	var title *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "title" {
			title = node
			return
		}
		for child := node.FirstChild; child != nil && title == nil; child = child.NextSibling {
			crawler(child)
		}
	}
	htmlDoc, err := html.Parse(bytes.NewReader([]byte(r.Body)))
	if err != nil {
		return nil, err
	}
	crawler(htmlDoc)
	if title != nil {
		return title, nil
	}
	return nil, fmt.Errorf("title not found")
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n) //nolint
	return buf.String()
}

func trimTitleTags(title string) string {
	// trim <title>*</title>
	titleBegin := strings.Index(title, ">")
	titleEnd := strings.Index(title, "</")
	if titleEnd < 0 || titleBegin < 0 {
		return title
	}
	return title[titleBegin+1 : titleEnd]
}
