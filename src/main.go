package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func ValidatePort() int{
	args := os.Args
	
	if len(args) != 2 {
		log.Fatal("A port should be specified")
	}
	port, err := strconv.Atoi(args[1])
	if err != nil || port < 1024 || port > 65535 {
		log.Fatal("A port should be numeric and between 1024 and 65535")
	}
	return port
}

func main(){
	port := ValidatePort()
	fmt.Println("Port: ", port)
}
