package response

type MovieResponse struct {
	CreatedAt   string `json:"created_at"`
	Name        string `json:"name"`
	Moviecode   string `json:"moviecode"`
	ReleaseDate string `json:"release_date"`
}

var _table_cs = "movies"

// TableName get sql table name movies
func (m *MovieResponse) TableName() string {
	return _table_cs
}
