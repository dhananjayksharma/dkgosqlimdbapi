package handlers

import (
	"imdbapi-service/internals/util"
	"imdbapi-service/pkg/v1/models/movies"

	"github.com/gin-gonic/gin"
)

// MovieHandler
type MovieHandler interface {
	GetMovieList(c *gin.Context)
	CreateMovie(c *gin.Context)
	UpdateMovieByID(c *gin.Context)
}

// moviesHandler
type moviesHandler struct {
	service movies.MovieService
}

// NewMovieHandler
func NewMovieHandler(service movies.MovieService) MovieHandler {
	return &moviesHandler{service: service}
}

// GetMovieList
func (srv *moviesHandler) GetMovieList(c *gin.Context) {

	// err := middleware.Claim(c)
	// if err != nil {
	// 	util.HandleError(c, err)
	// 	return
	// }

	resp, err := srv.service.GetMovieList(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}

// CreateMovie
func (srv *moviesHandler) CreateMovie(c *gin.Context) {
	resp, err := srv.service.CreateMovie(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}

// UpdateMovieByID
func (srv *moviesHandler) UpdateMovieByID(c *gin.Context) {
	resp, err := srv.service.UpdateMovieByID(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}
