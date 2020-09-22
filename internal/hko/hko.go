package hko

import (
	"log"
	"strings"
	"time"

	"github.com/pihao/ics/internal/h"
)

const solar_terms_path = "dist/solar-terms.ics"

type Event struct {
	T *time.Time
	V string
}

// GenSolarTerms 生成今后后五年内的节气日历
func GenSolarTerms() {
	es := GetEvent(5)
	s := event2ics(es)
	err := h.WriteFile(solar_terms_path, []byte(s))
	if err != nil {
		log.Println(err)
	}
}

func event2ics(es []*Event) string {
	var b strings.Builder
	b.WriteString(`BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//Pi Hao//Chinese Solar terms Calendar//EN
CALSCALE:GREGORIAN
METHOD:PUBLISH
X-WR-CALNAME:二十四节气
X-WR-TIMEZONE:Asia/Chongqing
X-WR-CALDESC:二十四节气 数据来自香港天文台
`)

	for _, e := range es {
		stamp := e.T.Format("20060102T150405Z")
		b.WriteString("BEGIN:VEVENT\n")
		b.WriteString("UID:solar-terms-" + stamp + "\n")
		b.WriteString("DTSTAMP:" + stamp + "\n")
		b.WriteString("DTSTART;VALUE=DATE:" + e.T.Format("20060102") + "\n")
		b.WriteString("DTEND;VALUE=DATE:" + e.T.Add(time.Hour*24).Format("20060102") + "\n")
		b.WriteString("STATUS:CONFIRMED\n")
		b.WriteString("SUMMARY:" + e.V + "\n")
		b.WriteString("END:VEVENT\n")
	}

	b.WriteString("END:VCALENDAR")

	return b.String()
}
