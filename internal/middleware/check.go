package middleware

import (
	"log"
	"net/http"
)

func Checking(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		if req.Method != http.MethodPost {

			log.Println("405 - Method Not Allowed")
			res.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		//if contentType := req.Header.Get("Content-Type"); contentType != "text/plain" {
		//
		//	res.Header().Set("Content-Type", "application/json")
		//
		//	log.Println("415 - Unsupported Media Type", contentType)
		//	res.WriteHeader(http.StatusUnsupportedMediaType)
		//	return
		//}

		next.ServeHTTP(res, req)
	})
}
