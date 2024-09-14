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

func (m ProfileModel) Insert(profile *ProfileRes) error {
	query := `
        INSERT INTO profiles (fullname, lastname, position, profileurn, link)
        VALUES ($1, $2, $3, $4, $5)`

	_, err := m.DB.Exec(query,
		profile.FullName,
		profile.LastName,
		profile.Position,
		profile.ProfileURN,
		profile.Link)

	return err
}

func (m ProfileModel) InsertMany(profiles []*ProfileRes) error {
	// Start a transaction
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails
	defer tx.Rollback()

	// Prepare the statement to be used for each insert
	stmt, err := tx.Prepare(`
        INSERT INTO profiles (fullname, lastname, position, email, link)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at, version`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Insert each profile
	for _, profile := range profiles {
		err := stmt.QueryRow(
			profile.FullName,
			profile.LastName,
			profile.Position,
			profile.ProfileURN,
			profile.Link,
		).Scan(&profile.ID, &profile.CreatedAt, &profile.Version)

		if err != nil {
			return err // This will trigger the deferred Rollback()
		}
	}

	// If we get here, all inserts were successful, so commit the transaction
	return tx.Commit()
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
