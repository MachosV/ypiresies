package create

import (
	"fmt"
	"net/http"
	"strconv"
	"webstorage"
)

func init() {
	mux := webstorage.GetMux()
	mux.HandleFunc("/createproswpiko", createProswpiko)
}

func createProswpiko(w http.ResponseWriter, r *http.Request) {
	onoma := r.PostFormValue("onoma")
	typos, _ := strconv.Atoi(r.PostFormValue("typos"))
	db := webstorage.GetDb()
	stmt, err := db.Prepare("INSERT INTO proswpiko (onoma,typos) VALUES(?,?);")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = stmt.Exec(onoma, typos)
	stmt.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		http.Redirect(w, r, "/proswpiko", http.StatusMovedPermanently)
	}
}
