package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"UserId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	completed bool   `json: copleted`
}

func performGetRequest() {
	myurl := "https://jsonplaceholder.typicode.com/todos/1"

	res, err := http.Get(myurl)
	if err != nil {
		fmt.Println("Error getting : ", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting Response : ", res.Status)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading: ", err)
		return
	}
	fmt.Println("Data : ", string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding : ", err)
		return
	}
	fmt.Println("Todo: ", todo)

}

func performPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "anshika",
		completed: true,
	}

	//convert the Todo struct to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling : ", err)
		return
	}

	//convert json data to string
	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myurl := "https://jsonplaceholder.typicode.com/todos"

	//send post request
	res, err := http.Post(myurl, "application/json", jsonReader)

	if err != nil {
		fmt.Println("Error sending request : ", err)
		return
	}

	defer res.Body.Close()

	// data, _ := ioutil.ReadAll(res.Body)
	// fmt.Println("Response : ", string(data))

	fmt.Println("Response status : ", res.Status)
}

func performUpdateRequest() {
	todo := Todo{
		UserID:    23899,
		Title:     "sontee",
		completed: true,
	}

	//convert the Todo struct to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling : ", err)
		return
	}

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	//convert json data to string
	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	//create PUT request
	req, err := http.NewRequest(http.MethodPut, myurl, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request : ", err)
		return
	}

	req.Header.Set("Content-type", "application/json")

	//send the request
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending the request : ", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response : ", string(data))

	fmt.Println("Response status : ", res.Status)
}

func performDeleteRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	//create DELETE request
	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("Error creating PUT request : ", err)
		return
	}

	//send the request
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending the request : ", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response status : ", res.Status)

}

func main() {
	fmt.Println("Now i am trying learn CRUD in goLang")
	// performGetRequest()
	//performPostRequest()
	//performUpdateRequest()
	//performDeleteRequest()

}
