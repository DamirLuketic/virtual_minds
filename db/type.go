package db

import "gorm.io/gorm"

type MariaDBDataStoreImpl struct {
	DB *gorm.DB
}

type DataStore interface {
	CreateCustomer(customer Customer) (Customer, error)
	GetCustomerByUUID(uuid string) (Customer, error)
	GetCustomers() ([]Customer, error)
	CreateIPBlackList(ipBlackList IPBlackList) (IPBlackList, error)
	GetIPBlackList() ([]IPBlackList, error)
	CreateUABlackList(uaBlackList UABlackList) (UABlackList, error)
	GetUABlackList() ([]UABlackList, error)
	CreateHourlyStats(hourlyStats HourlyStats) (HourlyStats, error)
	UpdateOrCreateHourlyStats(hourlyStats *HourlyStats) (*HourlyStats, error)
}
