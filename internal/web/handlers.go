package web

import (
	"encoding/json"
	"main/config"
	//"main/internal/models"
	"net/http"
)

func Home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//panic("oops! something went wrong")

		if r.URL.Path != "/" {
			notFound(w)
			return
		}
		w.Write([]byte("Zeby manga"))
	}
}

func Postsviewall(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // In production, set this to your frontend's origin
		//start := time.Now()
		company := r.URL.Query().Get("company")
		category := r.URL.Query().Get("category")
		keyword := r.URL.Query().Get("keyword")
		profiles := getAllProfiles(category, company, app)

		jsonData := getAllPosts(profiles, keyword, app)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(jsonData)
		if err != nil {
			// Log the error
			app.ErrorLog.Printf("Error writing response: %v", err)
			// Optionally, you could also set an error status code here
			serverError(app, w, err)
			return
		}

	}
}

func Profilesviewall(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // In production, set this to your frontend's origin
		company := r.URL.Query().Get("company")
		category := r.URL.Query().Get("category")
		profiles := getAllProfiles(category, company, app)
		app.ErrorLog.Printf("%v\n",app.DB.MaxOC)
		err := app.DB.Models.Profilesdb.InsertManyProfiles(profiles)
		if err != nil {
			app.ErrorLog.Fatal(err)
		}
		jsonData, err := json.Marshal(profiles)
		if err != nil {
			app.ErrorLog.Println("Error Marshalling to Json")
		}
		app.InfoLog.Printf("Profile Fetched :  %d", len(profiles))
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonData)
		if err != nil {
			// Log the error
			app.ErrorLog.Printf("Error writing response: %v", err)
			// Optionally, you could also set an error status code here
			serverError(app, w, err)
			return
		}
		//app.InfoLog.Printf("Profileviewall Request Finished in %.2f seconds :  \n", time.Since(start).Seconds())

	}
}
