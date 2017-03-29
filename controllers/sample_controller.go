package controllers

import "net/http"

// SampleController : Sample controller example
func SampleController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	data := "hello world"
	w.Write([]byte(data))
}
