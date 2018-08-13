package main

import (
	"log"
	"net/http"
	_ "views"
	_ "views/create"
	_ "views/delete"
	_ "views/retrieve"
	"webstorage"
	_ "webstorage"
)

func main() {
	mux := webstorage.GetMux()
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	log.Println("Server started..")
	http.ListenAndServe(":8000", mux)
}
