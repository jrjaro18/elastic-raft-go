package main

import (
	"time"

	s "github.com/jrjaro18/elastic-raft-go/server/internal/store"
)

func main() {
	store := s.NewStore()
	store.Add("riyan", "parag")
	store.Add("janhavi", "kapoor")
	store.Add("ananya", "pandey")
	store.Remove("janhavi")
	store.Add("ananya", "parag")
	time.Sleep(5 * time.Second)
	store.RebootLogFile()
}