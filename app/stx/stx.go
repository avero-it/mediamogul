package stx

import (
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

type Stx struct {
	verbose     bool // config
	router      *mux.Router
	debug       bool // config
	newRelicApp *newrelic.Application
	log         *logrus.Entry
}

func NewStx(router *mux.Router, newRelicApp *newrelic.Application, log *logrus.Entry) *Stx {
	return &Stx{
		verbose:     false,
		router:      router,
		debug:       false,
		newRelicApp: newRelicApp,
		log:         log,
	}
}

func (stx Stx) Run() error {

	select {}
}
