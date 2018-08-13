package webstorage

import (
	"net/http"
)

var mux *http.ServeMux

func init() {
	mux = http.NewServeMux()
}

func GetMux() *http.ServeMux {
	return mux
}
