package db

import "time"

type Customer struct {
	ID     int64  `gorm:"primary_key;type:int unsigned auto_increment;"`
	UUID   string `gorm:"not null;index:unique_uuid,unique;"`
	Name   string `gorm:"not null;"`
	Active bool   `gorm:"default:true;not null;"`
}

type IPBlackList struct {
	IP string `gorm:"primary_key;"`
}

type UABlackList struct {
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

// TableName overrides the table name used by IPBlackList to `ip_black_list`
func (IPBlackList) TableName() string {
	return "ip_black_list"
}

// TableName overrides the table name used by UABlackList to `ua_black_list`
func (UABlackList) TableName() string {
	return "ua_black_list"
}
