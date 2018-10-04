package main;

import (
	"net/http"
	"github.com/xferd/mono"
	"context"
)

func main() {
	mono.Boot(&MainWebService{})
}

type MainWebService struct{}

func(*MainWebService)Config() *mono.Config {
	return &mono.Config{
		Addr: ":8084",
	}
}

func(*MainWebService)Routes() map[string]mono.Handler {
	return map[string]mono.Handler{
		"/": func(c context.Context, w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello world"))
		},
		"/bye": func(c context.Context, w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("good bye"))
		},
	}
}