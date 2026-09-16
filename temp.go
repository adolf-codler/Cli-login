package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	//"gopkg.in/yaml.v3"
)

type Entries struct{
	Name[]Form `yaml:"name"`
}

type Form struct{
	Url string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}


func main(){
	WriteConfig()
}

func WriteConfig()error{
	var form Form
	fmt.Println("Enter url of the form")
	fmt.Scanln(&form.Url)
	fmt.Println("Enter username of the form")
	fmt.Scanln(&form.Username)
	fmt.Println("Enter password of the form")
	fmt.Scanln(&form.Password)

	var entry Entries
	entry.Name= append(entry.Name, form)
	
	file, err:=os.Create("data.yaml")
	if err!=nil{
		return fmt.Errorf("creating error: %w", err)
	}
	data, err:=yaml.Marshal(entry)
	_, err=file.Write(data)
	if err!=nil{
		return fmt.Errorf("Writing error: %w", err)
	}
	return nil
}
