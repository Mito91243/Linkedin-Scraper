package web

import (
	"main/config"
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
	mux.HandleFunc("/", Home(app))
	mux.HandleFunc("/post/viewall", Postsviewall())
	mux.HandleFunc("/profiles/view", Profilesviewall(app))
	//log.Print(starting)
	return RecoverPanic(app)(MWlogRequest(app)(MWsecureHeaders(app)(mux)))
}
