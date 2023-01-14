package mediamogul

import (
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

func TestNewMediaMogul(t *testing.T) {
	type args struct {
		router      *mux.Router
		newRelicApp *newrelic.Application
		log         *logrus.Entry
	}
	tests := []struct {
		name string
		args args
		want *MediaModul
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMediaMogul(tt.args.router, tt.args.newRelicApp, tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMediaMogul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMediaMogul_Run(t *testing.T) {
	type fields struct {
		verbose     bool
		router      *mux.Router
		debug       bool
		newRelicApp *newrelic.Application
		log         *logrus.Entry
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mediaMogul := MediaModul{
				verbose:     tt.fields.verbose,
				router:      tt.fields.router,
				debug:       tt.fields.debug,
				newRelicApp: tt.fields.newRelicApp,
				log:         tt.fields.log,
			}
			if err := mediaMogul.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
