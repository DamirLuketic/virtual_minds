package handler

import (
	"github.com/DamirLuketic/virtual_minds/clients/request"
	"github.com/DamirLuketic/virtual_minds/config"
	"github.com/DamirLuketic/virtual_minds/db"
	"log"
	"strings"
	"time"
)

func NewApiHandler(db db.DataStore, requestClient request.Client, c *config.ServerConfig) APIHandler {
	return &APIHandlerImpl{
		DB:            db,
		RequestClient: requestClient,
		APIUsername:   c.APIUser,
		APIPassword:   c.APIPassword,
	}
}

func (h *APIHandlerImpl) isCustomerValid(customerID string) (bool, error) {
	customers, err := h.DB.GetCustomers()
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

func (h *APIHandlerImpl) areIpAndUserAgentValid(ip, userAgent string) (bool, error) {
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

func (h *APIHandlerImpl) isIPValid(ip string) (bool, error) {
	ipBlackList, err := h.DB.GetIPBlackList()
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

func (h *APIHandlerImpl) isUserAgentValid(userAgent string) (bool, error) {
	uaBlackList, err := h.DB.GetUABlackList()
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

func (h *APIHandlerImpl) insertNotValidRequest(customerUUID string) {
	customerID := h.getCustomerID(customerUUID)
	hs := getNotValidRequestHourlyStatsEntity(&customerID)
	_, err := h.DB.CreateHourlyStats(*hs)
	if err != nil {
		log.Fatalf("APIHandler.insertNotValidRequest. Error: %s", ErrorOnInsertHourlyStats)
	}
}

func (h *APIHandlerImpl) insertValidRequest(customerUUID string) {
	customerID := h.getCustomerID(customerUUID)
	hs := getValidRequestHourlyStatsEntity(&customerID)
	_, err := h.DB.CreateHourlyStats(*hs)
	if err != nil {
		log.Fatalf("APIHandler.insertValidRequest. Error: %s", ErrorOnInsertHourlyStats)
	}
}

func (h *APIHandlerImpl) getCustomerID(customerUUID string) int64 {
	customer, err := h.DB.GetCustomerByUUID(customerUUID)
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
