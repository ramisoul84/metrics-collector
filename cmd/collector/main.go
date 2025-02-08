package main

import "fmt"

func main() {
	config, err := ReadConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println("token= ", config.UserToken)
	fmt.Println("host= ", config.TargetHost)

}
