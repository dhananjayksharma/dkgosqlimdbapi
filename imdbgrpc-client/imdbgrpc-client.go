package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/dhananjayksharma/dkgosqlimdbapi/imdbprotopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr     = flag.String("addr", "localhost:50059", "the address to connect to")
	gRPCConn *grpc.ClientConn
)

func main() {
	flag.Parse()
	gRPCConn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect server on port :%v, err:%v", addr, err)
	}

	defer gRPCConn.Close()

	addPerson(gRPCConn) // Unary
	//searchPerson(gRPCConn) // Server Streaming
	//addManyPerson(gRPCConn) // Client Streaming
	//searchPersonById(gRPCConn) // Bidirectional Server Streaming
}

// searchPersonById ...
func searchPersonById(grpcClient *grpc.ClientConn) {
	serviceClientConn := imdbprotopb.NewImdbServiceClient(grpcClient)

	stream, err := serviceClientConn.ListPersonById(context.Background())
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	reqList := []int32{40, 47, 25, 99, 88}
	fmt.Println("len(reqList):", len(reqList))
	for i, key := range reqList {
		time.Sleep(1 * time.Second)
		log.Printf("Streaming Client Search key sent %d: %v", i, key)
		req := &imdbprotopb.SearchPersonRequest{Query: fmt.Sprintf("%d", key)}
		if err := stream.Send(req); err != nil {
			log.Fatalf("can not send %v", err)
		}

		if i == len(reqList)-1 {
			if err := stream.CloseSend(); err != nil {
				log.Println(err)
			}
		}

		resp, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("can not receive %v", err)
		}

		log.Printf("Streaming Client Response received %v ", resp)
	}

}

// searchPerson ...
func searchPerson(grpcClient *grpc.ClientConn) {
	serviceClientConn := imdbprotopb.NewImdbServiceClient(grpcClient)
	ctx := context.Background()
	req := &imdbprotopb.SearchPersonRequest{Query: "m"}
	stream, err := serviceClientConn.ListPerson(ctx, req)
	if err != nil {
		log.Fatalf("Error while calling ListPerson: Conn: %v, err:%v", serviceClientConn, err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Println("Received all!")
			return
		}

		if err != nil {
			log.Println("Error while streaming %v", err)
		}

		fmt.Println("List Received:", res)
	}

	fmt.Println("End of listPerson")
}

// getRandomChar()
func getRandomChar() string {
	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyz"
	c := charset[rand.Intn(len(charset))]
	return string(c)
}

// addManyPerson ...
func addManyPerson(grpcClient *grpc.ClientConn) {
	serviceClientConn := imdbprotopb.NewImdbServiceClient(grpcClient)
	// personReq := []imdbprotopb.RequestAddPerson{ //Option one
	personReq := []*imdbprotopb.RequestAddPerson{
		{
			Name:   "user_" + getRandomChar() + "_" + getRandomChar(),
			Email:  getRandomChar() + "_" + getRandomChar() + "_user@gmail.com",
			Age:    23,
			Gender: 2,
		},
		{
			Name:   "user_" + getRandomChar() + "_" + getRandomChar(),
			Email:  getRandomChar() + "_" + getRandomChar() + "_user@gmail.com",
			Age:    23,
			Gender: 2,
		},
		{
			Name:   "user_" + getRandomChar() + "_" + getRandomChar(),
			Email:  getRandomChar() + "_" + getRandomChar() + "_user@gmail.com",
			Age:    23,
			Gender: 2,
		},
	}
	clientStream, err := serviceClientConn.AddManyPerson(context.Background())
	if err != nil {
		log.Fatalf("%v.AddManyPerson(_) = _, %v", serviceClientConn, err)
	}

	_ = serviceClientConn
	for i, person := range personReq {
		time.Sleep(1 * time.Second)
		//req := &imdbprotopb.RequestAddPerson{Name: person.Name, Email: person.Email, Gender: person.Gender, Age: person.Age}
		if err := clientStream.Send(person); err != nil {
			log.Printf("Failed to send add person : %v, err: %v", person, err)
		}
		fmt.Println("Sending person #%v, req:%v ", i, person.GetEmail())
	}

	resp, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", clientStream, err, nil)
	}
	log.Printf("Number of Person added by Server: Count: %v", resp)
}

// addPerson ..
func addPerson(grpcClient *grpc.ClientConn) {
	serviceClientConn := imdbprotopb.NewImdbServiceClient(grpcClient)

	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyz"
	c := charset[rand.Intn(len(charset))]

	personReq := &imdbprotopb.RequestAddPerson{
		Name:   "user_" + string(c),
		Email:  string(c) + "_user@gmail.com",
		Age:    23,
		Gender: 2,
	}

	ctx := context.Context(context.Background())
	resp, err := serviceClientConn.AddPerson(ctx, personReq)
	if err != nil {
		log.Printf("Failed to add person : %v, err: %v", personReq, err)
	}

	fmt.Println("Received from Server:", resp.Person.GetId(), resp.Person.GetEmail(), resp.Person.GetAge(), resp.Person.GetGender())
}
