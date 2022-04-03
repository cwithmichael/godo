package main

import (
	"database/sql"
	"flag"
	"github.com/cwithmichael/godo/pkg/models"
	"github.com/cwithmichael/godo/pkg/models/mysql"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

  _ "github.com/mattn/go-sqlite3"
	"github.com/golangcollege/sessions"
)

type contextKey string

const (
	contextKeyIsAuthenticated = contextKey("isAuthenticated")
	retries                   = 10
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	session  *sessions.Session
	todos    interface {
		Insert(string, string, int) (int, error)
		Get(int) (*models.Todo, error)
		Latest(int) ([]*models.Todo, error)
		Update(int, string, string, bool) error
		Delete(int) error
	}
	templateCache map[string]*template.Template
	users         interface {
		Insert(string, string, string) error
		Authenticate(string, string) (int, error)
		Get(int) (*models.User, error)
	}
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
  dsn := flag.String("dsn", "./test.db", "Sqlite data source name")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	db, err := openDB(*dsn)
	if err != nil {
		for i := 0; i < retries; i++ {
			db, err = openDB(*dsn)
			if err == nil {
				break
			}
			time.Sleep(3 * time.Second)
		}
		if err != nil {
			errorLog.Fatal(err)
		}
	}

	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		todos:         &mysql.TodoModel{DB: db},
		templateCache: templateCache,
		users:         &mysql.UserModel{DB: db},
	}

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s\n", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
