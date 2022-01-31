package handler

import (
	"github.com/DamirLuketic/virtual_minds/clients/request"
	"github.com/DamirLuketic/virtual_minds/db"
	"net/http"
)

type APIHandler interface {
	NewRequest(w http.ResponseWriter, r *http.Request)
	CustomerByDayStatistics(w http.ResponseWriter, r *http.Request)
}

type APIHandlerImpl struct {
	DB            db.DataStore
	RequestClient request.Client
	APIUsername   string
	APIPassword   string
}

type Request struct {
	CustomerUUID string `json:"customerUUID"`
	RemoteIP     string `json:"remoteIP"`
}
