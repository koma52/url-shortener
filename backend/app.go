package backend

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	var err error

	// Wait for db to start
	time.Sleep(5 * time.Second)
	connectionString := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv("DB_URL") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("MYSQL_DATABASE")
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Invalid DB connection string: %v", err)
	}
	if err = a.DB.Ping(); err != nil {
		log.Fatalf("Couldn't connect to database: %v", err)
	}

	a.Router.HandleFunc("/", HomeHandler).Methods("GET")
	a.Router.HandleFunc("/{shortcode}", a.RedirectHandler).Methods("GET")
	a.Router.HandleFunc("/api/info/{shortcode}", a.InfoHandler).Methods("GET")
	a.Router.HandleFunc("/api/list", a.ListHandler).Methods("GET")

	a.Router.HandleFunc("/api/shorten", a.ShortenHandler).Methods("POST")

	a.Router.HandleFunc("/api/toggle/{shortcode}", a.ToggleHandler).Methods("PUT")

	a.Router.HandleFunc("/api/delete/{shortcode}", a.DeleteHandler).Methods("DELETE")
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), a.Router))
}
