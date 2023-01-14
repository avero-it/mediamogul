package httpserver

import (
	"github.com/avero-it/mediamogul/app/aws/s3"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestNewServer(t *testing.T) {
	type args struct {
		m      *mux.Router
		s3c    s3.Client
		l      *logrus.Entry
		n      *newrelic.Application
		sttURI string
		h      httpHelper
	}
	tests := []struct {
		name string
		args args
		want *server
	}{
		{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewHttpServer(tt.args.m, tt.args.s3c, tt.args.l, tt.args.n, tt.args.sttURI, tt.args.h)
		})
	}
}
