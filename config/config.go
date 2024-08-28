package config

import (
	"log"
	"net/http"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Client   *http.Client
}