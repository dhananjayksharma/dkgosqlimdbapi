package person

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"imdbapi-service/internals/adapter/mysql/entities"
	"imdbapi-service/internals/adapter/mysql/query"
	"imdbapi-service/internals/consts"
	"imdbapi-service/internals/util"
	"imdbapi-service/pkg/v1/models"
	"imdbapi-service/pkg/v1/models/request"
	"imdbapi-service/pkg/v1/models/response"

	"github.com/dhananjayksharma/dkgosqlimdbapi/imdbprotopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gin-gonic/gin"
)

type PersonService interface {
	ListPerson(c *gin.Context) (models.Response, error)
	CreatePerson(c *gin.Context) (models.Response, error)
	CreateUserBygRPC(c *gin.Context) (models.Response, error)
	CreateManyUserBygRPC(c *gin.Context) (models.Response, error)
}

type personService struct {
	db         query.MySQLDBStoreAccess
	grpcClient *grpc.ClientConn
}

func NewPersonService(db query.MySQLDBStoreAccess, grpcClient *grpc.ClientConn) PersonService {
	return &personService{db: db, grpcClient: grpcClient}
}

// CreateUserBygRPC ...
func (service personService) CreateUserBygRPC(c *gin.Context) (models.Response, error) {
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	var resp = models.Response{}

	var responsePerson []response.PersonResponse
	err := service.db.ListPerson(ctx, &responsePerson)
	if err != nil {
		return resp, err
	}

	var addUserRequest request.AddUserInputRequest
	if err := c.BindJSON(&addUserRequest); err != nil {
		return resp, &util.BadRequest{ErrMessage: err.Error()}
	}

	var status uint8
	status = uint8(consts.ActiveStatus)
	_ = entities.Person{
		IsActive:  &status,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		Email:     addUserRequest.Email,
		Age:       addUserRequest.Age,
		Name:      addUserRequest.Name,
	}

	personReq := imdbprotopb.RequestAddPerson{
		Name:   addUserRequest.Name,
		Email:  addUserRequest.Email,
		Age:    int32(addUserRequest.Age),
		Gender: 2,
	}
	fmt.Println("Starting Client...")
	connection, err := grpc.Dial("localhost:50059", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connection.Close()

	grpcClient := imdbprotopb.NewImdbServiceClient(connection)

	// ctxnew := context.Context(context.Background())
	respn, err := grpcClient.AddPerson(ctx, &personReq)
	if err != nil {
		log.Printf("Failed to add person : %v, err: %v", personReq, err)
	}

	// fmt.Println("Received from Server:", resp.Person.GetEmail(), resp.Person.GetAge(), resp.Person.GetGender())
	log.Printf("respn:%v", respn)
	resp.Data = respn
	return resp, nil
}

func (service personService) CreatePerson(c *gin.Context) (models.Response, error) {
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	var resp = models.Response{}

	var addUserRequest request.AddUserInputRequest
	if err := c.BindJSON(&addUserRequest); err != nil {
		return resp, &util.BadRequest{ErrMessage: err.Error()}
	}

	var status uint8
	status = uint8(consts.ActiveStatus)
	addUser := entities.Person{
		IsActive:  &status,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		Email:     addUserRequest.Email,
		Age:       addUserRequest.Age,
		Name:      addUserRequest.Name,
	}

	err := service.db.CreatePerson(ctx, &addUser)
	if err != nil {
		return resp, err
	}

	var newSpotlightMaster []response.PersonResponse
	newSpotlightMaster = append(newSpotlightMaster, response.PersonResponse{
		IsActive: &status,
		Name:     addUserRequest.Name,
		Age:      addUserRequest.Age,
		Mobile:   addUserRequest.Mobile,
		Email:    addUserRequest.Email,
	})

	resp.Data = newSpotlightMaster
	resp.Message = consts.UserAddedSuccess
	return resp, nil
}

func (srv *personService) ListPerson(c *gin.Context) (models.Response, error) {
	var err error
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var personData []response.PersonResponse
	var resp = models.Response{}

	skip_number, err := strconv.ParseUint(c.Query("skip"), 10, 64)
	if skip_number < 0 || err != nil {
		if err != nil {
			return resp, err
		}
		err = errors.New(consts.SkipMessage)
		return resp, err
	}

	page_limit, _ := strconv.ParseUint(c.Query("limit"), 10, 64)

	if page_limit < 1 {
		err = errors.New(consts.PageLimitMessage)
		return resp, err
	}

	// var queryParams = request.QueryMembersInputRequest{Code: code, Limit: int(page_limit), Skip: int(skip_number)}

	err = srv.db.ListPerson(ctx, &personData)
	if err != nil {
		return resp, err
	}
	var responsePerson []response.PersonResponse
	for _, row := range personData {
		responsePerson = append(responsePerson, response.PersonResponse{
			IsActive:    row.IsActive,
			Name:        row.Name,
			Email:       row.Email,
			Age:         row.Age,
			Createddate: row.Createddate,
		})
	}

	resp.Data = responsePerson
	return resp, nil
}

// CreateManyUserBygRPC ...
func (service personService) CreateManyUserBygRPC(c *gin.Context) (models.Response, error) {
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	fmt.Println("Starting Client...")
	connection, err := grpc.Dial("localhost:50059", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connection.Close()

	grpcClient := imdbprotopb.NewImdbServiceClient(connection)

	var resp = models.Response{}

	var addUserRequest []request.AddUserInputRequest
	if err := c.BindJSON(&addUserRequest); err != nil {
		return resp, &util.BadRequest{ErrMessage: err.Error()}
	}

	// var status uint8
	// status = uint8(consts.ActiveStatus)
	for _, row := range addUserRequest {
		gender := int(row.Gender)
		personReq := imdbprotopb.RequestAddPerson{
			Name:   row.Name,
			Email:  row.Email,
			Age:    int32(row.Age),
			Gender: imdbprotopb.Gender(gender),
		}

		respn, err := grpcClient.AddPerson(ctx, &personReq)
		if err != nil {
			log.Printf("Failed to add person : %v, err: %v", personReq, err)
		}

		// fmt.Println("Received from Server:", resp.Person.GetEmail(), resp.Person.GetAge(), resp.Person.GetGender())
		log.Printf("respn:%v\n", respn)
	}
	resp.Data = addUserRequest
	return resp, nil
}
