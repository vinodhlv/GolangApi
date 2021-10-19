package main

import (
	"dataApi/Config"
	"dataApi/Routes"
)

func main() {

	defer Config.DB.Close()
	Config.Connect()

	//setting up router and control flow
	route := Routes.SetUpRouter()

	//running
	route.Run()

}
