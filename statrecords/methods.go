package statrecords

func (si *StatInfoExec) UpdateTimeRequests() {
	si.Mutex.Lock()
	defer si.Mutex.Unlock()
	si.statrecs.AccesStats++
}
func (si *StatInfoExec) UpdateMinTimeRequests() {
	si.Mutex.Lock()
	defer si.Mutex.Unlock()
	si.statrecs.MinStats++
}
func (si *StatInfoExec) UpdateMaxTimeRequests() {
	si.Mutex.Lock()
	defer si.Mutex.Unlock()
	si.statrecs.MaxStats++
}

func (si *StatInfoExec) GetStatRecords() StatRecords {
	si.Mutex.Lock()
	defer si.Mutex.Unlock()
	return si.statrecs
}
