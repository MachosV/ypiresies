package views

import (
	"fmt"
	"models"
	"net/http"
	"webstorage"
)

func init() {
	mux := webstorage.GetMux()
	mux.HandleFunc("/proswpiko", proswpiko)
}

func proswpiko(w http.ResponseWriter, r *http.Request) {
	adeies := []models.Adeia{}
	ypiresies := []models.YpiresiaAtomou{}
	proswpiko := []models.Atomo{}
	db := webstorage.GetDb()
	rows, err := db.Query("SELECT * FROM proswpiko;")
	if err != nil {
		fmt.Fprintf(w, "Database error %s\n", err.Error())
		return
	}
	atomo := models.Atomo{}
	for rows.Next() {
		atomo.Adeies = adeies
		atomo.YpiresiesAtomou = ypiresies
		err = rows.Scan(
			&atomo.Id,
			&atomo.Name,
			&atomo.Typos)
		proswpiko = append(proswpiko, atomo)
	}
	rows.Close()
	var t = webstorage.GetTemplate("proswpiko.html")
	if t != nil {
		t.ExecuteTemplate(w, "base", proswpiko)
	} else {
		fmt.Fprintf(w, "Error loading proswpiko.html")
	}
}
