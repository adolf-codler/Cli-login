package cus_parser

import (
	"fmt"
	"os"
)

const (
	config=iota
	manual
	edit_config
)

func Parse()(int, error){
	argv := os.Args
	argc := len(argv)
	if argc==1{
		return config, nil
	}
	command := argv[1]
	if command=="--login"{
		return manual, nil
	} else if command=="--edit"{
		return edit_config, nil
	} else {
		return 0, fmt.Errorf("Invalid command")	
	}
}
