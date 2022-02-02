package db

import (
	"database/sql"
	"fmt"
	"github.com/DamirLuketic/virtual_minds/config"
	mysqld "github.com/go-sql-driver/mysql"
	"gopkg.in/retry.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func NewMariaDBDataStore(c *config.DBConfig) DataStore {
	ds := MariaDBDataStoreImpl{}
	var err error
	dsn := getDSN(c)
	createDB(c)
	for a := retry.Start(connectRetryStrategy(), nil); a.Next(); {
		ds.DB, err = gorm.Open(mysql.Open(dsn))
		if err == nil {
			break
		} else {
			log.Println("DB connect fail")
		}
	}
	if err != nil {
		log.Fatalf("DB connect fail. Error: %s", err.Error())
	}
	ds.migrate()
	// Set mock data (for demo app purpose only)
	ds.setMockData()
	return &ds
}

func getDSN(c *config.DBConfig) string {
	cfg := mysqld.NewConfig()
	cfg.DBName = c.MySQLDatabase
	cfg.ParseTime = true
	cfg.User = c.MySQLUser
	cfg.Passwd = c.MySQLPassword
	cfg.Net = "tcp"
	cfg.Params = map[string]string{
		"charset": "utf8mb4",
		"loc":     "Local",
	}
	cfg.Addr = fmt.Sprintf("%v:%v", c.MySQLHost, c.MySQLPort)
	dsn := cfg.FormatDSN()
	return dsn
}

func connectRetryStrategy() retry.Strategy {
	return retry.LimitTime(30*time.Second,
		retry.Exponential{
			Initial: 1 * time.Second,
			Factor:  1.5,
		},
	)
}

func createDB(c *config.DBConfig) {
	dataSource := fmt.Sprintf("%s:%s@tcp(db:%s)/", c.MySQLUser, c.MySQLPassword, c.MySQLPort)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Printf("MySQL connected successfully")
	}
	for a := retry.Start(connectRetryStrategy(), nil); a.Next(); {
		_, err = db.Exec("DROP DATABASE IF EXISTS vm")
		if err == nil {
			break
		} else {
			log.Println("Successfully drop database..")
		}
	}
	if err != nil {
		log.Fatalf("DB connect fail. Error: %s", err.Error())
	}
	_, err = db.Exec("CREATE DATABASE vm")
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Println("Successfully created database..")
	}
}

func (ds *MariaDBDataStoreImpl) migrate() {
	err := ds.DB.Migrator().DropTable(&Customer{})
	handleError(err)
	err = ds.DB.AutoMigrate(&Customer{})
	handleError(err)
	err = ds.DB.Migrator().DropTable(&IPBlackList{})
	handleError(err)
	err = ds.DB.AutoMigrate(&IPBlackList{})
	handleError(err)
	err = ds.DB.Migrator().DropTable(&UABlackList{})
	handleError(err)
	err = ds.DB.AutoMigrate(&UABlackList{})
	handleError(err)
	err = ds.DB.Migrator().DropTable(&HourlyStats{})
	handleError(err)
	err = ds.DB.AutoMigrate(&HourlyStats{})
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func isRequestValid(hourlyStats *HourlyStats) bool {
	if hourlyStats.RequestCount != 0 {
		return true
	}
	return false
}
