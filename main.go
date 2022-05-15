package main

import (
	"github.com/lgcavalheiro/shortcat/model"
	"github.com/lgcavalheiro/shortcat/server"
	"github.com/lgcavalheiro/shortcat/util"
)

func main() {
	util.ConfigEnv()
	model.Setup()
	server.Launch()
}
