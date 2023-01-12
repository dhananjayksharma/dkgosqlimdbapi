package query

import (
	"context"

	"imdbapi-service/internals/adapter/mysql/entities"
	"imdbapi-service/pkg/v1/models/response"
)

type MockMySQLDBStore struct {
}

var _ MySQLDBStoreAccess = (*MockMySQLDBStore)(nil)

// UpdateMovieByID
func (ms *MockMySQLDBStore) UpdateMovieByID(ctx context.Context, user *entities.Movie, updateTypeData map[string]interface{}, code string) error {
	return nil
}

// ListMovieByCode
func (ms *MockMySQLDBStore) ListMovieByCode(ctx context.Context, user *entities.Movie, code string) error {
	return nil
}

// CreatePersonMember
func (ms *MockMySQLDBStore) CreatePersonMember(ctx context.Context, user *entities.Person) error {
	return nil
}

// GetMovieList
func (ms *MockMySQLDBStore) GetMovieList(ctx context.Context, movieData *[]response.MovieResponse) error {
	data := []response.MovieResponse{
		{
			Moviecode: "1454dddd",
			Name:      "TestMerchant",
			CreatedAt: "2022-06-04 16:40:28",
		}, {
			Moviecode: "1454dddd",
			Name:      "TestMerchant",
			CreatedAt: "2022-06-04 16:40:28",
		},
	}
	*movieData = data
	return nil
}

//func (ms *MockMySQLDBStore) ListPerson(ctx context.Context, person *[]response.PersonResponse, queryParams request.QueryMembersInputRequest)
// ListPerson
func (ms *MockMySQLDBStore) ListPerson(ctx context.Context, person *[]response.PersonResponse) error {
	data := []response.PersonResponse{
		{
			Name: "Ranbir",
		}, {
			Name: "Akshay",
		},
	}
	*person = data
	return nil
}

// CreateMovie
func (ms *MockMySQLDBStore) CreateMovie(ctx context.Context, user *entities.Movie) error {
	var data = entities.Movie{
		Name:      "Avatar: The Way of Water",
		Moviecode: "cadjq02gqpmvljdra98",
	}
	*user = data
	return nil
}

// CreatePerson
func (ms *MockMySQLDBStore) CreatePerson(ctx context.Context, user *entities.Person) error {
	var data = entities.Person{
		Name:  "Ajay",
		Email: "Ajayd@gmail.com",
	}
	*user = data
	return nil
}
