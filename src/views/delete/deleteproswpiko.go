package delete

import (
	"fmt"
	"net/http"
	"webstorage"
)

func init() {
	mux := webstorage.GetMux()
	mux.HandleFunc("/deleteproswpiko", deleteProswpiko)
}

func deleteProswpiko(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	db := webstorage.GetDb()
	stmt, err := db.Prepare("DELETE FROM proswpiko where id =?;")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = stmt.Exec(id)
	stmt.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		http.Redirect(w, r, "/proswpiko", http.StatusMovedPermanently)
	}
}
