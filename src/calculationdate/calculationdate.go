package calculationdate

import (
	"fmt"
	"strconv"
	"time"
)

const (
	hour   = 24
	minute = 60
	year   = 365
	second = 60
	week   = 7
)

type durationDate struct {
	DayNameFrom   string `json:"daynamefrom"`
	DayNameTo     string `json:"daynameto"`
	ResultDay     string `json:"resultday"`
	ResultMonth   string `json:"resultmonth"`
	ResultYear    string `json:"resultyear"`
	ResultOverDay string `json:"resultoverday"`
	Second        string `json:"second"`
	Minute        string `json:"minute"`
	Hours         string `json:"hours"`
	Week          string `json:"week"`
	DayOverOfWeek string `json:"dayoverofweek"`
	Percent       string `json:"percent"`
}

func MakeJSON(startday, startmonth, startyear, endday, endmonth, endyear int) durationDate {
	day, _ := strconv.Atoi(CalBetweenDate(startday, startmonth, startyear, endday, endmonth, endyear))
	return durationDate{
		DayNameFrom: SetFullNameDate(startday, startmonth, startyear),
		DayNameTo:   SetFullNameDate(endday, endmonth, endyear),
		ResultDay:   CalBetweenDate(startday, startmonth, startyear, endday, endmonth, endyear),
		Second:      TransferDayToSecond(day),
		Minute:      TransferDayToMinute(day),
		Hours:       TransferDayToHours(day),
		Percent:     CalculatePercentOfYears(day),
	}
}

func CalBetweenDate(startday, startmonth, startyear, endday, endmonth, endyear int) string {

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

func SetFullNameDate(year, month, day int) string {
	fullnamedate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	weekdays := fullnamedate.Weekday()
	days := fullnamedate.Day()
	months := fullnamedate.Month()
	years := fullnamedate.Year()
	return fmt.Sprintf("%s %s %s %s", weekdays.String(), strconv.Itoa(days), months.String(), strconv.Itoa(years))
}
func CalculatePercentOfYears(days int) string {
	percentofyears := (float64(days) / year) * 100
	return strconv.FormatFloat(percentofyears, 'f', 2, 64)
}

func CheckEqualDay(startday, startmouth, startyear, endday, endmonth, endyear int) string {
	if startday == endday && startmouth == endmonth && startyear == endyear {
		return "1"
	}
	return "not equal"
}
