package main

import "net/http"

func HandlerReadMessage(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, http.StatusOK, struct{}{})
}
