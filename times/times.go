package times

import (
	"github.com/yungsem/gotool/pattern"
	"time"
)

// Format 将时间按照指定的格式格式化成字符串
func Format(t time.Time, pattern string) string {
	return t.Format(pattern)
}

// Parse 将字符串按照指定的格式转换为 time.Time
func Parse(str string, pattern string) (time.Time, error) {
	t, err := time.Parse(pattern, str)
	if err != nil {
		return t, err
	}
	return t, nil
}

// ZeroTime 将 t 的时间置为 0 ，日期保持不变
func ZeroTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.YearDay(), 0, 0, 0, 0, t.Location())
}

// TimeList 返回 start 和 end 之间的时间点，每两个时间点之间的间隔 step
func TimeList(start time.Time, end time.Time, step time.Duration) []string {
	var timeList []string
	for IsLTE(start, end) {
		hm := Format(start, pattern.HourMinute)
		timeList = append(timeList, hm)
		start = start.Add(step)
	}
	return timeList
}
