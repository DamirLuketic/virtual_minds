package db

type DataStore interface {
	CreateCustomer(customer Customer) (Customer, error)
	CreateIPBlacklist(ipBlacklist IPBlacklist) (IPBlacklist, error)
	CreateUABlacklist(uaBlacklist UABlacklist) (UABlacklist, error)
	CreateHourlyStats(hourlyStats HourlyStats) (HourlyStats, error)
}
