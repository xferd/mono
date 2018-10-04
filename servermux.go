package mono

import(
	"net/http"
	"context"
	"log"
)

type monoServerMux struct {}

type Handler func(context.Context, http.ResponseWriter, *http.Request)

func(*monoServerMux)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	routes := s.Routes()
	url := r.URL.String()
	if handler, ok := routes[url]; ok {
		handler(context.Background(), w, r)
	}
	log.Println("It works! from mono mux", r.URL)
}

