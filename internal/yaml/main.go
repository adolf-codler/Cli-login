package cus_yaml

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)


const (
	config=iota
	manual
	edit_config
)

type Entries struct{
	Forms []Form `yaml:"forms"`
}

type Form struct{
	Name string `yaml:"name"`
	Url string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func WriteConfig()error{// {{{
	file, err:=os.Create("config.yaml")
	if err!=nil{
		return fmt.Errorf("File creating Error: %w", err)
	}
	var form Form
	var entry Entries

	fmt.Println("Enter name of the form")
	fmt.Scanln(&form.Name)
	fmt.Println("Enter url of the form")
	fmt.Scanln(&form.Url)
	fmt.Println("Enter username of the form")
	fmt.Scanln(&form.Username)
	fmt.Println("Enter password of the form")
	fmt.Scanln(&form.Password)

	entry.Forms = append(entry.Forms, form)
	data, err:=yaml.Marshal(entry)
	if err!=nil{
		return fmt.Errorf("Marshal Error: %w", err)
	}

	_, err=file.Write(data)
	if err!=nil{
		return fmt.Errorf("Wrting Error: %w", err)
	}
	return nil
}// }}}

func ReadConfig()(string, string, error){
	
}

func EditConfig(){

}

