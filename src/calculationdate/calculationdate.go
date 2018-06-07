package calculationdate

import (
	"strconv"
	//"calculationdate"

	"time"
)

const hour = 24
const minute = 60
const year = 365
const second = 60
const week = 7

func CalBetweenDate(startday int, startmonth int, startyear int, endday int, endmonth int, endyear int) string {

	startdate := time.Date(startyear, time.Month(startmonth), startday, 0, 0, 0, 0, time.UTC)
	enddate := time.Date(endyear, time.Month(endmonth), endday, 0, 0, 0, 0, time.UTC)

	diff := enddate.Sub(startdate)

	hours := diff.Hours()
	days := (hours / hour) + 1
	return strconv.FormatFloat(days, 'f', 0, 64)

}

func TransferDayToMinute(days int) string {

	minutes := days * hour * minute

	return strconv.Itoa(minutes)

}

func TransferDayToSecond(days int) string {

	seconds := days * hour * minute * second
	return strconv.Itoa(seconds)
}

func TransferDayToWeek(days int) (string, string) {

	weeks := days / week
	overdays := days % week

	return strconv.Itoa(weeks), strconv.Itoa(overdays)

}
func TransferDayToHours(days int) string {

	hours := days * hour
	return strconv.Itoa(hours)
}

func CalculatePercentOfYears(days int) string {

	percentofyears := (float64(days) / year) * 100
	return strconv.FormatFloat(percentofyears, 'f', 2, 64)
}
