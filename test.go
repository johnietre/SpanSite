package main

import (
  "encoding/json"
  "net/http"
)

type TestStr struct {
  Firstname string `json:"firstname"`
  Middlename string `json:"middlename"`
  Lastname string `json:"lastname"`
  Suffix string `json:"suffix"`
}

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    writer := json.NewEncoder(w)
    err := writer.Encode(TestStr{"Johnie", "Vance", "Rodgers", "III"})
    if err != nil {
      println(err.Error())
    }
  })
  http.ListenAndServe("localhost:8080", nil)
}
