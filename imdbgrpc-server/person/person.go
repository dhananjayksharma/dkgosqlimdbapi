package person

import (
	"errors"
	"fmt"
	"strings"
)

type PersonData struct {
	ID    string
	Name  string
	Age   int32
	Email string
}

var personList = []PersonData{
	{
		ID:    "99",
		Name:  "ajay",
		Age:   23,
		Email: "ajay123@gmail.com",
	},
	{
		ID:    "88",
		Name:  "Dhananjay",
		Age:   35,
		Email: "Dhananjay334@gmail.com",
	},
	{
		ID:    "67",
		Name:  "Rishabh",
		Age:   27,
		Email: "Rishabh3g343@gmail.com",
	},
	{
		ID:    "10",
		Name:  "riya",
		Age:   24,
		Email: "riya@gmail.com",
	},
	{
		ID:    "99",
		Name:  "abhinav",
		Age:   28,
		Email: "abhinav@gmail.com",
	},
	{
		ID:    "25",
		Name:  "umesh",
		Age:   28,
		Email: "umesh@gmail.com",
	},
	{
		ID:    "47",
		Name:  "amit",
		Age:   28,
		Email: "amit@gmail.com",
	},
	{
		ID:    "40",
		Name:  "ramesh",
		Age:   28,
		Email: "ramesh787@gmail.com",
	},
}

// List
func List(key string) ([]PersonData, error) {
	var personResp = []PersonData{}
	for i, row := range personList {
		if strings.Contains(row.Name, strings.ToLower(key)) {
			fmt.Printf("Search found at id #%d, Name:%v\n", i, row.Name)
			personResp = append(personResp, row)
		}
	}

	if len(personResp) == 0 {
		errString := fmt.Sprintf("data not found for key:%s", key)
		return nil, errors.New(errString)
	}
	return personResp, nil
}

//ListById ...
func ListById(key string) (PersonData, error) {
	var personResp = PersonData{}
	for _, row := range personList {
		if row.ID == key {
			personResp = row
			return personResp, nil
		}
	}

	errString := fmt.Sprintf("data not found for key:%s", key)
	return personResp, errors.New(errString)

}
