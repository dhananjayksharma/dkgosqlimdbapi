package handlers

import (
	"imdbapi-service/internals/util"
	"imdbapi-service/pkg/v1/models/person"

	"github.com/gin-gonic/gin"
)

type PersonHandler interface {
	List(c *gin.Context)
	CreateUser(c *gin.Context)
	CreateUserBygRPC(c *gin.Context)
	CreateManyUserBygRPC(c *gin.Context)
}

type personHandler struct {
	service person.PersonService
}

func NewPersonHandler(service person.PersonService) PersonHandler {
	return &personHandler{service: service}
}

// List ...
func (srv *personHandler) List(c *gin.Context) {
	resp, err := srv.service.ListPerson(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}

// CreateUser ...
func (srv *personHandler) CreateUser(c *gin.Context) {
	resp, err := srv.service.CreatePerson(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}

// CreateUserBygRPC ...
func (srv *personHandler) CreateUserBygRPC(c *gin.Context) {
	resp, err := srv.service.CreateUserBygRPC(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}

// CreateManyUserBygRPC ...
func (srv *personHandler) CreateManyUserBygRPC(c *gin.Context) {
	resp, err := srv.service.CreateManyUserBygRPC(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}
