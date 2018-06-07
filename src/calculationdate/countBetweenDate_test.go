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

func Test_Tranfer_Day_To_Minute_Input_152_Should_Be_218880(t *testing.T) {

	day := 152
	resultMinute := TranferDayToMinute(day)
	expectedMinute := "218880"

	if resultMinute != expectedMinute {
		t.Errorf("expectedMinute %s but %s ", expectedMinute, resultMinute)

	}

}

func Test_Tranfer_Day_To_Minute_Input_152_Should_Be_13073760(t *testing.T) {

	day := 9079
	resultMinute := TranferDayToMinute(day)
	expectedMinute := "13073760"

	if resultMinute != expectedMinute {
		t.Errorf("expectedHour %s but %s ", expectedMinute, resultMinute)

	}

}

func Test_Tranfer_Day_To_Second_Input_152_Should_Be_13132800(t *testing.T) {
	day := 152
	resultHour := TranferDayToSecond(day)
	expectedHour := "13132800"

	if resultHour != expectedHour {
		t.Errorf("expectedSecond %s but %s ", expectedHour, resultHour)

	}

}
