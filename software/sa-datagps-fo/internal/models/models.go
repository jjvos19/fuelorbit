package models

import (
	"database/sql/driver"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Data struct {
	gorm.Model
	Id            uint      `gorm:"primary key"`
	TankerId      uint      `gorm:"not null"`
	GpsCoordinate Point     `gorm:"type:point"`
	Volume        float64   `gorm:"null"`
	StateMotor    string    `gorm:"not null"`
	HashDevice    string    `gorm:"not null"`
	SendDate      time.Time `gorm:"column:send_date;type:date"`
	GroupBlck     int       `gorm:"not null"`
	//HashBlck      uint      `gorm:"column:hash_blck;null"`
}

func (Data) TableName() string {
	return "tkr_data"
}

func (p Point) Value() (driver.Value, error) {
	out := []byte{'('}
	out = strconv.AppendFloat(out, p.X, 'f', -1, 64)
	out = append(out, ',')
	out = strconv.AppendFloat(out, p.Y, 'f', -1, 64)
	out = append(out, ')')
	return out, nil
}

type Group struct {
	gorm.Model
	Id        uint   `gorm:"primary key"`
	HashGroup string `gorm:"not null"`
	IdStart   uint   `gorm:"not null;column:data_id_start"`
	IdFinish  uint   `gorm:"not null;column:data_id_finish"`
	HashBlkc  string `gorm:"null"`
}

func (Group) TableName() string {
	return "tkr_group"
}
