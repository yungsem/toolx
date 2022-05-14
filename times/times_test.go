package times

import (
	"fmt"
	"github.com/yungsem/toolx/pattern"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	ti := Format(time.Now(), pattern.Date)
	fmt.Printf("%s\n", ti)

	ti = Format(time.Now(), pattern.DateTime)
	fmt.Printf("%s\n", ti)
}

func TestParse(t *testing.T) {
	ti, err := Parse("2021-03-26", pattern.Date)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%v\n", ti)

	ti, err = Parse("2021-03-26 21:39:12", pattern.DateTime)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%v\n", ti)
}

func TestTimeList(t *testing.T) {
	start, err := Parse("2021-04-05 20:00:00", pattern.DateTime)
	if err != nil {
		t.Error(err)
	}
	end, err := Parse("2021-04-06 08:00:00", pattern.DateTime)
	if err != nil {
		t.Error(err)
	}
	timeList := TimeList(start, end, 40*time.Minute)
	fmt.Println(timeList)
}
