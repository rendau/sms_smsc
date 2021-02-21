package httpapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *St) router() http.Handler {
	r := mux.NewRouter()

	r.PathPrefix("/send").HandlerFunc(a.hSend).Methods("POST")
	r.PathPrefix("/bcast").HandlerFunc(a.hBcast).Methods("POST")
	r.PathPrefix("/balance").HandlerFunc(a.hGetBalance).Methods("GET")
	r.PathPrefix("/cron/check_balance").HandlerFunc(a.hCronCheckBalance).Methods("GET")

	return r
}
