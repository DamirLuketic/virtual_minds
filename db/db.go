package db

import (
	"gorm.io/gorm"
)

type MariaDBDataStore struct {
	db *gorm.DB
}

func (ds *MariaDBDataStore) CreateCustomer(customer Customer) (Customer, error) {
	re := ds.db.Create(&customer)
	if re.Error != nil {
		return customer, re.Error
	}
	return customer, nil
}

func (ds *MariaDBDataStore) CreateIPBlacklist(ipBlacklist IPBlacklist) (IPBlacklist, error) {
	re := ds.db.Create(&ipBlacklist)
	if re.Error != nil {
		return ipBlacklist, re.Error
	}
	return ipBlacklist, nil
}

func (ds *MariaDBDataStore) CreateUABlacklist(uaBlacklist UABlacklist) (UABlacklist, error) {
	re := ds.db.Create(&uaBlacklist)
	if re.Error != nil {
		return uaBlacklist, re.Error
	}
	return uaBlacklist, nil
}

func (ds *MariaDBDataStore) CreateHourlyStats(hourlyStats HourlyStats) (HourlyStats, error) {
	re := ds.db.Create(&hourlyStats)
	if re.Error != nil {
		return hourlyStats, re.Error
	}
	return hourlyStats, nil
}
