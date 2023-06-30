package websiterecords

import "time"

func (wi *WebsiteInfoExec) CheckWebsite(websites []WebsiteRecords) {
	wi.Mutex.Lock()
	defer wi.Mutex.Unlock()
	wi.websites = websites
}
func (wi *WebsiteInfoExec) CheckWebsiteTime(url string, start time.Time) {
	wi.Mutex.Lock()
	defer wi.Mutex.Unlock()
	wi.websitesDuration[url] = time.Since(start)

}
