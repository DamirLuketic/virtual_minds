package handler

import (
	"github.com/DamirLuketic/virtual_minds/clients/request"
	"github.com/DamirLuketic/virtual_minds/db"
	"github.com/DamirLuketic/virtual_minds/localtime"
	"net/http"
)

type APIHandler interface {
	NewRequest(w http.ResponseWriter, r *http.Request)
	CustomerPerDayStatistics(w http.ResponseWriter, r *http.Request)
}

type APIHandlerImpl struct {
	DB            db.DataStore
	RequestClient request.Client
	LocalTime     localtime.Time
	APIUsername   string
	APIPassword   string
}

type Request struct {
	CustomerUUID string `json:"customerUUID"`
	RemoteIP     string `json:"remoteIP"`
}

type HourlyStatsResponse struct {
	ValidRequests    int64 `json:"valid_requests"`
	NotValidRequests int64 `json:"not_valid_requests"`
	TotalRequests    int64 `json:"total_requests"`
}
