package main

import (
    _ "github.com/lib/pq"
    "database/sql"
    "net/http"
    "log"
    "os"
    "fmt"
    "encoding/json"
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

type Configuration struct {
    Dbuser string
    Dbname string
    Dbpassword string
    Dbhost string
    Dbport int
}

func main() {
    var err error
    file, _ := os.Open("sensueventshim.json")
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    decoder.Decode(&configuration)
    db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", configuration.Dbhost, configuration.Dbport, configuration.Dbuser, configuration.Dbname, configuration.Dbpassword))
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
