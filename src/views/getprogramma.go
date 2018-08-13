package views

import (
	"algorithm"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"webstorage"
)

func init() {
	mux := webstorage.GetMux()
	mux.HandleFunc("/getprogramma", getprogramma)
}

func getprogramma(w http.ResponseWriter, r *http.Request) {
	month, err := strconv.Atoi(r.PostFormValue("month"))
	year, err := strconv.Atoi(r.PostFormValue("year"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	algorithm.Algorithm(date)
}
