package views

import (
	"fmt"
	"models"
	"net/http"
	"webstorage"
)

func init() {
	mux := webstorage.GetMux()
	mux.HandleFunc("/ypiresies", ypiresies)
}

func ypiresies(w http.ResponseWriter, r *http.Request) {
	db := webstorage.GetDb()
	res, err := db.Query("SELECT id,perigrafi,typos,proswpiko FROM ypiresia;")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ypiresies := []models.Ypiresia{}
	ypiresia := models.Ypiresia{}
	for res.Next() {
		err = res.Scan(
			&ypiresia.Id,
			&ypiresia.Perigrafi,
			&ypiresia.Typos,
			&ypiresia.Proswpiko)
		ypiresies = append(ypiresies, ypiresia)
	}
	res.Close()
	var t = webstorage.GetTemplate("ypiresies.html")
	if t != nil {
		t.ExecuteTemplate(w, "base", ypiresies)
	} else {
		fmt.Fprintf(w, "Error loading ypiresies.html")
	}
}
