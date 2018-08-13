package create

import (
	"fmt"
	"net/http"
	"strconv"
	"webstorage"
)

func init() {
	mux := webstorage.GetMux()
	mux.HandleFunc("/createypiresia", createYpiresia)
}

func createYpiresia(w http.ResponseWriter, r *http.Request) {
	onoma := r.PostFormValue("onoma")
	typos, err := strconv.Atoi(r.PostFormValue("typos"))
	proswpiko, err := strconv.Atoi(r.PostFormValue("proswpiko"))
	stepallagis, err := strconv.Atoi(r.PostFormValue("stepallagis"))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db := webstorage.GetDb()
	stmt, err := db.Prepare("insert into ypiresia (perigrafi,typos,proswpiko,stepallagis) values(?,?,?,?);")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = stmt.Exec(onoma, typos, proswpiko, stepallagis)
	stmt.Close()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/ypiresies", http.StatusMovedPermanently)
}
