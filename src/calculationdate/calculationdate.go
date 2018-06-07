package calculationdate

import (
	"strconv"
	//"calculationdate"

	"time"
)

const hour = 24

func CalBetweenDate(startday int, startmonth int, startyear int, endday int, endmonth int, endyear int) string {

	startdate := time.Date(startyear, time.Month(startmonth), startday, 0, 0, 0, 0, time.UTC)
	enddate := time.Date(endyear, time.Month(endmonth), endday, 0, 0, 0, 0, time.UTC)

	diff := enddate.Sub(startdate)

	hours := diff.Hours()
	days := (hours / hour) + 1
	return strconv.FormatFloat(days, 'f', 0, 64)

}
