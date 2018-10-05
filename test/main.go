package main;

import (
	"github.com/xferd/mono"
	_ "github.com/xferd/mono/test/controller"
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
