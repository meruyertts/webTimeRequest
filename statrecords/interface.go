package statrecords

import "sync"

type StatRecords struct {
	AccesStats int
	MinStats   int
	MaxStats   int
}

type StatInfo interface {
	GetStatRecords() StatRecords
	UpdateTimeRequests()
	UpdateMinTimeRequests()
	UpdateMaxTimeRequests()
}
type StatInfoExec struct {
	statrecs StatRecords
	Mutex    sync.Mutex
}
