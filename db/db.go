package db

func (ds *MariaDBDataStoreImpl) CreateCustomer(customer Customer) (Customer, error) {
	re := ds.db.Create(&customer)
	if re.Error != nil {
		return customer, re.Error
	}
	return customer, nil
}

func (ds *MariaDBDataStoreImpl) GetCustomerByUUID(uuid string) (customer Customer, err error) {
	re := ds.db.Where("uuid = ?", uuid).First(&customer)
	if re.Error != nil {
		return customer, re.Error
	}
	return customer, nil
}

func (ds *MariaDBDataStoreImpl) GetCustomers() (customers []Customer, err error) {
	re := ds.db.Find(&customers)
	if re.Error != nil {
		return customers, re.Error
	}
	return customers, nil
}

func (ds *MariaDBDataStoreImpl) CreateIPBlackList(ipBlackList IPBlackList) (IPBlackList, error) {
	re := ds.db.Create(&ipBlackList)
	if re.Error != nil {
		return ipBlackList, re.Error
	}
	return ipBlackList, nil
}

func (ds *MariaDBDataStoreImpl) GetIPBlackList() (ipBlackList []IPBlackList, err error) {
	re := ds.db.Find(&ipBlackList)
	if re.Error != nil {
		return ipBlackList, re.Error
	}
	return ipBlackList, nil
}

func (ds *MariaDBDataStoreImpl) CreateUABlackList(uaBlackList UABlackList) (UABlackList, error) {
	re := ds.db.Create(&uaBlackList)
	if re.Error != nil {
		return uaBlackList, re.Error
	}
	return uaBlackList, nil
}

func (ds *MariaDBDataStoreImpl) GetUABlackList() (uaBlackList []UABlackList, err error) {
	re := ds.db.Find(&uaBlackList)
	if re.Error != nil {
		return uaBlackList, re.Error
	}
	return uaBlackList, nil
}

func (ds *MariaDBDataStoreImpl) CreateHourlyStats(hourlyStats HourlyStats) (HourlyStats, error) {
	re := ds.db.Create(&hourlyStats)
	if re.Error != nil {
		return hourlyStats, re.Error
	}
	return hourlyStats, nil
}
