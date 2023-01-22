package middleware

import (
	"log"
	"net/http"
)

// BasicAuthMiddleware ..
type BasicAuthMiddleware struct {
	Next http.Handler
}

func (bam *BasicAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if bam.Next == nil {
		bam.Next = http.DefaultServeMux
	}

	if r.Method != http.MethodGet {
		username, password, ok := r.BasicAuth()
		if !ok {
			log.Println("Error parsing basic auth")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if username != "admin" {
			log.Println("Username is not correct")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if password != "123456" {
			log.Println("Password is not correct")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	bam.Next.ServeHTTP(w, r)

}
