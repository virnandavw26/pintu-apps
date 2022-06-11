package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestHelloServerHandler(t *testing.T) {
  r := httptest.NewRequest("GET", "/", nil)
  w := httptest.NewRecorder()
  h := http.HandlerFunc(HelloServer)

  h.ServeHTTP(w, r)

  // check HTTP response status code
  if s := w.Code; s != http.StatusOK {
          t.Errorf("handler return wrong status code: got %v, want %v", s, http.StatusOK)
  }

  // check HTTP response body
  want := "Hello, Pintu! It's Nanda :)"
  if w.Body.String() != want {
          t.Errorf("handler return wrong status code: got %v, want %v", w.Body.String(), want)
  }
}
