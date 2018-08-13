package retrieve

import (
	"fmt"
	"models"
	"net/http"
	"time"
	"webstorage"
)

func init() {
	mux := webstorage.GetMux()
	mux.HandleFunc("/retrieveproswpiko", retrieveProswpiko)
}

func retrieveProswpiko(w http.ResponseWriter, r *http.Request) {
	var arxi int64
	var telos int64
	id := r.URL.Query().Get("id")
	db := webstorage.GetDb()
	res, err := db.Query("select * from proswpiko where id =?;", id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		http.Redirect(w, r, "/proswpiko", http.StatusMovedPermanently)
		return
	}
	atomo := models.Atomo{}
	ypiresies := []models.YpiresiaAtomou{}
	for res.Next() {
		err = res.Scan(
			&atomo.Id,
			&atomo.Name,
			&atomo.Typos)
	}
	res, err = db.Query("select id,aitia,arxi,telos from adeies where atomo = ?", atomo.Id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		http.Redirect(w, r, "/proswpiko", http.StatusMovedPermanently)
		return
	}
	for res.Next() {
		adeia := models.Adeia{}
		err = res.Scan(
			&adeia.Id,
			&adeia.Aitia,
			&arxi,
			&telos,
		)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		adeia.Arxi = time.Unix(arxi, 0)
		adeia.Telos = time.Unix(telos, 0)
		atomo.Adeies = append(atomo.Adeies, adeia)
	}

	atomo.YpiresiesAtomou = ypiresies
	res.Close()
	var t = webstorage.GetTemplate("atomo.html")
	if t != nil {
		t.ExecuteTemplate(w, "base", atomo)
	} else {
		fmt.Fprintf(w, "Error loading atomo.html")
	}
}
