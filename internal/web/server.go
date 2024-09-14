package web

import (
	"log"
	"main/internal/models"
	"net/http"
	"os"
	"time"
)

type Application struct {
    ErrorLog *log.Logger
    InfoLog  *log.Logger
    Client   *http.Client
    DB       DatabaseConfig
}

type DatabaseConfig struct {
    Dsn          string
    MaxOpenConns int
    MaxIdleConns int
    MaxIdleTime  string
    Models       models.DbModels
}

func Start() {
	//Setting Loggers,client for the application
	app := &Application{
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		Client: &http.Client{
			Timeout: time.Second * 30,
		},
	}

	srv := &http.Server{
		Addr:     ":80",
		ErrorLog: app.ErrorLog,
		Handler:  app.routes(),
	}

	app.InfoLog.Printf("Starting server on 4040")
	err := srv.ListenAndServe()
	app.ErrorLog.Fatal(err)
}

func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home())
	mux.HandleFunc("/post/viewall", app.Postsviewall())
	mux.HandleFunc("/profiles/view", app.Profilesviewall())
	//log.Print(starting)
	return app.RecoverPanic(app.MWlogRequest(app.MWsecureHeaders(mux)))
}
