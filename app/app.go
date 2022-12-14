package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct { // Hold the Application, expose reference to the router and the DB that the App uses
	Router mux.Router
	DB     *sql.DB
}

// Initialize the Application
func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = *mux.NewRouter()
	a.Router.HandleFunc("/salut", func(resp http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(resp, "Salut - Gorilla!")
	})
}

// Run it
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, &a.Router))
}
