package request

type AddMovieInputRequest struct {
	Name    string `json:"name" binding:"required,min=5,max=55"`
	Code    string `json:"code" binding:"required,min=16,max=24,alphanum"`
	Address string `json:"address" binding:"required,min=5,max=465"`
}

type UpdateMovieInputRequest struct {
	Name        string `json:"name" binding:"required,min=5,max=55"`
	ReleaseDate string `json:"release_date"`
}
