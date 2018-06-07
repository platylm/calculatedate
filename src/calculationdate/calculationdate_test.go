package calculationdate

import (
	"testing"
)

func Test_Cal_Between_Date_Input_Startday_Endday_Should_Be_152(t *testing.T) {

	startday := 4
	startmouth := 1
	startyear := 2018
	endday := 4
	endmonth := 6
	endyear := 2018

	resultDay := CalBetweenDate(startday, startmouth, startyear, endday, endmonth, endyear)

	expectedDay := "152"

	if resultDay != expectedDay {
		t.Errorf("expectedDay %s but %s ", expectedDay, resultDay)

	}

}

func Test_Cal_Between_Date_Input_Startday_Endday_Should_Be_9079(t *testing.T) {

	startday := 27
	startmouth := 7
	startyear := 1993
	endday := 4
	endmonth := 6
	endyear := 2018

	resultDay := CalBetweenDate(startday, startmouth, startyear, endday, endmonth, endyear)

	expectedDay := "9079"

	if resultDay != expectedDay {
		t.Errorf("expectedDay %s but %s ", expectedDay, resultDay)

	}

}

func Test_Transfer_Day_To_Minute_Input_152_Should_Be_218880(t *testing.T) {

	day := 152
	resultMinute := TransferDayToMinute(day)
	expectedMinute := "218880"

	if resultMinute != expectedMinute {
		t.Errorf("expectedMinute %s but %s ", expectedMinute, resultMinute)

	}

}

func Test_Transfer_Day_To_Minute_Input_152_Should_Be_13073760(t *testing.T) {

	day := 9079
	resultMinute := TransferDayToMinute(day)
	expectedMinute := "13073760"

	if resultMinute != expectedMinute {
		t.Errorf("expectedHour %s but %s ", expectedMinute, resultMinute)

	}

}

func Test_Transfer_Day_To_Second_Input_152_Should_Be_13132800(t *testing.T) {
	day := 152
	resultSecond := TransferDayToSecond(day)
	expectedSecond := "13132800"

	if resultSecond != expectedSecond {
		t.Errorf("expectedSecond %s but %s ", expectedSecond, resultSecond)

	}

}

func Test_Transfer_Day_To_Week_Input_152_Should_Be_21_Week_5_Day(t *testing.T) {
	day := 152
	resultWeek, resultDay := TransferDayToWeek(day)
	expectedWeek := "21"
	expectedDay := "5"

	if resultWeek != expectedWeek {
		t.Errorf("expectedWeek %s but %s ", expectedWeek, resultWeek)

	}
	if resultDay != expectedDay {
		t.Errorf("expectedDay %s but %s ", expectedDay, resultDay)

	}
}

func Test_Transfer_Day_To_Week_Input_9079_Should_Be_1297_Week_0_Day(t *testing.T) {
	day := 9079
	resultWeek, resultDay := TransferDayToWeek(day)
	expectedWeek := "1297"
	expectedDay := "0"

	if resultWeek != expectedWeek {
		t.Errorf("expectedWeek %s but %s ", expectedWeek, resultWeek)

	}
	if resultDay != expectedDay {
		t.Errorf("expectedDay %s but %s ", expectedDay, resultDay)

	}

}
func Test_Tranfer_Day_To_Hours_Input_152_Should_Be_3648(t *testing.T) {
	day := 152
	resultHour := TransferDayToHours(day)
	expectedHour := "3648"

	if resultHour != expectedHour {
		t.Errorf("expectedSecond %s but %s ", expectedHour, resultHour)

	}

}

func Test_Tranfer_Day_To_Hours_Input_9079_Should_Be_217896(t *testing.T) {
	day := 9079
	resultHour := TransferDayToHours(day)
	expectedHour := "217896"

	if resultHour != expectedHour {
		t.Errorf("expectedHour %s but %s ", expectedHour, resultHour)

	}

}

func Test_Calculate_Percent_Of_Years_Input_152_Should_Be_41_dot_64(t *testing.T) {
	day := 152
	resultPercentOfYears := CalculatePercentOfYears(day)
	expectedPercentOfYears := "41.64"

	if resultPercentOfYears != expectedPercentOfYears {
		t.Errorf("expectedSecond %s but %s ", expectedPercentOfYears, resultPercentOfYears)

	}

}

func Test_Calculate_Percent_Of_Years_Input_9079_Should_Be_2487_dot_40(t *testing.T) {
	day := 9079
	resultPercentOfYears := CalculatePercentOfYears(day)
	expectedPercentOfYears := "2487.40"

	if resultPercentOfYears != expectedPercentOfYears {
		t.Errorf("expectedSecond %s but %s ", expectedPercentOfYears, resultPercentOfYears)

	}

}
