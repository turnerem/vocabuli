package main

//import (
//  "testing"
//  "net/http"
//  "net/http/httptest"
//)
//
//func TestGETTranslation(t *testing.T) {
//  t.Run("returns translation for fruhstuck", func(t *testing.T) {
//    request, _ := http.NewRequest(http.MethodGet, "/translate/fruhstuck", nil)
//    response := httptest.NewRecorder()
//
//    DictionaryServer(response, request)
//
//    got := response.Body.String()
//    want := "breakfast"
//
//    if got != want {
//      t.Errorf("got %q want %q", got, want)
//    }
//  })
//}
