package db

import "log"

func (ds *MariaDBDataStoreImpl) setMockData() {
	ds.migrateCustomerData()
	ds.migrateIPBlackListData()
	ds.migrateUABlackListData()
}

func (ds *MariaDBDataStoreImpl) migrateCustomerData() {
	data := getCustomerData()
	for _, d := range data {
		_, err := ds.CreateCustomer(d)
		if err != nil {
			log.Fatalf("Error on inserting mockdata. Error: %s", err.Error())
		}
		log.Printf("Data for customer with id: %d created in DB", d.ID)
	}
}

func (ds *MariaDBDataStoreImpl) migrateIPBlackListData() {
	data := getIPBlacklistData()
	for _, d := range data {
		_, err := ds.CreateIPBlackList(d)
		if err != nil {
			log.Fatalf("Error on inserting mockdata. Error: %s", err.Error())
		}
		log.Printf("Data for IPBlackListData with IP: %s created in DB", d.IP)
	}
}

func (ds *MariaDBDataStoreImpl) migrateUABlackListData() {
	data := getUABlackListData()
	for _, d := range data {
		_, err := ds.CreateUABlackList(d)
		if err != nil {
			log.Fatalf("Error on inserting mockdata. Error: %s", err.Error())
		}
		log.Printf("Data for UABlackList with UA: %s created in DB", d.UA)
	}
}

func getCustomerData() []*Customer {
	return []*Customer{
		{
			UUID:   "72167dcf-3618-4af7-8815-ca51f3bb775a",
			Name:   "Big News Media Corp",
			Active: true,
		},
		{
			UUID:   "c334dde1-4566-4ee2-b389-68cfff63c7d4",
			Name:   "Online Mega Store",
			Active: true,
		},
		{
			UUID:   "b5676b7c-b430-4eaf-8c21-1e613f508b92",
			Name:   "Nachoroo Delivery",
			Active: false,
		},
		{
			UUID:   "27e90d55-4194-4aea-81a6-7dc000ea1737",
			Name:   "Euro Telecom Group",
			Active: true,
		},
	}
}

func getIPBlacklistData() []*IPBlackList {
	return []*IPBlackList{
		{
			IP: "0",
		},
		{
			IP: "2130706433",
		},
		{
			IP: "4294967295",
		},
	}
}

func getUABlackListData() []*UABlackList {
	return []*UABlackList{
		{
			UA: "A6-Indexer",
		},
		{
			UA: "Googlebot-News",
		},
		{
			UA: "Googlebot",
		},
	}
}
