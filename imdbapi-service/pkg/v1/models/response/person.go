package response

type PersonResponse struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Age         int8   `json:"age"`
	ID          int    `json:"-"`
	IsActive    *uint8 `json:"isActive"`
	Mobile      string `json:"mobile"`
	Updateddate string `json:"updateddate"`
	Createddate string `json:"createddate"`
}

var _table_person = "person"

// TableName get sql table name person
func (m *PersonResponse) TableName() string {
	return _table_person
}
