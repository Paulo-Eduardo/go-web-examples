package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Logging() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {log.Println(r.URL.Path, time.Since(start))} ()
			hf(w,r)
		}
	}
}

func Method(m string) Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			hf(w,r)
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
			f = m(f)
	}
	return f
}

func HelloMiddleware(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func ServeAdvanceMiddleware() {
	http.HandleFunc("/", Chain(HelloMiddleware, Method("GET"), Logging()))
	http.ListenAndServe(":8000", nil)
}
