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
