package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
  store := NewInMemoryPlayerStore()
  server := NewPlayerServer(store)
  player := "Pepper"

  pepperWin := newPostWinRequest(player)
  server.ServeHTTP(httptest.NewRecorder(), pepperWin)
  server.ServeHTTP(httptest.NewRecorder(), pepperWin)
  server.ServeHTTP(httptest.NewRecorder(), pepperWin)

  response := httptest.NewRecorder()
  server.ServeHTTP(response, newGetScoreRequest(player))
  assertStatus(t, response.Code, http.StatusOK)

  assertResponseBody(t, response.Body.String(), "3")
}