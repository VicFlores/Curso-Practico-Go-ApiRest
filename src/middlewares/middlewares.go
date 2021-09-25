package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {

		return func(res http.ResponseWriter, req *http.Request) {
			flag := true // simulando loqueo

			if flag {
				hf(res, req)
			} else {
				fmt.Fprintf(res, "No auth")
				return
			}
		}
	}
}

func Logging() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(res http.ResponseWriter, req *http.Request) {

			start := time.Now()

			defer func() {
				log.Println(req.Method, req.URL.Path, time.Since(start))
			}()

			hf(res, req) // hace que salte al siguiente middleware
		}
	}
}
