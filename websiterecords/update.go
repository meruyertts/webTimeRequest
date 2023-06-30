package websiterecords

import (
	"math"
	"time"
)

func (wi *WebsiteInfoExec) UpdateTimeInfo(url string) time.Duration {
	wi.Mutex.Lock()
	defer wi.Mutex.Unlock()
	if duration, ok := wi.websitesDuration[url]; ok {
		return duration
	}
	return 0
}

func (wi *WebsiteInfoExec) UpdateMinTimeInfo() string {
	wi.Mutex.Lock()
	defer wi.Mutex.Unlock()
	minTime := time.Duration(math.MaxInt64)
	minURL := ""
	for webURL, webDuration := range wi.websitesDuration {
		if webDuration < minTime {
			minTime = webDuration
			minURL = webURL
		}
	}
	return minURL
}
func (wi *WebsiteInfoExec) UpdateMaxTimeInfo() string {
	wi.Mutex.Lock()
	defer wi.Mutex.Unlock()
	maxTime := time.Duration(math.MinInt64)
	maxURL := ""
	for webURL, webDuration := range wi.websitesDuration {
		if webDuration > maxTime {
			maxTime = webDuration
			maxURL = webURL
		}
	}
	return maxURL
}
