package entities

import (
	"time"
)

var _table_person = "person"

type Person struct {
	ID        int       `gorm:"column:id;primary_key"`
	Email     string    `gorm:"column:email;index:Code_Email_UniqueIndex"`
	Name      string    `gorm:"column:name"`
	Age       int8      `gorm:"column:age"`
	Mobile    string    `gorm:"column:mobile"`
	IsActive  *uint8    `gorm:"column:is_active"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName get sql table name person
func (m *Person) TableName() string {
	return _table_person
}
