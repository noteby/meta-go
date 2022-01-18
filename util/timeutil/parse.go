package timeutil

import (
	"fmt"
	"time"
)

// 计算从传入时间起到当前时间过去了多久
// 注意：传入时间不能大于当前时间，否则无意义
func TimeSince(t time.Time) string {
	var since string

	now := time.Now()
	duration := now.Sub(t)

	hours := duration.Hours()
	minutes := duration.Minutes()
	seconds := duration.Seconds()

	if hours > 23 {
		if hours < 24*31 {
			since = fmt.Sprintf("%d 天前", int(hours/24))
		} else {
			yearInterval := now.Year() - t.Year()
			nowMonth := int(now.Month())
			tMonth := int(t.Month())
			nowDay := now.Day()
			tDay := t.Day()
			// 实际相差的完整年数
			if nowMonth < tMonth || nowMonth == tMonth && nowDay < tDay {
				yearInterval--
			}
			// 实际相差的完整月数
			monthInterval := nowMonth + 12 - tMonth
			if nowDay < tDay {
				monthInterval--
			}
			monthInterval = yearInterval*12 + monthInterval%12

			if monthInterval < 12 {
				since = fmt.Sprintf("%d 月前", monthInterval)
			} else {
				since = fmt.Sprintf("%d 年前", monthInterval/12)
			}
		}
	} else if hours > 1 {
		since = fmt.Sprintf("%d 小时前", int(hours))
	} else if minutes > 1 {
		since = fmt.Sprintf("%d 分钟前", int(minutes))
	} else if seconds > 5 {
		since = fmt.Sprintf("%d 秒前", int(seconds))
	} else {
		since = fmt.Sprint("刚刚")
	}

	return since
}
