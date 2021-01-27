package main

import (
  "database/sql"
  "encoding/json"
  "errors"
  "fmt"
  "html/template"
  "log"
  "net/http"
  "strings"
  // "sync"

  _ "github.com/mattn/go-sqlite3"
)

// Entry holds information for a word
type Entry struct {
  Word string `json:"word"`
  Def string `json:"def"`
  Gender string `json:"gender"`
}

type Error struct {
  Error string `json:"error"`
}

var (
  // ip string = "129.119.172.61"
  ip string = "localhost"
  port string = "8000"
  dbPath string = "dict.db"
  db *sql.DB
  temp *template.Template
)

func main() {
  var err error
  db, err = sql.Open("sqlite3", dbPath)
  if err != nil {
    log.Panic(err)
  }
  defer db.Close()

  parse()

  fs := http.FileServer(http.Dir("templates"))
  http.Handle("/static/", http.StripPrefix("/static", fs))
  http.HandleFunc("/", pageHandler)
  http.HandleFunc("/api", apiHandler)
  log.Panic(http.ListenAndServe(ip + ":" + port, nil))
}

func parse() {
  temp = template.Must(template.ParseFiles("templates/index.html"))
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
  temp.Execute(w, nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
  if err := r.ParseForm(); err != nil {
    log.Println(err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }
  w.Header()["Content-Type"] = []string{"application/json"}
  encoder := json.NewEncoder(w)
  word := strings.Join(r.Form["word"], " ")
  if word == "" {
    if err := encoder.Encode(Error{"must provide word"}); err != nil {
      log.Println(err)
    }
    return
  }
  if r.Method == http.MethodPost {
    method := r.FormValue("method")
    oldDef := strings.Join(r.Form["old"], " ")
    newDef := strings.Join(r.Form["new"], " ")
    if method == "suggestion" {
      /* Handle suggestions */
    } else if method == "admin" {
      email, password := r.FormValue("email"), r.FormValue("password")
      row := db.QueryRow(fmt.Sprintf(`SELECT * FROM users WHERE email="%s"`, email))
      var e, p string
      if err := row.Scan(&e, &p); errors.Is(err, sql.ErrNoRows) {
        if err := encoder.Encode(Error{"invalid username or password"}); err != nil {
          log.Println(err)
        }
        return
      } else if err != nil {
        log.Println(err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
      }
      /* Hash password */
      if p != password {
        if err := encoder.Encode(Error{"invalid username or password"}); err != nil {
          log.Println(err)
        }
        return
      }
      /* Create, Update, or Delete word */
      println(oldDef, newDef)
    } else {
      if err := encoder.Encode(Error{"invalid method"}); err != nil {
        log.Println(err)
        return
      }
    }
    return
  }
  stmt := `SELECT * FROM words`
  if word == "*" {
    stmt += fmt.Sprintf(` ORDER BY word`)
  } else {
      stmt += fmt.Sprintf(` WHERE word="%s" ORDER BY def`, word)
  }
  rows, err := db.Query(stmt)
  if err != nil {
    log.Println(err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }
  defer rows.Close()
  var def, gender string
  var entries []Entry
  for rows.Next() {
    if err := rows.Scan(&word, &def, &gender); err != nil {
      log.Println(err)
      http.Error(w, "Internal Server Error", http.StatusInternalServerError)
      return
    }
    entries = append(entries, Entry{word, def, gender})
  }
  if err := encoder.Encode(entries); err != nil {
    log.Println(err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
  }
}
