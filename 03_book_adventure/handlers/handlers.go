package handlers

import "net/http"

func Homepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))

}
