package main

import (
	_ "github.com/lib/pq"
	"main/cmd"
	"main/config"
	"main/internal/web"
)

func main() {
	app := config.InitializeConfig()
	if app.Mode == "prod" {
		web.Start(app)
	} else {
		cmd.Start(app)
	}
}