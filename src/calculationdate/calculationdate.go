package calculationdate

import (
	"fmt"
	"strconv"
	"time"
)

const hour = 24
const minute = 60
const year = 365
const second = 60
const week = 7

<<<<<<< HEAD
func CalBetweenDate(startday, startmonth, startyear, endday, endmonth, endyear int) string {
=======
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
		DayNameFrom: SetFullNameDate(startyear, startmonth, startday),
		DayNameTo:   SetFullNameDate(endyear, endmonth, endday),
		ResultDay:   CalBetweenDate(startday, startmonth, startyear, endday, endmonth, endyear),
		Second:      TransferDayToSecond(day),
		Minute:      TransferDayToMinute(day),
		Hours:       TransferDayToHours(day),
		Percent:     CalculatePercentOfYears(day),
	}
}

<<<<<<< HEAD
func CalBetweenDate(startday, startmonth, startyear, endday, endmonth, endyear int) string {
=======
func CalBetweenDate(startday int, startmonth int, startyear int, endday int, endmonth int, endyear int) string {
>>>>>>> b0cb900223e910bf8e41804ab7f933807156c2ad
>>>>>>> 287351a5a44827c083bcea9d24d9312ab1d15b29

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

<<<<<<< HEAD
func SetFullNameDate(year, month, day int) string {

=======
func SetFullNameDate(year int, month int, day int) string {
>>>>>>> b0cb900223e910bf8e41804ab7f933807156c2ad
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
