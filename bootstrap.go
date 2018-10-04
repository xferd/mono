package mono;

import (
	"net/http"
	"log"
)

type WebService interface {
	Config() *Config
	Routes() map[string]Handler
}

var (
	c *Config
	s WebService
)

func Boot(service WebService) {
	s = service
	conf := service.Config()
	if nil == conf {
		conf = &Config{}
	}

	log.Println("Started!")
	http.ListenAndServe(conf.addr(), &monoServerMux{})
}
