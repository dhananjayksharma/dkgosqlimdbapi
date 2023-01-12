package handlers

import (
	"time"

	"imdbapi-service/pkg/v1/models/movies"
	"imdbapi-service/pkg/v1/models/person"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(movieService movies.MovieService, personService person.PersonService) *gin.Engine {
	r := gin.Default()
	corsConfig := CORS()

	r.Use(corsConfig)
	healthHandler := NewHealthHandler()
	r.GET("/health", healthHandler.Health)

	// NewMovieHandler
	movieHandler := NewMovieHandler(movieService)
	// NewPersonHandler
	personHandler := NewPersonHandler(personService)
	{
		v1Group := r.Group("/movies")
		{
			// secured := v1Group.Group("/secured").Use(middleware.Auth())
			{
				v1Group.PUT("/movie/:code", movieHandler.UpdateMovieByID)
				v1Group.POST("/movies", movieHandler.CreateMovie)
				v1Group.GET("/movies", movieHandler.GetMovieList)
				// v1Group.GET("/person/:code", personHandler.)
				v1Group.GET("/person", personHandler.List)
				v1Group.POST("/person", personHandler.CreateUser)
				v1Group.POST("/person/grpc", personHandler.CreateUserBygRPC)
				v1Group.POST("/person/grpc/many", personHandler.CreateManyUserBygRPC)

			}

		}
	}
	return r
}

func CORS() gin.HandlerFunc {
	config := cors.Config{}
	config.AllowHeaders = []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"}
	// config.AllowAllOrigins = true
	config.AllowBrowserExtensions = true
	config.AllowCredentials = true
	config.AllowWildcard = true
	config.AllowOrigins = []string{"*"}
	config.MaxAge = time.Hour * 12
	return cors.New(config)
}
