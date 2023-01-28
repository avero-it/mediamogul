package mediamogul

import (
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

type MediaMogul struct {
	verbose     bool // config
	router      *mux.Router
	debug       bool // config
	newRelicApp *newrelic.Application
	log         *logrus.Entry
}

func NewMediaMogul(router *mux.Router, newRelicApp *newrelic.Application, log *logrus.Entry) *MediaMogul {
	return &MediaMogul{
		verbose:     false,
		router:      router,
		debug:       false,
		newRelicApp: newRelicApp,
		log:         log,
	}
}

func (mediaMogul MediaMogul) Run() error {

	select {}
}
