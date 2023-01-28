package httpserver

import "net/http"

func (srv server) Routes() {
	srv.router.HandleFunc("/healthz", srv.healthz()).Methods(http.MethodGet)
}
