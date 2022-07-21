package main

import (
	"pxj/courseSystem/api"
	"pxj/courseSystem/config"
)

func main() {
	// Init config resource run server
	config.Init()
	err := api.InitResource()
	if err != nil {
		return
	}
	api.RunServer()
}
