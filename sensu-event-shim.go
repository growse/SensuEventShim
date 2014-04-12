package sensu-event-shim

import (
  "database/sql"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
  
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":5151", nil)
}
