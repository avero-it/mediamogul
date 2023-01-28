package httpserver

import (
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

type Server interface {
	Routes()
}

// guard
var _ Server = server{}

type server struct {
	router      *mux.Router
	log         *logrus.Entry
	newRelicApp *newrelic.Application
	STTURI      string
	helper      httpHelper
}

func NewHttpServer(m *mux.Router, l *logrus.Entry, n *newrelic.Application, sttURI string, h httpHelper) *server {
	return &server{
		log:         l,
		router:      m,
		newRelicApp: n,
		STTURI:      sttURI,
		helper:      h,
	}
}
