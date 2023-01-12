package entities

import (
	"time"
)

var _table_mc = "movies"

type Movie struct {
	ID          int64     `gorm:"primaryKey;autoIncrement;not null"`
	Moviecode   string    `gorm:"column:moviecode;index:Code_UniqueIndex" json:"moviecode"`
	Name        string    `gorm:"column:name" json:"name"`
	ReleaseDate time.Time `gorm:"column:release_date" json:"release_date"`
	Status      *uint8    `gorm:"column:status;default:1" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name movies
func (m *Movie) TableName() string {
	return _table_mc
}
