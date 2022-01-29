package handler

type Request struct {
	CustomerUUID string `json:"customerUUID"`
	RemoteIP     string `json:"remoteIP"`
}
