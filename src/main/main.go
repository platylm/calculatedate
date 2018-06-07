package main

import (
	"apicalculatedate"
	"net/http"
)

func main() {
	http.HandleFunc("/apicalculatedurationdate", apicalculatedate.ApiGetStartAndEndDate)
	http.ListenAndServe(":9002", nil)

}
