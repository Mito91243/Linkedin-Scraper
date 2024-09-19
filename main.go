package main

import (
	"context"
	"database/sql"
	"flag"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"main/cmd"
	"main/config"
	"main/internal/web"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load("./.env")

	//Setting Dev Mode
	mode := flag.String("m", "prod", "Enviroment Mode")
	dsn := flag.String("db-dsn", os.Getenv("link_db_dsn"), "PostgreSQL DSN")
	maxOC := flag.Int("db-max-open-conns", 25, "PostgreSQL max open connections")
	maxIC := flag.Int("db-max-idle-conns", 25, "PostgreSQL max idle connections")
	maxIT := flag.String("db-max-idle-time", "15m", "PostgreSQL max connection idle time")
	flag.Parse()

	//Setting Loggers,client for the application
	app := &config.Application{
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		Client: &http.Client{
			Timeout: time.Second * 30,
		},
		DB: struct {
			Dsn          string
			MaxOpenConns int
			MaxIdleConns int
			MaxIdleTime  string
		}{
			Dsn:          *dsn,
			MaxOpenConns: *maxOC,
			MaxIdleConns: *maxIC,
			MaxIdleTime:  *maxIT,
		},
	}

	if err != nil {
		app.ErrorLog.Printf("Error Loading .env")
		return
	}

	db, err := openDB(app)
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
	// Defer a call to db.Close() so that the connection pool is closed before the
	// main() function exits.
	defer db.Close()
	// Also log a message to say that the connection pool has been successfully
	// established.
	app.InfoLog.Printf("database connection pool established")

	// Setting mode to launch while sending loggers to files
	if *mode == "prod" {
		web.Start()
	} else {
		cmd.Start(app)
	}
}

func openDB(app *config.Application) (*sql.DB, error) {
	// Use sql.Open() to create an empty connection pool, using the DSN from the config
	// struct.
	db, err := sql.Open("postgres", app.DB.Dsn)
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
