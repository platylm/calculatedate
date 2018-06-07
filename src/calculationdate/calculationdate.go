package calculationdate

import (
	"strconv"
	//"calculationdate"

	"time"
)

const hour = 24
const minute = 60
const year = 365

func CalBetweenDate(startday int, startmonth int, startyear int, endday int, endmonth int, endyear int) string {

	startdate := time.Date(startyear, time.Month(startmonth), startday, 0, 0, 0, 0, time.UTC)
	enddate := time.Date(endyear, time.Month(endmonth), endday, 0, 0, 0, 0, time.UTC)

	diff := enddate.Sub(startdate)

	hours := diff.Hours()
	days := (hours / hour) + 1
	return strconv.FormatFloat(days, 'f', 0, 64)

}

func TranferDayToMinute(days int) string {

	minutes := days * hour * minute

	return strconv.Itoa(minutes)

}

func TranferDayToSecond(days int) string {

	seconds := days * 24 * 60 * 60
	return strconv.Itoa(seconds)
}

func TransferDayToHours(days int) string {

	hours := days * hour
	return strconv.Itoa(hours)
}

func CalculatePercentOfYears(days int) string {

	percentofyears := (float64(days) / year) * 100
	return strconv.FormatFloat(percentofyears, 'f', 2, 64)
}
