package main

import (
  "net/http"
  "fmt"
)

// DictionaryServer is a stand-in for the real api
func DictionaryServer(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "breakfast")
}
