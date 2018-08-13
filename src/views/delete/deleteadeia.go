package delete

import (
	"fmt"
	"net/http"
	"webstorage"
)

func init() {
	mux := webstorage.GetMux()
	mux.HandleFunc("/deleteadeia", deleteAdeia)
}

func deleteAdeia(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	idatomou := r.PostFormValue("idatomou")
	db := webstorage.GetDb()
	stmt, err := db.Prepare("DELETE FROM adeies where id =?;")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = stmt.Exec(id)
	stmt.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		http.Redirect(w, r, "/retrieveproswpiko?id="+idatomou, http.StatusMovedPermanently)
	}
}
