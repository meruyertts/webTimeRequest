package websitestats

import "time"

func (ws *WebsiteStats) CheckWebsite() {
	websites := ws.WebsiteInfo.GetWebsites()
	for i := range websites {
		websites[i].Access, websites[i].AccessTime = ws.WebsiteAccess.CheckAvailability(websites[i].URL)
		if websites[i].Access {
			ws.WebsiteInfo.CheckWebsiteTime(websites[i].URL, websites[i].AccessTime)
		}

	}
	ws.WebsiteInfo.CheckWebsite(websites)
}

func (ws *WebsiteStats) UpdateTimeInfo(url string) time.Duration {
	return ws.WebsiteInfo.UpdateTimeInfo(url)
}

func (ws *WebsiteStats) UpdateMinTimeInfo() string {
	return ws.WebsiteInfo.UpdateMinTimeInfo()
}

func (ws *WebsiteStats) UpdateMaxTimeInfo() string {
	return ws.WebsiteInfo.UpdateMaxTimeInfo()
}
