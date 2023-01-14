package httpserver

import (
	"github.com/avero-it/mediamogul/app/aws/s3"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_server_handleS3AudioToNLP(t *testing.T) {

	type fields struct {
		router      *mux.Router
		s3client    s3.Client
		log         *logrus.Entry
		newRelicApp *newrelic.Application
		STTURI      string
		h           httpHelper
	}

	_ = fields{}

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "",
			fields: fields{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new request to the handler func
			req, err := http.NewRequest("POST", "/api/s3audiotonlp", nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a new ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			srv := server{
				router:      tt.fields.router,
				s3client:    tt.fields.s3client,
				log:         tt.fields.log,
				newRelicApp: tt.fields.newRelicApp,
				STTURI:      tt.fields.STTURI,
			}

			handler := srv.handleS3AudioToNLP()
			handler.ServeHTTP(rr, req)

			// Check the status code of the response
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

		})
	}
}

func TestErrorHandler(t *testing.T) {
	// Create a dummy request and response writer for the test
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Call the errorHandler function with a sample status code
	errorHandler(w, r, http.StatusInternalServerError)

	// Check that the response status code was set correctly
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	}
}
