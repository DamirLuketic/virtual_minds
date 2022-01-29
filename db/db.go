package db

func (ds *MariaDBDataStoreImpl) CreateCustomer(customer Customer) (Customer, error) {
	re := ds.db.Create(&customer)
	if re.Error != nil {
		return customer, re.Error
	}
	return customer, nil
}

func (ds *MariaDBDataStoreImpl) CreateIPBlacklist(ipBlacklist IPBlacklist) (IPBlacklist, error) {
	re := ds.db.Create(&ipBlacklist)
	if re.Error != nil {
		return ipBlacklist, re.Error
	}
	return ipBlacklist, nil
}

func (ds *MariaDBDataStoreImpl) CreateUABlacklist(uaBlacklist UABlacklist) (UABlacklist, error) {
	re := ds.db.Create(&uaBlacklist)
	if re.Error != nil {
		return uaBlacklist, re.Error
	}
	return uaBlacklist, nil
}

func (ds *MariaDBDataStoreImpl) CreateHourlyStats(hourlyStats HourlyStats) (HourlyStats, error) {
	re := ds.db.Create(&hourlyStats)
	if re.Error != nil {
		return hourlyStats, re.Error
	}
	return hourlyStats, nil
}
