package movies

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"imdbapi-service/internals/adapter/mysql/entities"
	"imdbapi-service/internals/adapter/mysql/query"
	"imdbapi-service/internals/consts"
	"imdbapi-service/internals/util"
	"imdbapi-service/pkg/v1/models"
	"imdbapi-service/pkg/v1/models/request"
	"imdbapi-service/pkg/v1/models/response"

	"github.com/gin-gonic/gin"
)

type MovieService interface {
	GetMovieList(c *gin.Context) (models.Response, error)
	CreateMovie(c *gin.Context) (models.Response, error)
	UpdateMovieByID(c *gin.Context) (models.Response, error)
}

type movieService struct {
	db query.MySQLDBStoreAccess
}

func NewMovieService(db query.MySQLDBStoreAccess) MovieService {
	return &movieService{db: db}
}

// UpdateMovieByID
func (service movieService) UpdateMovieByID(c *gin.Context) (models.Response, error) {
	var err error
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var movieData entities.Movie
	var resp = models.Response{}

	var updateMerchantRequest request.UpdateMovieInputRequest
	if err := c.BindJSON(&updateMerchantRequest); err != nil {
		return resp, &util.BadRequest{ErrMessage: err.Error()}
	}

	code := strings.Trim(c.Param("code"), "")
	if len(code) == 0 {
		err = errors.New(consts.InvalidCode)
		return resp, err
	}

	var responseMovie []response.MovieResponse
	err = service.db.GetMovieList(ctx, &responseMovie)
	if err != nil {
		return resp, err
	}
	if len(responseMovie) == 0 {
		err = errors.New(fmt.Sprintf(consts.ErrorDataNotFoundCode, code))
		return resp, err
	}

	var updateTypeData = make(map[string]interface{})
	updateTypeData["name"] = updateMerchantRequest.Name

	err = service.db.UpdateMovieByID(ctx, &movieData, updateTypeData, code)
	if err != nil {
		return resp, err
	}

	err = service.db.GetMovieList(ctx, &responseMovie)
	if err != nil {
		return resp, err
	}
	resp.Data = responseMovie
	resp.Message = fmt.Sprintf(consts.MovieUpdatedSuccess, code)
	return resp, nil
}

// CreateMovie
func (srv movieService) CreateMovie(c *gin.Context) (models.Response, error) {
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var resp = models.Response{}
	var addMerchantRequest request.AddMovieInputRequest
	if err := c.BindJSON(&addMerchantRequest); err != nil {
		return resp, &util.BadRequest{ErrMessage: err.Error()}
	}

	currTime := time.Now()
	var status uint8
	status = uint8(consts.ActiveStatus)
	addMerchant := entities.Movie{
		UpdatedAt: currTime,
		CreatedAt: currTime,
		Moviecode: addMerchantRequest.Code,
		Status:    &status,
		Name:      addMerchantRequest.Name,
	}

	err := srv.db.CreateMovie(ctx, &addMerchant)
	if err != nil {
		return resp, err
	}

	var newMerchantMaster []response.MovieResponse
	newMerchantMaster = append(newMerchantMaster, response.MovieResponse{
		Name:      addMerchantRequest.Name,
		Moviecode: addMerchantRequest.Code,
		CreatedAt: currTime.String(),
	})

	resp.Data = newMerchantMaster
	resp.Message = consts.MovieAddedSuccess
	return resp, nil
}

// GetMovieList movies
func (srv *movieService) GetMovieList(c *gin.Context) (models.Response, error) {
	var err error
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var movieData []response.MovieResponse
	var resp = models.Response{}
	if err != nil {
		return resp, err
	}

	err = srv.db.GetMovieList(ctx, &movieData)
	if err != nil {
		return resp, err
	}

	var outMSM []response.MovieResponse
	for _, row := range movieData {
		outMSM = append(outMSM, response.MovieResponse{
			Name:      row.Name,
			Moviecode: row.Moviecode,
			CreatedAt: row.CreatedAt,
		})
	}
	resp.Data = outMSM
	return resp, nil
}
