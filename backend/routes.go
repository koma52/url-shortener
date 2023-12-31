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

	err := a.DB.QueryRow(dbSelect).Scan(&longUrl, &active)
	if err != nil {
		var e MessageBody
		e.Message = "Could not find shortcode"
		respond(e, http.StatusNotFound, w)
		return
	}

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

	err := a.DB.QueryRow(dbSelect).Scan(&shortCode, &longUrl, &active, &created)
	if err != nil {
		var e MessageBody
		e.Message = "Could not find shortcode"
		respond(e, http.StatusNotFound, w)
		return
	}

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

	res, err := a.DB.Exec(dbInsert)
	if err != nil {
		var e MessageBody
		e.Message = "Something went wrong while shortening"
		respond(e, http.StatusInternalServerError, w)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		var e MessageBody
		e.Message = "Something went wrong while reading db"
		respond(e, http.StatusInternalServerError, w)
		return
	}

	u.Url = r.Host + "/" + strconv.FormatInt(id, 10)

	respond(u, http.StatusOK, w)
}

func (a *App) ToggleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortcode := vars["shortcode"]

	var active bool
	var dbUpdate string
	var e MessageBody

	dbSelect := "SELECT active FROM shortenedurls WHERE shortcode=" + shortcode
	err := a.DB.QueryRow(dbSelect).Scan(&active)
	if err != nil {
		var e MessageBody
		e.Message = "Could not find shortcode"
		respond(e, http.StatusNotFound, w)
		return
	}

	if active {
		dbUpdate = "UPDATE shortenedurls SET active=0 WHERE shortcode=" + shortcode
		e.Message = "Toggled inactive"
	} else {
		dbUpdate = "UPDATE shortenedurls SET active=1 WHERE shortcode=" + shortcode
		e.Message = "Toggled active"
	}

	_, err = a.DB.Exec(dbUpdate)
	if err != nil {
		var e MessageBody
		e.Message = "Something went wrong while updating db"
		respond(e, http.StatusInternalServerError, w)
		return
	}

	respond(e, http.StatusOK, w)
}

func (a *App) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortcode := vars["shortcode"]

	dbDelete := "DELETE FROM shortenedurls WHERE shortcode=" + shortcode
	dbSelect := "SELECT shortcode FROM shortenedurls WHERE shortcode=" + shortcode

	// Check if shortened url exists
	err := a.DB.QueryRow(dbSelect).Scan(&shortcode)
	if err != nil {
		var e MessageBody
		e.Message = "Could not find shortcode"
		respond(e, http.StatusNotFound, w)
		return
	}

	_, err = a.DB.Exec(dbDelete)
	if err != nil {
		var e MessageBody
		e.Message = "Something went wrong while deleting"
		respond(e, http.StatusInternalServerError, w)
		return
	}

	var e MessageBody
	e.Message = "Deleted shortcode"
	respond(e, http.StatusOK, w)
}

func respond(body interface{}, resCode int, w http.ResponseWriter) {
	response, _ := json.Marshal(body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resCode)
	w.Write(response)
}
