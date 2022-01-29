package handler

import (
	"github.com/DamirLuketic/virtual_minds/config"
	"github.com/DamirLuketic/virtual_minds/db"
)

type APIHandler struct {
	db          db.DataStore
	APIUsername string
	APIPassword string
}

func NewApiHandler(db db.DataStore, c *config.ServerConfig) APIHandler {
	return APIHandler{
		db:          db,
		APIUsername: c.APIUser,
		APIPassword: c.APIPassword,
	}
}
