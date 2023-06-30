package websiterecords

import "time"

func (wi *WebsiteInfoExec) GetWebsites() []WebsiteRecords {
	wi.Mutex.Lock()
	defer wi.Mutex.Unlock()
	return wi.websites
}
func NewWebsiteInfo() *WebsiteInfoExec {
	websiteList := Websites
	webURL := make([]WebsiteRecords, len(websiteList))
	for i := range websiteList {
		webURL[i].URL = websiteList[i]
	}
	return &WebsiteInfoExec{
		websites:         webURL,
		websitesDuration: make(map[string]time.Duration),
	}
}
