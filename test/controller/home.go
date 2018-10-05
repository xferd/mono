package controller

import(
	"context"
	"net/http"

	"github.com/xferd/mono"
)

func init() {
	mono.Route("/", home)
}

func home(c context.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("homepage"))
}