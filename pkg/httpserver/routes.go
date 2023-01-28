package httpserver

import "net/http"

func (srv server) Routes() {
	srv.router.HandleFunc("/healthz", srv.healthz()).Methods(http.MethodGet)
	srv.router.HandleFunc("/api/s3audiotonlp", srv.handleS3AudioToNLP()).Methods(http.MethodPost)
}
