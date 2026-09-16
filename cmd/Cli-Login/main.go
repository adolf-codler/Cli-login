package main

import (
	"adolf-codler/cli-login/internal/parser"
	cus_yaml "adolf-codler/cli-login/internal/yaml"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)


const (
	config=iota
	manual
	edit_config
)

func main(){
	proceed, err:=cus_parser.Parse()
	if err!=nil{
		log.Fatalln(err)
	}
	var username, password string
	if proceed == config{
		username, password, err = cus_yaml.ReadConfig()
	} else if proceed == edit_config{
		cus_yaml.EditConfig()
	} else{
		username, password, err=InputIdPass()
		if err!=nil{
			log.Fatalln(err)
		}
	}
	form_url:= "http://10.5.50.1/login"

	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)

	header:="application/x-www-form-urlencoded"

	resp, err := http.Post(form_url, header, strings.NewReader(data.Encode()))
	if err!=nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	log.Println("-> Logged in")
}
