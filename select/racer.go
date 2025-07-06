package racer

import (
  "net/http"
  "time"
)

var tenSecondTimeout = 10 * time.Second

type RacerError string

func (e RacerError) Error() string {
  return string(e)
}

const (
  ErrRacerError = RacerError("Reached timeout.")
)

func Racer(a, b string) (winner string, error error) {
  return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
  select {
  case <-ping(a):
    return a, nil
  case <-ping(b):
    return b, nil
  case <-time.After(timeout):
    return "", ErrRacerError
  }
}

func ping(url string) chan struct{} {
  ch := make(chan struct{})
  go func() {
    http.Get(url)
    close(ch)
  }()
  return ch
}
