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

type MovieResponse struct {
	CreatedAt   string `json:"created_at"`
	Name        string `json:"name"`
	Moviecode   string `json:"moviecode"`
	ReleaseDate string `json:"release_date"`
}

// TableName get sql table name movies
func (m *MovieResponse) TableName() string {
	return _table_mc
}

type MovieDetailsResponse struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Age         string `json:"age"`
	ID          int    `json:"-"`
	IsActive    *uint8 `json:"isActive"`
	Mobile      string `json:"mobile"`
	Updateddate string `json:"updateddate"`
	Createddate string `json:"createddate"`
}

// TableName get sql table name movies
func (m *MovieDetailsResponse) TableName() string {
	return _table_mc
}
