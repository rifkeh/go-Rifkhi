package main

import (
	"competence/config"
	"competence/route"
)

func main(){
	if err := config.InitDB(); err != nil{
		panic(err)
	}
	e := route.New()

	e.Logger.Fatal(e.Start(":8080"))
}