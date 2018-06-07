package apicalculatedate

import (
	"calculationdate"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type dateInput struct {
	Startday   int
	Startmonth int
	Startyear  int
	Endday     int
	Endmonth   int
	Endyear    int
}

// func CheckApiStatusCode(w http.ResponseWriter, r *http.Request) {
// 	testStatusCode := statusCode{Status: 200}
// 	json, _ := json.Marshal(testStatusCode)
// 	w.Write(json)
// }

var startEndDate dateInput

func ApiGetStartAndEndDate(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("startdate")
	endDate := r.URL.Query().Get("enddate")

	spitStartDate := strings.Split(startDate, "/")
	spitEndDate := strings.Split(endDate, "/")

	startEndDate = dateInput{
		Startday:   converseStringToInt(spitStartDate[0]),
		Startmonth: converseStringToInt(spitStartDate[1]),
		Startyear:  converseStringToInt(spitStartDate[2]),
		Endday:     converseStringToInt(spitEndDate[0]),
		Endmonth:   converseStringToInt(spitEndDate[1]),
		Endyear:    converseStringToInt(spitEndDate[2]),
	}

	result, _ := json.Marshal(calculationdate.MakeJSON(startEndDate.Startday, startEndDate.Startmonth, startEndDate.Startyear, startEndDate.Endday, startEndDate.Endmonth, startEndDate.Endyear))
	w.Write(result)

}

func converseStringToInt(number string) int {
	s, _ := strconv.Atoi(number)

	return s
}
