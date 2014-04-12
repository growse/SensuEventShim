package main

import (
    _ "github.com/lib/pq"
    "database/sql"
    "net/http"
    "log"
    "io/ioutil"
)

func eventhandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            log.Print(err)
            http.Error(w, "Bad Request", 400)
            return
        }
        _, err = stmt.Exec(body)
        if err != nil {
            log.Print(err)
            http.Error(w, "Bad Request", 400)
            return
        }
    } else {
        http.Error(w, "Bad Request", 400)
    }
}

var db *sql.DB
var stmt *sql.Stmt

func main() {
    var err error
    db, err = sql.Open("postgres", "user=sensuevents dbname=sensuevents password=password")
    if err != nil {
        log.Fatal(err)
    }
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
    stmt, err = db.Prepare("INSERT INTO events (event) VALUES($1)")
    if err != nil {
        log.Fatal(err)
    }
    http.HandleFunc("/events", eventhandler)
    http.ListenAndServe(":5151", nil)
}
