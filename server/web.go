package main

import (
  "fmt"
  "html/template"
  "log"
  "net/http"
  "os"
)

var (
  webLogger *log.Logger
  errorPage string
)

func init() {
  webLogger = log.New(os.Stdout, "Web: ", log.LstdFlags)
  f, err := os.Open(tf+"error.html")
  defer f.Close()
  if err != nil {
    webLogger.Panic(err)
  }
  var blines [512]byte
  l, err := f.Read(blines[:])
  if err != nil {
    webLogger.Panic(err)
  }
  errorPage = string(blines[:l])
}

func RunWeb() {
  server := http.Server{
    Addr: IP + WEB_PORT,
    Handler: routes(),
  }
  webLogger.Panic(server.ListenAndServe())
}

func routes() *http.ServeMux {
  r := http.NewServeMux()
  r.HandleFunc("/", homeHandler)

  static := http.FileServer(http.Dir("../static"))
  r.Handle("/static/", http.StripPrefix("/static", static))

  return r
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  ts, err := template.ParseFiles(tf+"indexs.html")
  if err != nil {
    webLogger.Println(err)
    fmt.Fprintf(w, errorPage, "500")
    return
  }
  ts.Execute(w, nil)
}