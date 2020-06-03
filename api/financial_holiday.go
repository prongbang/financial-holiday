package api

import (
	"encoding/json"
	"net/http"

	"github.com/prongbang/bank-holiday/pkg/holiday"
)

var utility holiday.Utility

func init() {
	utility = holiday.NewUtility()
}

// FinancialHolidayHandler is the handler
func FinancialHolidayHandler(w http.ResponseWriter, r *http.Request) {
	year := r.URL.Query().Get("year")
	if len(year) >= 4 {
		holidayData, _ := json.Marshal(utility.GetFinancialHoliday(year))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(holidayData)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 BadRequest"))
	}
}
