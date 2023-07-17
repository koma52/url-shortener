package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UrlBody struct {
	Url string `json:"url"`
}

type MessageBody struct {
	Message string `json:"message"`
}

type InfoBody struct {
	ShortCode int64  `json:"shortcode"`
	LongUrl   string `json:"url"`
	Active    bool   `json:"active"`
	Created   string `json:"created"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Nothing here yet")
}

func (a *App) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortcode := vars["shortcode"]

	var longUrl string
	var active bool

	dbSelect := "SELECT longurl,active FROM shortenedurls WHERE shortcode=" + shortcode

	a.DB.QueryRow(dbSelect).Scan(&longUrl, &active)

	if active {
		http.Redirect(w, r, longUrl, http.StatusSeeOther)
	} else {
		var e MessageBody
		e.Message = "Url is not active"
		respond(e, http.StatusNotAcceptable, w)
	}
}

func (a *App) InfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortcode := vars["shortcode"]

	var shortCode int64
	var longUrl string
	var active bool
	var created string

	var i InfoBody

	dbSelect := "SELECT * FROM shortenedurls WHERE shortcode=" + shortcode

	a.DB.QueryRow(dbSelect).Scan(&shortCode, &longUrl, &active, &created)

	i.ShortCode = shortCode
	i.LongUrl = longUrl
	i.Active = active
	i.Created = created

	respond(i, http.StatusOK, w)
}

func (a *App) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	var u UrlBody

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		fmt.Println(err)
		return
	}
	defer r.Body.Close()

	dbInsert := "INSERT INTO shortenedurls (longurl, active) VALUES ('" + u.Url + "', 1)"

	res, _ := a.DB.Exec(dbInsert)

	id, _ := res.LastInsertId()

	u.Url = r.Host + "/" + strconv.FormatInt(id, 10)

	respond(u, http.StatusOK, w)
}

func (a *App) ToggleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortcode := vars["shortcode"]

	var active bool
	var dbUpdate string

	dbSelect := "SELECT active FROM shortenedurls WHERE shortcode=" + shortcode
	a.DB.QueryRow(dbSelect).Scan(&active)

	if active {
		dbUpdate = "UPDATE shortenedurls SET active=0 WHERE shortcode=" + shortcode
	} else {
		dbUpdate = "UPDATE shortenedurls SET active=1 WHERE shortcode=" + shortcode
	}

	a.DB.Exec(dbUpdate)

	w.WriteHeader(http.StatusOK)
}

func (a *App) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortcode := vars["shortcode"]

	dbDelete := "DELETE FROM shortenedurls WHERE shortcode=" + shortcode

	a.DB.Exec(dbDelete)

	w.WriteHeader(http.StatusOK)
}

func respond(body interface{}, resCode int, w http.ResponseWriter) {
	response, _ := json.Marshal(body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resCode)
	w.Write(response)
}
