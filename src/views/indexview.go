package views

import (
	"fmt"
	"net/http"
	"webstorage"
)

func init() {
	mux := webstorage.GetMux()
	mux.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	var t = webstorage.GetTemplate("index.html")
	if t != nil {
		t.ExecuteTemplate(w, "base", nil)
	} else {
		fmt.Fprintf(w, "Error loading index.html")
	}
}
