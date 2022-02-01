package db

func (ds *MariaDBDataStoreImpl) CreateCustomer(customer *Customer) (*Customer, error) {
	re := ds.DB.Create(&customer)
	if re.Error != nil {
		return customer, re.Error
	}
	return customer, nil
}

func (ds *MariaDBDataStoreImpl) GetCustomerByUUID(uuid string) (customer *Customer, err error) {
	re := ds.DB.Where("uuid = ?", uuid).First(&customer)
	if re.Error != nil {
		return customer, re.Error
	}
	return customer, nil
}

func (ds *MariaDBDataStoreImpl) GetCustomers() (customers []*Customer, err error) {
	re := ds.DB.Find(&customers)
	if re.Error != nil {
		return customers, re.Error
	}
	return customers, nil
}

func (ds *MariaDBDataStoreImpl) CreateIPBlackList(ipBlackList *IPBlackList) (*IPBlackList, error) {
	re := ds.DB.Create(&ipBlackList)
	if re.Error != nil {
		return ipBlackList, re.Error
	}
	return ipBlackList, nil
}

func (ds *MariaDBDataStoreImpl) GetIPBlackList() (ipBlackList []*IPBlackList, err error) {
	re := ds.DB.Find(&ipBlackList)
	if re.Error != nil {
		return ipBlackList, re.Error
	}
	return ipBlackList, nil
}

func (ds *MariaDBDataStoreImpl) CreateUABlackList(uaBlackList *UABlackList) (*UABlackList, error) {
	re := ds.DB.Create(&uaBlackList)
	if re.Error != nil {
		return uaBlackList, re.Error
	}
	return uaBlackList, nil
}

func (ds *MariaDBDataStoreImpl) GetUABlackList() (uaBlackList []*UABlackList, err error) {
	re := ds.DB.Find(&uaBlackList)
	if re.Error != nil {
		return uaBlackList, re.Error
	}
	return uaBlackList, nil
}

func (ds *MariaDBDataStoreImpl) CreateHourlyStats(hourlyStats *HourlyStats) (*HourlyStats, error) {
	re := ds.DB.Create(&hourlyStats)
	if re.Error != nil {
		return hourlyStats, re.Error
	}
	return hourlyStats, nil
}

func (ds *MariaDBDataStoreImpl) UpdateOrCreateHourlyStats(hourlyStats *HourlyStats) (*HourlyStats, error) {
	isValid := isRequestValid(hourlyStats)
	hourlyStatsDB := &HourlyStats{}
	err := ds.DB.
		Where("customer_id = ? AND time = ?", *hourlyStats.CustomerID, *hourlyStats.Time).
		First(hourlyStatsDB)
	if err.Error != nil {
		return ds.CreateHourlyStats(hourlyStats)
	}
	if isValid {
		hourlyStatsDB.RequestCount++
	} else {
		hourlyStatsDB.InvalidCount++
	}
	err = ds.DB.Save(hourlyStatsDB)
	return hourlyStatsDB, err.Error
}

func isRequestValid(hourlyStats *HourlyStats) bool {
	if hourlyStats.RequestCount != 0 {
		return true
	}
	return false
}
