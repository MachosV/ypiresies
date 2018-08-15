package views

import (
	"algorithm"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	filename := "Υπηρεσίες " + r.PostFormValue("month") + " " + r.PostFormValue("year") + ".xlsx"
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	algorithm.Algorithm(date)
	os.Rename("temp.xlsx", filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening file, /getprogramma")
	}
	w.Header().Set("Content-Disposition", "attachment; filename=\""+filename+"\"")
	io.Copy(w, f)
	os.Remove(filename)
}
