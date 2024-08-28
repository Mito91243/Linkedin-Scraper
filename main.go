package main

import (
	"flag"
	"log"
	"main/cmd"
	"main/server"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	//Setting Dev Mode
	mode := flag.String("mode", "prod", "Enviroment")
	flag.Parse()

	//Setting Loggers for the application
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Setting mode to launch while sending loggers to files
	if *mode == "dev" {
		server.Start(app.errorLog, app.infoLog)
	} else {
		cmd.Start(app.errorLog, app.infoLog)
	}
}
