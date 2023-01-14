package mediamogul

import (
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

type MediaModul struct {
	verbose     bool // config
	router      *mux.Router
	debug       bool // config
	newRelicApp *newrelic.Application
	log         *logrus.Entry
}

func NewMediaMogul(router *mux.Router, newRelicApp *newrelic.Application, log *logrus.Entry) *MediaModul {
	return &MediaModul{
		verbose:     false,
		router:      router,
		debug:       false,
		newRelicApp: newRelicApp,
		log:         log,
	}
}

func (mediaMogul MediaModul) Run() error {

	select {}
}
