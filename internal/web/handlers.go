package web

import (
	"net/http"
)

func (app *Application) Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//panic("oops! something went wrong")

		if r.URL.Path != "/" {
			app.notFound(w)
			return
		}
		w.Write([]byte("Zeby manga"))
	}
}

func (app *Application) Postsviewall() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (app *Application) Profilesviewall() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // In production, set this to your frontend's origin
		//start := time.Now()
		company := r.URL.Query().Get("company")
		category := r.URL.Query().Get("category")
		jsonData := app.getAllProfiles(category, company)
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		_, err := w.Write(jsonData)
		if err != nil {
			// Log the error
			app.ErrorLog.Printf("Error writing response: %v", err)
			// Optionally, you could also set an error status code here
			app.serverError(w, err)
			return
		}
		//app.InfoLog.Printf("Profileviewall Request Finished in %.2f seconds :  \n", time.Since(start).Seconds())

	}
}
