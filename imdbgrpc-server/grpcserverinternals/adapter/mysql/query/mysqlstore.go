package query

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/dhananjayksharma/dkgosqlimdbapi/imdbgrpc-server/grpcserverinternals/adapter/mysql/entities"

	"github.com/dhananjayksharma/dkgosqlimdbapi/imdbgrpc-server/grpcserverinternals/consts"
	"github.com/dhananjayksharma/dkgosqlimdbapi/imdbgrpc-server/grpcserverinternals/request"
	"github.com/dhananjayksharma/dkgosqlimdbapi/imdbgrpc-server/grpcserverinternals/util"
	"gorm.io/gorm"
)

type mySQLDBStore struct {
	db *gorm.DB
}

func NewMySQLDBStore(db *gorm.DB) MySQLDBStoreAccess {
	return &mySQLDBStore{db: db}
}

type MySQLDBStoreAccess interface {
	GetMovieList(ctx context.Context, movieData *[]entities.MovieResponse) error
	CreateMovie(ctx context.Context, movieData *entities.Movie) error
	ListMovieByCode(ctx context.Context, movie *entities.Movie, code string) error
	UpdateMovieByID(ctx context.Context, user *entities.Movie, updateTypeData map[string]interface{}, code string) error

	CreatePerson(ctx context.Context, user *entities.Person) error
	ListPerson(ctx context.Context, user *[]entities.PersonResponse) error
	// ListPersonByName(ctx context.Context, user *[]entities.PersonResponse, queryParams request.QueryMembersInputRequest) error
}

// CreatePerson
func (ms *mySQLDBStore) CreatePerson(ctx context.Context, user *entities.Person) error {
	result := ms.db.Debug().WithContext(ctx).Create(&user)
	err := result.Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			_userMsg := fmt.Sprintf(consts.ErrUserAlreadyExists, user.Email)
			return &util.BadRequest{ErrMessage: _userMsg}
		} else {
			return &util.InternalServer{ErrMessage: err.Error()}
		}
	}

	return nil
}

// UpdateMovieByID
func (ms *mySQLDBStore) UpdateMovieByID(ctx context.Context, user *entities.Movie, updateTypeData map[string]interface{}, code string) error {

	var updateFields = make(map[string]interface{})
	for key, val := range updateTypeData {
		updateFields[key] = val
	}

	result := ms.db.Debug().WithContext(ctx).Model(&user).Where("moviecode=?", code).Omit("moviecode", "id").Updates(updateFields)

	log.Println("UpdateByID updated rows: ", result.RowsAffected)
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	} else if result.RowsAffected == 0 {
		err := fmt.Sprintf(consts.ErrorUpdateType, code)
		return &util.InternalServer{ErrMessage: err}
	}
	return nil
}

// CreateMovie
func (ms *mySQLDBStore) CreateMovie(ctx context.Context, movie *entities.Movie) error {
	result := ms.db.Debug().WithContext(ctx).Create(&movie)
	err := result.Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			_userMsg := fmt.Sprintf(consts.ErrMerchantAlreadyExists, movie.Moviecode)
			return &util.BadRequest{ErrMessage: _userMsg}
		} else {
			return &util.InternalServer{ErrMessage: err.Error()}
		}
	}

	return nil
}

// ListMovieDetail
func (ms *mySQLDBStore) ListMovieDetail(ctx context.Context, merchant *[]entities.MovieDetailsResponse, queryParams request.QueryMembersInputRequest) error {

	result := ms.db.Debug().WithContext(ctx).Model(&entities.MovieDetailsResponse{}).Select("person.name, person.email, person.mobile, person.is_active, person.created_at, movies.name as MovieName").Joins("left join merchants on movies.code = person.fk_code").Where("fk_code=?", queryParams.Code).Limit(queryParams.Limit).Offset(queryParams.Skip).Scan(&merchant)
	if result.RowsAffected == 0 {
		return &util.DataNotFound{ErrMessage: fmt.Sprintf(consts.ErrorDataNotFoundCode, queryParams.Code)}
	}
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}

// GetMovieList
func (ms *mySQLDBStore) GetMovieList(ctx context.Context, movieData *[]entities.MovieResponse) error {
	result := ms.db.WithContext(ctx).Model(&entities.MovieResponse{}).Select("moviecode,  name, is_active, created_at, updated_at").Find(&movieData)
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}
func (ms *mySQLDBStore) ListMovieByCode(ctx context.Context, movie *entities.Movie, code string) error {
	result := ms.db.WithContext(ctx).Model(&entities.MovieResponse{}).Select("moviecode,  name, is_active, created_at, updated_at").Find(&movie).Where("moviecode=?", code)
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}

//func (ms *mySQLDBStore) ListPerson(ctx context.Context, personData *[]entities.PersonResponse, queryParams request.QueryMembersInputRequest)
// ListPerson
func (ms *mySQLDBStore) ListPerson(ctx context.Context, personData *[]entities.PersonResponse) error {
	result := ms.db.WithContext(ctx).Model(&entities.PersonResponse{}).Select("id,  email, name, mobile, age, is_active, created_at, updated_at").Find(&personData)
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}
