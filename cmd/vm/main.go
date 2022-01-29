package main

import (
	"github.com/DamirLuketic/virtual_minds/api/handler"
	"github.com/DamirLuketic/virtual_minds/api/security"
	"github.com/DamirLuketic/virtual_minds/config"
	"github.com/DamirLuketic/virtual_minds/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	conf := config.NewServerConfig()
	srv := configureServer(conf)
	log.Fatal(srv.ListenAndServe())
}

func configureServer(conf *config.Config) *http.Server {
	return &http.Server{
		Addr:         conf.ServerConfig.Port,
		Handler:      newHandler(conf),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func newHandler(conf *config.Config) http.Handler {
	mdb := db.NewMariaDBDataStore(conf.DBConfig)
	handlers := handler.NewApiHandler(mdb, conf.ServerConfig)
	router := mux.NewRouter()
	appendMiddleware(router, conf)
	apiSubRouter := router.PathPrefix("/api").Subrouter()
	apiSubRouter.HandleFunc("/new_request", handlers.NewRequest).Methods(http.MethodPost)
	apiSubRouter.HandleFunc("/customer_statistic", handlers.CustomerByDayStatistics).Methods(http.MethodGet)
	return router
}

func appendMiddleware(router *mux.Router, conf *config.Config) {
	middleware := security.NewSecurity(conf.ServerConfig)
	router.Use(middleware.ValidateAuthorizationMiddleware)
}
