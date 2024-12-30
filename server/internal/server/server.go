package server

import (
	s "github.com/jrjaro18/elastic-raft-go/server/internal/store"
)

type IP string

type Server struct {
	Self  IP
	Peers []IP
	Store *s.Store
	Term  uint
}
