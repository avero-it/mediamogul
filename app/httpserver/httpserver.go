package httpserver

import (
	"github.com/avero-it/mediamogul/app/aws/s3"
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
	s3client    s3.Client
	log         *logrus.Entry
	newRelicApp *newrelic.Application
	STTURI      string
	helper      httpHelper
}

func NewHttpServer(m *mux.Router, s3c s3.Client, l *logrus.Entry, n *newrelic.Application, sttURI string, h httpHelper) *server {
	return &server{
		log:         l,
		router:      m,
		s3client:    s3c,
		newRelicApp: n,
		STTURI:      sttURI,
		helper:      h,
	}
}
