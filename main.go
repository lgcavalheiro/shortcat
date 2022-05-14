package main

import (
	"github.com/lgcavalheiro/shortcat/model"
	"github.com/lgcavalheiro/shortcat/server"
)

func main() {
	model.Setup()
	server.Launch()
}
