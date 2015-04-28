package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	ur := "https://pacific-tor-1459.herokuapp.com/data"
	fmt.Println("URL:>", ur)
	body := "Test"
	res, err := http.PostForm(ur, url.Values{"data": {body}})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
