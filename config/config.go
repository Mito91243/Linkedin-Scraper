package config

import (
	//"fmt"
	"context"
	"database/sql"
	"flag"
	"github.com/joho/godotenv"
	"log"
	"main/internal/models"
	"net/http"
	"os"
	"time"
	//"runtime/debug"
)

type Application struct {
	Mode     string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Client   *http.Client
	DB       DBconfig
}

type DBconfig struct {
	Dsn    string
	MaxOC  int
	MaxIC  int
	MaxIT  string
	Models models.DbModels
}

func InitializeConfig() *Application {
	err := godotenv.Load("./.env")

	// Setting Dev Mode
	mode := flag.String("m", "prod", "Enviroment Mode")
	dsn := flag.String("db-dsn", os.Getenv("link_db_dsn"), "PostgreSQL DSN")
	maxOC := flag.Int("db-max-open-conns", 25, "PostgreSQL max open connections")
	maxIC := flag.Int("db-max-idle-conns", 25, "PostgreSQL max idle connections")
	maxIT := flag.String("db-max-idle-time", "15m", "PostgreSQL max connection idle time")
	flag.Parse()

	configs := &DBconfig{
		Dsn:   *dsn,
		MaxOC: *maxOC,
		MaxIC: *maxIC,
		MaxIT: *maxIT,
	}

	app := &Application{
		Mode:     *mode,
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		Client: &http.Client{
			Timeout: time.Second * 30,
		},
		DB: *configs,
	}

	if err != nil {
		app.ErrorLog.Printf("Error Loading .env")
		return nil
	}

	db, err := openDB()

	if err != nil {
		app.ErrorLog.Fatal(err)
	}

	// Defer a call to db.Close() so that the connection pool is closed before the
	// main() function exits.
	defer db.Close()
	// Also log a message to say that the connection pool has been successfully
	// established.
	app.InfoLog.Printf("database connection pool established")

	return app
}

func openDB() (*sql.DB, error) {
	// Use sql.Open() to create an empty connection pool, using the DSN from the config
	// struct.
	db, err := sql.Open("postgres", os.Getenv("link_db_dsn"))
	if err != nil {
		return nil, err
	}
	// Create a context with a 5-second timeout deadline.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Use PingContext() to establish a new connection to the database, passing in the
	// context we created above as a parameter. If the connection couldn't be
	// established successfully within the 5 second deadline, then this will return an
	// error.
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	// Return the sql.DB connection pool.
	return db, nil
}
