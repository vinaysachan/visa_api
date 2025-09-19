package models

// AppCountry maps to `app_country` table
type AppCountry struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Code   string `gorm:"type:char(2);not null" json:"code"`
	Name   string `gorm:"type:varchar(200);not null" json:"name"`
	Status uint8  `gorm:"type:tinyint unsigned;not null" json:"status"`
	Flag   string `gorm:"type:varchar(300);not null" json:"flag"`
}

// TableName overrides the default table name
func (AppCountry) TableName() string {
	return "app_country"
}

// AppArrivalPort maps to `app_arrival_port` table
type AppArrivalPort struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name   string `gorm:"type:varchar(400);not null" json:"name"`
	Status uint8  `gorm:"type:tinyint unsigned;default:0" json:"status"`
}

// TableName overrides the default table name
func (AppArrivalPort) TableName() string {
	return "app_arrival_port"
}
