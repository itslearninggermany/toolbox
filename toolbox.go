package toolbox

import (
	"fmt"
	"time"
)

// Version: 01.07.2019

// Creates a actually TempusString in the Format: "day-month-year_hour-min-sec"
func TempusString() string {
	now := time.Now()
	y, m, d := now.Date()
	h, min, s := now.Clock()
	return (fmt.Sprint(d, "-", m, "-", y, "_", h, "-", min, "-", s))
}

/*

func TestingNil(err error, ret bool) {
	if err != nil {
		fmt.Println(err)
		if ret {
			return
		}
	}
}


func TestingNilAndLog(l *log.Logger, err error, ret bool) {
	if err != nil {
		l.Println(err)
		if ret {
			return
		}
	}
}

*/
