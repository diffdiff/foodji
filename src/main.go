package main

import (
	"github.com/diffdiff/foodji/app"
	"github.com/diffdiff/foodji/app/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
