package create

import (
	"fmt"
	"net/http"
	"time"
	"webstorage"
)

func init() {
	mux := webstorage.GetMux()
	mux.HandleFunc("/createadeia", createAdeia)
}

func createAdeia(w http.ResponseWriter, r *http.Request) {
	atomo := r.PostFormValue("atomo")
	aitia := r.PostFormValue("aitia")
	arxi, err := time.Parse("2006-01-02", r.PostFormValue("arxi"))
	telos, err := time.Parse("2006-01-02", r.PostFormValue("telos"))
	redirectURL := "/retrieveproswpiko?id=" + atomo
	if err != nil {
		fmt.Println(err)
		return
	}
	arxiToStore := arxi.Unix()
	telosToStore := telos.Unix()
	db := webstorage.GetDb()
	stmt, err := db.Prepare("insert into adeies (aitia,arxi,telos,atomo) values(?,?,?,?);")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = stmt.Exec(aitia, arxiToStore, telosToStore, atomo)
	stmt.Close()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
	}
}
