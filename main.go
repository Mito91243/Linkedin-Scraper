package main

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"main/cmd"
	"main/config"
	"main/internal/web"
	"net/http"
	"os"
	"time"
)

func main() {

	//Setting Dev Mode
	mode := flag.String("m", "prod", "Enviroment Mode")
	flag.Parse()

	//Setting Loggers,client for the application
	app := &config.Application{
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		Client: &http.Client{
			Timeout: time.Second * 30,
		},
	}

	err := godotenv.Load("./.env")
	if err != nil {
		app.ErrorLog.Printf("Error Loading .env")
		return
	}

	// Setting mode to launch while sending loggers to files
	if *mode == "prod" {
		web.Start()
	} else {
		cmd.Start(app)
	}
}
