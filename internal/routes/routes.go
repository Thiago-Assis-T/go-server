package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/thiago", thiagoHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1> Hello World!</h1>")
}

func thiagoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1> Hello Thiago! </h1>")
}
