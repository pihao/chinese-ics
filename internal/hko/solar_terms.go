package hko

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

func GetEvent(years int) []*Event {
	n := time.Now().Year()
	var es []*Event
	for y := n - 1; y <= n+years; y++ {
		getSourceText(y, func(b []byte) {
			e, ok := parseEvent(b, y)
			if ok {
				es = append(es, e)
			}
		})
	}
	return es
}

func parseEvent(b []byte, year int) (*Event, bool) {
	b, err := decodeBIG5(b)
	if err != nil {
		log.Printf("decode BIG5 failed: %v", err)
		return nil, false
	}

	if !bytes.HasPrefix(b, []byte(strconv.Itoa(year))) {
		return nil, false
	}

	fs := bytes.Fields(b)
	if len(fs) != 4 || !bytes.HasPrefix(fs[2], []byte("星期")) {
		return nil, false
	}

	t, err := time.ParseInLocation("2006年1月2日", string(fs[0]), hkloc)
	if err != nil {
		log.Printf("parse time failed: %v", err)
		return nil, false
	}

	return &Event{
		T: &t,
		V: simp(string(fs[3])),
	}, true
}

func getSourceText(year int, parse func([]byte)) {
	url := fmt.Sprintf("https://www.hko.gov.hk/tc/gts/time/calendar/text/files/T%dc.txt", year)
	log.Println("GET", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("GET source failed: %v", err)
		return
	}

	defer resp.Body.Close()
	scan := bufio.NewScanner(resp.Body)
	for scan.Scan() {
		parse(scan.Bytes())
	}
	if err := scan.Err(); err != nil {
		log.Printf("read body error: %v", err)
	}
}

func decodeBIG5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

var trad2simp = map[string]string{
	"小寒": "小寒",
	"大寒": "大寒",
	"立春": "立春",
	"雨水": "雨水",
	"驚蟄": "惊蛰",
	"春分": "春分",
	"清明": "清明",
	"穀雨": "谷雨",
	"立夏": "立夏",
	"小滿": "小满",
	"芒種": "芒种",
	"夏至": "夏至",
	"小暑": "小暑",
	"大暑": "大暑",
	"立秋": "立秋",
	"處暑": "处暑",
	"白露": "白露",
	"秋分": "秋分",
	"寒露": "寒露",
	"霜降": "霜降",
	"立冬": "立冬",
	"小雪": "小雪",
	"大雪": "大雪",
	"冬至": "冬至",
}

func simp(t string) string {
	s, ok := trad2simp[t]
	if ok {
		return s
	}
	return t
}

var hkloc *time.Location

func init() {
	l, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		log.Fatal(err)
	}
	hkloc = l
}
