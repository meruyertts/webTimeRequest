package websiteaccess

import (
	"net/http"
	"time"
)

func (wa *WebsiteAccessExec) CheckAvailability(url string) (bool, time.Time) {
	start := time.Now()
	resp, err := http.Get("https://" + url)
	if err != nil {
		return false, start
	}
	if resp.StatusCode != 200 {
		return false, start

	}
	resp.Body.Close()
	return true, start
}
