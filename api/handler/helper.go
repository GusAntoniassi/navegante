package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
)

func AddContentType(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	contentTypeHeader := "application/json"

	w.Header().Set("Content-Type", contentTypeHeader)
	next(w, r)
}

func GetNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("{\"error\":\"Not found\"}"))

	fmt.Printf(
		"[notfound] %s | %d | \t %s | %s | %s %s\n",
		time.Now().Format(negroni.LoggerDefaultDateFormat),
		http.StatusNotFound,
		"", // This would be the duration
		r.Host,
		r.Method,
		r.RequestURI,
	)
}
