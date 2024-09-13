package models

import (
	"database/sql" // New import
	//"time"
)

// Define a MovieModel struct type which wraps a sql.DB connection pool.
type ProfileModel struct {
	DB *sql.DB
}
type DbModels struct {
	Profilesdb ProfileModel
}

// Add a placeholder method for inserting a new record in the movies table.
func (m ProfileModel) Insert(profile *ProfileRes) error {
	return nil
}

// Add a placeholder method for fetching a specific record from the movies table.
func (m ProfileModel) Get(id int) (*ProfileRes, error) {
	return nil, nil
}
func (m ProfileModel) Delete(id int64) error {
	return nil
}

// For ease of use, we also add a New() method which returns a Models struct containing
// the initialized MovieModel.
func NewModels(db *sql.DB) DbModels {
	return DbModels{
		Profilesdb: ProfileModel{DB: db}, 
	}
}

//func (m ProfileModel) InsertMany(ProfileRes *[]ProfileRes)
/* Add a placeholder method for updating a specific record in the movies table.
func (m ProfileModel) Update(movie *Movie) error {
return nil
}*/
// Add a placeholder method for deleting a specific record from the movies table.
