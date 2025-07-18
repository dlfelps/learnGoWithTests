package main

import (
  "time"
)

type SpySleeper struct {
  Calls int
}

func (s *SpySleeper) Sleep() {
  s.Calls++
}

type SpyCountdownOperations struct {
  Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
  s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
  s.Calls = append(s.Calls, write)
  return
}

const write = "write"
const sleep = "sleep"

type SpyTime struct {
  durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
  s.durationSlept = duration
}