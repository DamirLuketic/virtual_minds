package db

import "time"

type Customer struct {
	ID     int64  `gorm:"primary_key;type:int unsigned auto_increment;"`
	Name   string `gorm:"not null;"`
	Active bool   `gorm:"default:true;not null;"`
}

type IPBlacklist struct {
	IP string `gorm:"primary_key;"`
}

type UABlacklist struct {
	UA string `gorm:"primary_key;"`
}

type HourlyStats struct {
	ID           int64      `gorm:"primary_key;type:int unsigned auto_increment;"`
	CustomerID   *int64     `gorm:"index:customer_idx;type:int unsigned;not null;index:unique_customer_time,unique;"`
	Time         *time.Time `gorm:"not null;index:unique_customer_time,unique;"`
	RequestCount int64      `gorm:"type:int unsigned;not null;default:0"`
	InvalidCount int64      `gorm:"type:int unsigned;not null;default:0"`
	Customer     Customer   `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION;"`
}
