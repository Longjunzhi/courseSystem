package main

import (
	"pxj/courseSystem/api"
	"pxj/courseSystem/config"
)

func main() {
	// Init config resource run server
	config.Init()
	api.InitResource()
	api.RunServer()
}
