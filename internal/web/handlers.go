package web

import (
	//"encoding/json"
	"main/config"
	"net/http"
	"time"
	//"main/cmd"
)

func Home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			notFound(w)
			return
		}
		w.Write([]byte("Zeby manga"))
	}
}

func Postsviewall(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func Profilesviewall(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		jsonData := getAllProfiles("18","microsoft",app)
		w.Header().Set("Content-Type","application/json")

		w.WriteHeader(http.StatusOK)
		_,err := w.Write(jsonData)
		if err != nil {
			// Log the error
			app.ErrorLog.Printf("Error writing response: %v", err)
			// Optionally, you could also set an error status code here
			serverError(w,err,app)
			return
		}
		app.InfoLog.Printf("Profileviewall Request Finished in %.2f seconds\n", time.Since(start).Seconds())

	}
}