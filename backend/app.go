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

	// Wait for db to start
	time.Sleep(5 * time.Second)
	connectionString := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(db:3306)/" + os.Getenv("MYSQL_DATABASE")
	a.DB, _ = sql.Open("mysql", connectionString)

	a.Router.HandleFunc("/", HomeHandler).Methods("GET")
	a.Router.HandleFunc("/{shortcode}", a.RedirectHandler).Methods("GET")
	a.Router.HandleFunc("/info/{shortcode}", a.InfoHandler).Methods("GET")

	a.Router.HandleFunc("/shorten", a.ShortenHandler).Methods("POST")

	a.Router.HandleFunc("/{shortcode}", a.ToggleHandler).Methods("PUT")

	a.Router.HandleFunc("/{shortcode}", a.DeleteHandler).Methods("DELETE")
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":52520", a.Router))
}
