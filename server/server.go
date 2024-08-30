package server

import (
	//"log"
	"main/config"
	"main/internal/web"
	"net/http"
)

func Start(app *config.Application) {

	srv := &http.Server{
		Addr:     ":80",
		ErrorLog: app.ErrorLog,
		Handler:  routes(app),
	}

	app.InfoLog.Printf("Starting server on 4040")
	err := srv.ListenAndServe()
	app.ErrorLog.Fatal(err)
}

func routes(app *config.Application) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", web.Home(app))
	mux.HandleFunc("/post/viewall", web.Postsviewall(app))
	mux.HandleFunc("/profiles/view", web.Profilesviewall(app))
	//log.Print(starting)
	return web.MWlogRequest(web.MWsecureHeaders(mux))
}
