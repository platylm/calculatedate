package apicalculatedate

// func Test_CheckApiStatusCode_inpui_url_should_be_return_200(t *testing.T) {
// 	url := "http://localhost:9002/apicalculatedurationdate?startdate=4/1/2018&enddate=4/6/2018"
// 	expected := 200
// 	req := httptest.NewRequest("GET", url, nil)
// 	res := httptest.NewRecorder()

// 	CheckApiStatusCode(res, req)

// 	if res.Code != expected {
// 		t.Fatalf("Expected status code == %d , but got %d ", expected, res.Code)
// 	}

// }

// func Test_apiApiGetStartAndEndDate_input_start_date_4_1_2018_and_end_date_4_6_2018_should_be_4_1_2018_and_4_6_2018(t *testing.T) {
// 	url := "http://localhost:9002/apicalculatedurationdate?startdate=4/1/2018&enddate=4/6/2018"
// 	expected := dateInput{
// 		Startday:   4,
// 		Startmonth: 1,
// 		Startyear:  2018,
// 		Endday:     4,
// 		Endmonth:   6,
// 		Endyear:    2018,
// 	}
// 	req := httptest.NewRequest("GET", url, nil)
// 	startenddate := ApiGetStartAndEndDate(req)

// 	if startenddate != expected {
// 		t.Errorf("Should be %d, but got %d ", expected, startenddate)
// 	}

// }
