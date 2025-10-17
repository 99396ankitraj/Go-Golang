package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("I am trying to learn how URL actually works")
	myURL := "https://jsonplaceholder.typicode.com/todos/1"
	fmt.Printf("Type of URL: %T\n", myURL)

	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("can't parse URL ", err)
		return
	}
	fmt.Printf("Type of URL : %T\n", parsedURL)

	fmt.Println("Schema: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("RawQuery: ", parsedURL.RawQuery)

	//Modifying URL components
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=ankitraj12"

	//constructing a URL string from a URL object
	newUrl := parsedURL.String()

	fmt.Println("New URL: ", newUrl)
}
