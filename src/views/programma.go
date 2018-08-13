package views

import (
	"fmt"
	"net/http"
	"webstorage"
)

func init() {
	mux := webstorage.GetMux()
	mux.HandleFunc("/programma", programma)
}

func programma(w http.ResponseWriter, r *http.Request) {
	var t = webstorage.GetTemplate("programma.html")
	if t != nil {
		t.ExecuteTemplate(w, "base", nil)
	} else {
		fmt.Fprintf(w, "Error loading programma.html")
	}
}
