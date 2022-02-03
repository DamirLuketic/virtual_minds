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
	ErrorMissingCustomerUUIDOrDate        = "missing customerUUID or date in query params"
	ErrorNotValidDate                     = "date not valid"
)

// NewRequest godoc
// @Summary Handle Request
// @Description Handle Request
// @Tags Request
// @Produce
// @Success 200
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
// @Failure 500 "Internal server error"
// @Router /api/new_request
// @Param command body Request
func (h *APIHandlerImpl) NewRequest(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(ErrorMalformedRequestBody)
		return
	}
	reqBody := Request{}
	err = json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(http.StatusText(http.StatusInternalServerError))
		return
	}
	customerUUID := reqBody.CustomerUUID
	if customerUUID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(ErrorMissingCustomerUUIDInRequestBody)
		return
	}
	isCustomerValid, err := h.isCustomerValid(customerUUID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(http.StatusText(http.StatusInternalServerError))
		return
	}
	if !isCustomerValid {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(ErrorCustomerNotValid)
		return
	}
	remoteIP := reqBody.RemoteIP
	if remoteIP == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(ErrorMissingRemoteIPInRequestBody)
		return
	}
	userAgent := r.UserAgent()
	isValid, err := h.areIpAndUserAgentValid(remoteIP, userAgent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(http.StatusText(http.StatusInternalServerError))
		// Request inserted in DB,
		// but to take into consideration that failure occurred due to internal error
		h.insertNotValidRequest(customerUUID)
		return
	}
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(ErrorIPorUAonBlackList)
		h.insertNotValidRequest(customerUUID)
		return
	}
	h.insertValidRequest(customerUUID)
	h.RequestClient.New(remoteIP, w, r)
}

// CustomerPerDayStatistics godoc
// @Summary Get customers statistic for selected day
// @Description Get customers statistic for selected day
// @Tags Statistic
// @Produce HourlyStatsResponse
// @Param customerUUID query string true "Customer UUID"
// @Param date query string true "Requested statistics date. Format YYYY-MM-DD"
// @Success 200
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
// @Failure 500 "Internal server error"
// @Router /api/customer_statistic
func (h *APIHandlerImpl) CustomerPerDayStatistics(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	customerUUID := queryParams.Get("customerUUID")
	date := queryParams.Get("date")
	if customerUUID == "" || date == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(ErrorMissingCustomerUUIDOrDate)
		return
	}
	customerID := h.getCustomerID(customerUUID)
	isCustomerValid, err := h.isCustomerValid(customerUUID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(http.StatusText(http.StatusInternalServerError))
		return
	}
	if !isCustomerValid {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(ErrorCustomerNotValid)
		return
	}
	validDate, err := h.inspectDate(date)
	if !isCustomerValid {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(ErrorNotValidDate)
		return
	}
	hourlyStats, err := h.DB.GetDailyCustomerHourlyStats(&customerID, validDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(http.StatusText(http.StatusInternalServerError))
		return
	}
	response := countRequests(hourlyStats)
	responseBody, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseBody)
	if err != nil {
		fmt.Printf("handler.CustomerByDayStatistics. error writing response. Error: %s", err.Error())
	}
}
