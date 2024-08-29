package server

import (
	//"log"
	"main/config"
	"main/internal/web"
	"net/http"
)

func Start(app *config.Application) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", web.Home(app))
	mux.HandleFunc("/post/viewall", web.Postsviewall(app))
	mux.HandleFunc("/profile/viewall", web.Profilesviewall(app))
	//log.Print(starting)

	srv := &http.Server{
		Addr: ":4040",
		ErrorLog: app.ErrorLog,
		Handler: mux,
		}
		app.InfoLog.Printf("Starting server on 4040")
		err := srv.ListenAndServe()
		app.ErrorLog.Fatal(err)
	}

