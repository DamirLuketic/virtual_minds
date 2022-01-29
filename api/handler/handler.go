package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ErrorMalformedRequestBody             = "malformed JSON"
	ErrorMissingCustomerUUIDInRequestBody = "missing customerUUID in request body"
	ErrorMissingRemoteIPInRequestBody     = "missing remoteIP in request body"
	ErrorCustomerNotValid                 = "customer not valid"
	ErrorIPorUAonBlackList                = "ip or userAgent on black list"
	ErrorOnFetchingCustomerData           = "error on fetching customer data"
	ErrorOnInsertHourlyStats              = "error on insert hourly stats"
)

func (h *APIHandler) NewRequest(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf(ErrorMalformedRequestBody)
		return
	}
	reqBody := Request{}
	err = json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Errorf(http.StatusText(http.StatusInternalServerError))
		return
	}
	customerUUID := reqBody.CustomerUUID
	if customerUUID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf(ErrorMissingCustomerUUIDInRequestBody)
		return
	}
	isCustomerValid, err := h.isCustomerValid(customerUUID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Errorf(http.StatusText(http.StatusInternalServerError))
		return
	}
	if !isCustomerValid {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf(ErrorCustomerNotValid)
		return
	}
	remoteIP := reqBody.RemoteIP
	if remoteIP == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf(ErrorMissingRemoteIPInRequestBody)
		return
	}
	userAgent := r.UserAgent()
	isValid, err := h.areIpAndUserAgentValid(remoteIP, userAgent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Errorf(http.StatusText(http.StatusInternalServerError))
		// Request inserted in DB,
		// but to take into consideration that failure occurred due to internal error
		h.insertNotValidRequest(customerUUID)
		return
	}
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf(ErrorIPorUAonBlackList)
		h.insertNotValidRequest(customerUUID)
		return
	}
	h.insertValidRequest(customerUUID)
	// TODO: Implement rest
}

func (h *APIHandler) CustomerByDayStatistics(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}
