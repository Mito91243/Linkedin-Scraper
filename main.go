package main

import (
	"flag"
	"log"
	"main/cmd"
	"main/config"
	"main/server"
	"os"
	"time"
	"net/http"
)



func main() {
	//Setting Dev Mode
	mode := flag.String("mode", "prod", "Enviroment")
	flag.Parse()

	//Setting Loggers,client for the application
	app := &config.Application{
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		Client: &http.Client{
            Timeout: time.Second * 30,
        },
	}

	// Setting mode to launch while sending loggers to files
	if *mode == "dev" {
		server.Start(app)
	} else {
		cmd.Start(app)
	}
}
