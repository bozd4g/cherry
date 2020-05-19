package utils

import "net/http"

func StatusNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("404 - Not found!"))
	if err != nil {
		panic(err)
	}
}

func StatusInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte("Internal server error"))
	if err != nil {
		panic(err)
	}
}