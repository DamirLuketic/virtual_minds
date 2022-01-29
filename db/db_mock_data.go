package db

import "log"

func (ds *MariaDBDataStoreImpl) setMockData() {
	ds.migrateCustomerData()
	ds.migrateIPBlacklistData()
	ds.migrateUABlacklistData()
}

func (ds *MariaDBDataStoreImpl) migrateCustomerData() {
	data := getCustomerData()
	for _, d := range data {
		_, err := ds.CreateCustomer(d)
		if err != nil {
			log.Fatalf("Error on inserting mockdata. Error: %s", err.Error())
		}
		log.Printf("Data for customer with id: %s created in DB", d.ID)
	}
}

func (ds *MariaDBDataStoreImpl) migrateIPBlacklistData() {
	data := getIPBlacklistData()
	for _, d := range data {
		_, err := ds.CreateIPBlacklist(d)
		if err != nil {
			log.Fatalf("Error on inserting mockdata. Error: %s", err.Error())
		}
		log.Printf("Data for IPBlacklistData with IP: %s created in DB", d.IP)
	}
}

func (ds *MariaDBDataStoreImpl) migrateUABlacklistData() {
	data := getUABlacklistData()
	for _, d := range data {
		_, err := ds.CreateUABlacklist(d)
		if err != nil {
			log.Fatalf("Error on inserting mockdata. Error: %s", err.Error())
		}
		log.Printf("Data for UABlacklist with UA: %s created in DB", d.UA)
	}
}

func getCustomerData() []Customer {
	return []Customer{
		{
			Name:   "Big News Media Corp",
			Active: true,
		},
		{
			Name:   "Online Mega Store",
			Active: true,
		},
		{
			Name:   "Nachoroo Delivery",
			Active: false,
		},
		{
			Name:   "Euro Telecom Group",
			Active: true,
		},
	}
}

func getIPBlacklistData() []IPBlacklist {
	return []IPBlacklist{
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

func getUABlacklistData() []UABlacklist {
	return []UABlacklist{
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
