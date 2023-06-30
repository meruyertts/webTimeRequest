package main

import (
	"time"
	"webTimeRequest/handler"
	"webTimeRequest/statrecords"
	"webTimeRequest/websiteaccess"
	"webTimeRequest/websiterecords"
	"webTimeRequest/websitestats"
)

func main() {
	websiteStat := &websitestats.WebsiteStats{
		WebsiteAccess: &websiteaccess.WebsiteAccessExec{},
		WebsiteInfo:   websiterecords.NewWebsiteInfo(),
		StatInfo:      &statrecords.StatInfoExec{},
	}
	go func() {
		websiteStat.CheckWebsite()
		time.Sleep(time.Minute)
	}()
	newhandler := handler.NewHandler(*websiteStat)
	newhandler.InitRoutes()

}
