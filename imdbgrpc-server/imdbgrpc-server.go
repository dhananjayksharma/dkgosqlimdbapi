package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/dhananjayksharma/dkgosqlimdbapi/imdbgrpc-server/grpcserverinternals/adapter/mysql"
	"github.com/dhananjayksharma/dkgosqlimdbapi/imdbgrpc-server/grpcserverinternals/adapter/mysql/entities"
	"github.com/dhananjayksharma/dkgosqlimdbapi/imdbgrpc-server/grpcserverinternals/adapter/mysql/query"
	"github.com/dhananjayksharma/dkgosqlimdbapi/imdbprotopb"
	"google.golang.org/grpc"

	"github.com/dhananjayksharma/dkgosqlimdbapi/imdbgrpc-server/person"

	"github.com/spf13/viper"
)

type imdbServiceServer struct {
	imdbprotopb.UnimplementedImdbServiceServer
	db query.MySQLDBStoreAccess
}

func (iss *imdbServiceServer) ListPersonById(stream imdbprotopb.ImdbService_ListPersonByIdServer) error {
	for {
		req, err := stream.Recv()
		time.Sleep(1 * time.Second)
		if err == io.EOF {
			fmt.Println("End of receive from client")
			break
		}

		if err != nil {
			fmt.Println("Error while receive from client %v", err)
		}

		fmt.Printf("Server received searching request for key: %v\n", req.Query)
		resp, err := person.ListById(req.Query)

		if err != nil {
			fmt.Printf("No data for key %v\n", req.Query)
		}

		fmt.Printf("Streaming Server Response: Result found for key %v, Data:%v\n", req.Query, resp)

		respData := &imdbprotopb.Person{Id: resp.ID, Name: resp.Name, Age: resp.Age, Email: resp.Email}

		if err := stream.Send(respData); err != nil {
			log.Printf("send error %v", err)
		}

		if err != nil {
			log.Printf("SERVER: error sending data:  error %v", err)
			return err
		}
	}

	return nil
}

//ListPerson server streaming
func (iss *imdbServiceServer) ListPerson(req *imdbprotopb.SearchPersonRequest, stream imdbprotopb.ImdbService_ListPersonServer) error {

	personList, err := person.List(req.Query)
	fmt.Println("Server data: ", personList)
	if err != nil {
		fmt.Println("Error while fetching person list")
		return nil
	}

	for i, row := range personList {
		time.Sleep(1 * time.Second)
		resp := &imdbprotopb.Person{Email: row.Email, Name: row.Name, Age: row.Age, Id: string(row.ID)}
		fmt.Printf("Serve streaming Data:%d, Resp:%v\n", i, resp)
		err := stream.Send(resp)
		if err != nil {
			fmt.Println("Error while sending server streaming")
		}
	}
	log.Println("Completed")
	return nil
}

// AddManyPerson
func (iss *imdbServiceServer) AddManyPerson(stream imdbprotopb.ImdbService_AddManyPersonServer) error {
	fmt.Printf("receiving...\n\n")
	ctx := context.Background()

	var personCount int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&imdbprotopb.ResponseAddManyPerson{Count: personCount})
		}
		if err != nil {
			return err
		}
		fmt.Println("Person req:", req)
		requestPerson := entities.Person{
			Email: req.Email,
			Name:  req.Name,
			Age:   int8(req.Age),
		}

		err = iss.db.CreatePerson(ctx, &requestPerson)
		if err != nil {
			log.Printf("create person failed :%v", err)
		}
		personCount++
	}
	return nil
}

// AddPerson
func (iss *imdbServiceServer) AddPerson(ctx context.Context, req *imdbprotopb.RequestAddPerson) (*imdbprotopb.ResponseAddPerson, error) {
	fmt.Printf("AddPerson received: %v\n\n", req)

	requestPerson := entities.Person{
		Email: req.Email,
		Name:  req.Name,
		Age:   int8(req.Age),
	}

	err := iss.db.CreatePerson(ctx, &requestPerson)
	if err != nil {
		log.Printf("create person failed :%v", err)
	}
	resp := &imdbprotopb.ResponseAddPerson{
		Person: &imdbprotopb.Person{Id: string(requestPerson.ID), Name: req.Name, Email: req.Email, Age: req.Age, Gender: imdbprotopb.Gender(req.Gender.Number())},
	}
	fmt.Printf("Sent AddPerson : %v\n\n", resp)

	return resp, nil
}

var portName = "localhost:50059"

func startgRPCServer() {
	// Set the file name of the configurations file
	if os.Getenv("MICROSERVICECDEMONEWAPI") == "local" {
		viper.SetConfigName("config-local")
	} else {
		viper.SetConfigName("config")
	}

	log.Println("Current Config :", os.Getenv("MICROSERVICECDEMONEWAPI"))

	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	dbReadWrite := viper.GetString("ENV_VAR_RW")
	dbConnection, err := mysql.DBConn(dbReadWrite)
	if err != nil {
		log.Fatalf("MySQL connection error , %v", err)
	} else {
		fmt.Println("dbConnection connected")
	}

	db := query.NewMySQLDBStore(dbConnection)

	lis, err := net.Listen("tcp", portName)
	if err != nil {
		log.Fatalf("failed to listen on port :%v, err:%v", portName, err)
	}

	grcpServer := grpc.NewServer()
	imdbprotopb.RegisterImdbServiceServer(grcpServer, &imdbServiceServer{db: db})
	fmt.Printf("Server start listing on port %v \n", lis.Addr())
	if err := grcpServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func main() {
	fmt.Println("Server start gRPC ")
	startgRPCServer()
}
