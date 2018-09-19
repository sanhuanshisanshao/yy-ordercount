package spider

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

func GetRemainFreeTimes(str string) (s int64, err error) {
	buf := bytes.NewBuffer([]byte(str))
	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		return 0, err
	}

	doc.Find(".data_tit").Find("span").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			str := selection.Find("span").Text()
			s, _ = strconv.ParseInt(str, 10, 64)
		}
	})
	return
}
