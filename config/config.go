package config

import (
	//"fmt"
	"log"
	"net/http"
	//"runtime/debug"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Client   *http.Client
	DB       struct {
		Dsn          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}
}
