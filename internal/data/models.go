package data

import (
	"database/sql" // New import
	"errors"
	"fmt"
	"main/internal/models"
	//"time"
)

// Define a MovieModel struct type which wraps a sql.DB connection pool.
type ProfileModel struct {
	DB *sql.DB
}
type CompanyModel struct {
	DB *sql.DB
}
type DbModels struct {
	Profilesdb ProfileModel
	Companydb CompanyModel
}

func (m ProfileModel) Insert(profile *models.ProfileRes) error {
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
func (m ProfileModel) InsertManyProfiles(profiles []models.ProfileRes) error {
    if m.DB == nil {
        return errors.New("database connection is nil")
    }

    // Check if the database is open
    if err := m.DB.Ping(); err != nil {
        return fmt.Errorf("database connection error: %w", err)
    }

    tx, err := m.DB.Begin()
    if err != nil {
        return fmt.Errorf("failed to begin transaction: %w", err)
    }
    defer tx.Rollback()

    // Prepare statement for inserting company
    companyStmt, err := tx.Prepare(`
    INSERT INTO companies (company_identifier)
    VALUES ($1)
    ON CONFLICT (company_identifier) DO NOTHING`)
    if err != nil {
        return fmt.Errorf("failed to prepare company statement: %w", err)
    }
    defer companyStmt.Close()

    // Prepare statement for inserting profile
    profileStmt, err := tx.Prepare(`
    INSERT INTO profiles (full_name, position, email, category, link, company_identifier, version)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING id, created_at, version`)
    if err != nil {
        return fmt.Errorf("failed to prepare profile statement: %w", err)
    }
    defer profileStmt.Close()

    for i, profile := range profiles {
        // Insert company (this will do nothing if it already exists)
        _, err := companyStmt.Exec(profile.CompanyID)
        if err != nil {
            return fmt.Errorf("failed to insert company for profile %d: %w", i, err)
        }

        // Insert profile
        err = profileStmt.QueryRow(
            profile.FullName,
            profile.Position,
            profile.ProfileURN,  // Assuming this is the email
            profile.Category,
            profile.Link,
            profile.CompanyID,
            1,  // Default version
        ).Scan(&profile.ID, &profile.CreatedAt, &profile.Version)

        if err != nil {
            return fmt.Errorf("failed to insert profile %d: %w", i, err)
        }
    }

    if err := tx.Commit(); err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    }

    return nil
}

// Add a placeholder method for fetching a specific record from the movies table.
func (m ProfileModel) Get(id int) (*models.ProfileRes, error) {

	return nil, nil
}
func (m ProfileModel) Delete(id int64) error {
	return nil
}

// For ease of use, we also add a New() method which returns a Models struct containing
// the initialized MovieModel.
func NewModels(db *sql.DB) DbModels {
	if db == nil {
		fmt.Printf("DB NOT INTIALIZED IN NEW MODELS")
	}
	return DbModels{
		Profilesdb: ProfileModel{DB: db},
		Companydb: CompanyModel{DB: db},
	}
}

//func (m ProfileModel) InsertMany(ProfileRes *[]ProfileRes)
/* Add a placeholder method for updating a specific record in the movies table.
func (m ProfileModel) Update(movie *Movie) error {
return nil
}*/
// Add a placeholder method for deleting a specific record from the movies table.
