package views

import (
	"datastorage"
	"fmt"
	"models"
	"net/http"
	"webstorage"
)

func init() {
	mux := webstorage.GetMux()
	mux.HandleFunc("/proswpiko", proswpiko)
}

type Context struct {
	Metafrastes []models.Atomo
	Akroates    []models.Atomo
	Stratiwtes  []models.Atomo
}

func proswpiko(w http.ResponseWriter, r *http.Request) {
	metafrastes := datastorage.LoadAtoma(datastorage.METAFRASTIS)
	akroates := datastorage.LoadAtoma(datastorage.AKROATIS)
	stratiwtes := datastorage.LoadAtoma(datastorage.STRATIWTIS)
	var context Context
	context.Metafrastes = metafrastes
	context.Akroates = akroates
	context.Stratiwtes = stratiwtes
	var t = webstorage.GetTemplate("proswpiko.html")
	if t != nil {
		t.ExecuteTemplate(w, "base", context)
	} else {
		fmt.Fprintf(w, "Error loading proswpiko.html")
	}
}
