package handler

import (
	"github.com/DamirLuketic/virtual_minds/config"
	"github.com/DamirLuketic/virtual_minds/db"
	"log"
	"strings"
	"time"
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

func (h *APIHandler) isCustomerValid(customerID string) (bool, error) {
	customers, err := h.db.GetCustomers()
	if err != nil {
		return false, err
	}
	for _, customer := range customers {
		if customer.UUID == customerID && customer.Active {
			return true, nil
		}
	}
	return false, nil
}

func (h *APIHandler) areIpAndUserAgentValid(ip, userAgent string) (bool, error) {
	ipValid, err := h.isIPValid(ip)
	if ipValid == false || err != nil {
		return false, err
	}
	userAgentValid, err := h.isUserAgentValid(userAgent)
	if userAgentValid == false || err != nil {
		return false, err
	}
	return true, nil
}

func (h *APIHandler) isIPValid(ip string) (bool, error) {
	ipBlackList, err := h.db.GetIPBlackList()
	if err != nil {
		return false, err
	}
	for _, ipBl := range ipBlackList {
		if ipBl.IP == ip {
			return false, nil
		}
	}
	return true, nil
}

func (h *APIHandler) isUserAgentValid(userAgent string) (bool, error) {
	uaBlackList, err := h.db.GetUABlackList()
	if err != nil {
		return false, err
	}
	for _, uaBl := range uaBlackList {
		if strings.Contains(userAgent, uaBl.UA) {
			return false, nil
		}
	}
	return true, nil
}

func (h *APIHandler) insertNotValidRequest(customerUUID string) {
	customerID := h.getCustomerID(customerUUID)
	hs := getNotValidRequestHourlyStatsEntity(&customerID)
	_, err := h.db.CreateHourlyStats(*hs)
	if err != nil {
		log.Fatalf("APIHandler.insertNotValidRequest. Error: %s", ErrorOnInsertHourlyStats)
	}
}

func (h *APIHandler) insertValidRequest(customerUUID string) {
	customerID := h.getCustomerID(customerUUID)
	hs := getValidRequestHourlyStatsEntity(&customerID)
	_, err := h.db.CreateHourlyStats(*hs)
	if err != nil {
		log.Fatalf("APIHandler.insertValidRequest. Error: %s", ErrorOnInsertHourlyStats)
	}
}

func (h *APIHandler) getCustomerID(customerUUID string) int64 {
	customer, err := h.db.GetCustomerByUUID(customerUUID)
	if err != nil {
		log.Fatalf("APIHandler.getCustomerID. Error: %s", ErrorOnFetchingCustomerData)
		return 0
	}
	return customer.ID
}

func getValidRequestHourlyStatsEntity(customerID *int64) *db.HourlyStats {
	return &db.HourlyStats{
		CustomerID:   customerID,
		Time:         getCurrentUTCTime(),
		RequestCount: 1,
	}
}

func getNotValidRequestHourlyStatsEntity(customerID *int64) *db.HourlyStats {
	return &db.HourlyStats{
		CustomerID:   customerID,
		Time:         getCurrentUTCTime(),
		InvalidCount: 1,
	}
}

func getCurrentUTCTime() *time.Time {
	currentTime := time.Now().UTC()
	return &currentTime
}
