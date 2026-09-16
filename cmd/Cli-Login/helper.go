package main

import(
	"bufio"
	"os"
	"strings"
	"fmt"
)


func InputIdPass()(string, string, error){
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Username:")
	username, err:=reader.ReadString('\n')
	if err!=nil{
		return "", "", fmt.Errorf("ReadString Error: %w", err)
	}
	username= strings.TrimSpace(username)
	fmt.Println("Password:")
	password, err:=reader.ReadString('\n')
	if err!=nil{
		return "", "", fmt.Errorf("ReadString Error: %w", err)
	}
	password= strings.TrimSpace(password)
	return username, password, nil
}
