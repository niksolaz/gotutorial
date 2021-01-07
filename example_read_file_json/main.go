package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Items struct {
	Users    []User    `json:"users"`
	Projects []Project `json:"projects"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

type Project struct {
	ID          int    `json:"projects_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	jsonfile, err := os.Open("info.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Json file opened")

	value, _ := ioutil.ReadAll(jsonfile)

	// we initialize our Items array
	var items Items

	json.Unmarshal(value, &items)

	for i := 0; i < len(items.Users); i++ {
		fmt.Println("User Type: " + items.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(items.Users[i].Age))
		fmt.Println("User Name: " + items.Users[i].Name)
		fmt.Println("Facebook Url: " + items.Users[i].Social.Facebook)
		fmt.Println("Twitter Url: " + items.Users[i].Social.Twitter)
	}

	for j := 0; j < len(items.Projects); j++ {
		fmt.Println("User ID: " + strconv.Itoa(items.Projects[j].ID))
		fmt.Println("User Title: " + items.Projects[j].Title)
		fmt.Println("User Description: " + items.Projects[j].Description)
	}

	defer jsonfile.Close()
}
