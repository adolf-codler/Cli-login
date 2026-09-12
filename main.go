package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main(){
	form_url:= "http://10.5.50.1/login"

	data := url.Values{}
	data.Set("username", "DH1151")
	data.Set("password", "1149")

	header:="application/x-www-form-urlencoded"

	resp, err := http.Post(form_url, header, strings.NewReader(data.Encode()))
	if err!=nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	log.Println("-> Logged in")
}
