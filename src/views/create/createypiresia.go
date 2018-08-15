package create

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	color := r.PostFormValue("color")
	color = strings.Split(color, "#")[1]
	color = "FF" + strings.ToUpper(color)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db := webstorage.GetDb()
	stmt, err := db.Prepare("insert into ypiresia (perigrafi,typos,proswpiko,color) values(?,?,?,?);")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = stmt.Exec(onoma, typos, proswpiko, color)
	stmt.Close()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/ypiresies", http.StatusMovedPermanently)
}
