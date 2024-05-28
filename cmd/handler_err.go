package main

import "net/http"

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 400, "Something went wrong")
}
