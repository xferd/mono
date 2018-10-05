package mono

import(
	"net/http"
	"context"
	"log"
)

var(
	routes map[string]Handler
)

func init() {
	routes = make(map[string]Handler)
}

type monoServerMux struct {}

type Handler func(context.Context, http.ResponseWriter, *http.Request)

func(*monoServerMux)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	if handler, ok := routes[url]; ok {
		handler(context.Background(), w, r)
	}
	log.Println("It works! from mono mux", r.URL)
}

func Route(url string, handler Handler) {
	routes[url] = handler
}
