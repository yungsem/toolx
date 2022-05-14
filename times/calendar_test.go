package times

import (
	"fmt"
	"github.com/yungsem/toolx/pattern"
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	date, err := Parse("2021-04-05", pattern.Date)
	if err != nil {
		t.Error(err)
	}
	info := Info(date)
	fmt.Printf("%+v\n", info)

	ti := time.Now()
	fmt.Println(ti.ISOWeek())
}

func TestFirstDayOfISOWeek(t *testing.T) {
	date := FirstDayOfYearWeek(2020, 53)
	fmt.Println(date)
}

func TestRangeOfYearWeek(t *testing.T) {
	r := RangeOfYearWeek(2021, 13)
	fmt.Printf("%+v\n", r)
	fmt.Println(r.start)
	fmt.Println(r.end)
}

func TestZeroTime(t *testing.T) {
	ti := ZeroTime(time.Now())

	fmt.Println(ti)
	fmt.Println(Format(ti, pattern.DateTime))
}

func TestGap(t *testing.T) {
	t1 := time.Now()
	time.Sleep(3 * time.Second)
	t2 := time.Now()
	gap := Gap(t1, t2)
	fmt.Println(gap)
}


