package response

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
