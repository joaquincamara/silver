package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer next(w, r)
		log.Printf("TIME: %s, REMOTE-IP: %s, HOST: %s, URI: %s, METHOD: %s", time.Now(), r.RemoteAddr, r.Host, r.RequestURI, r.Method)
	})
}
